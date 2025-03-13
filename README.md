 # YCF-CLI
 ## CLI для работы с cервисом Yandex Cloud Functions
На данный момент поддерживается только деплой функций

Пример проекта можно увидеть [здесь](https://github.com/MIMBOL22/ycf-cli-template/)
Скачать CLI можно [тут](https://github.com/MIMBOL22/ycf-cli/releases)

 ## Настройка проекта
 ### Главный конфиг
Для начала, вам необходимо создать в корне своего проекта главный конфиг с названием mono.ycf.yaml :
```yaml
base_dir: ./functions
service_account_key_path: ./authorized_key.json
auth_type: oauth
s3_bucket: for-function
```
 + base_dir - папка с функциями (каждая функция, как отдельная папка)
 + service_account_key_path - путь до ключа сервисного аккаунта (если выбран такой способ аутентификации)
 + auth_type - способ авторизации в Yandex Cloud API: `oauth`, `instance` или `service_account`
 + s3_bucket - имя S3 бакета в Yandex Cloud, через который будут загружены функции

 ### .env
Далее, настраиваем секреты
```env
OAUTH_TOKEN=gDfcbR38svDn59xgGde
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
```
 + OAUTH_TOKEN - OAuth токен для авторизации в Yandex Cloud API (если выбран такой способ аутентификации) Получить его можно [здесь](https://oauth.yandex.ru/authorize?response_type=token&client_id=1a6990aa636648e9b2ef855fa7bec2fb)
 + AWS_ACCESS_KEY_ID - Идентификатор статичного ключа сервисного аккаунта (используется для S3)
 + AWS_SECRET_ACCESS_KEY - Секрет статичного ключа сервисного аккаунта (используется для S3)
   
 ### Конфиги функции
 Теперь необходимо в каждой папке функции (напомню, каждая из этих папок должна находиться в указанном base_dir) c названием ycf.yaml:
 ```yaml
 environments:
  - name: production
    id: d4enfpb6roib91vangsn
    memory: 128
    runtime: nodejs18
    description: "d"
    entrypoint:
      file: src/index.ts
      function: handler
    timeout: 3
    service_account_id: ajeq7unktvsanbnpqg6u
    additional_files:
      - test.json
  - name: dev
    id: d4e8cgis2shkqamr5pov
    memory: 128
    runtime: nodejs18
    description: "d"
    entrypoint:
      file: src/index.ts
      function: handler
    timeout: 3
    service_account_id: ajeqqg6u7unktvsanbnp
    additional_files:
      - test.json
 ```
 + environments - Окружения (dev, prod и т.д.) (массив)
 + environments.name - Название окружения (необходимо только для выбора на деплое, в Cloud не загружается)
 + environments.id - ID функции в Yandex Cloud, куда будет загружаться функция
 + environments.memory - Выделенная память (в МБ)
 + environments.runtime - Тип runtime-a (в нашем случае - nodejs18, но в будущем могут быть и более свежие версии)
 + environments.description - Описание версии функции
 + environments.entrypoint.file - Файл с точкой входа в функцию
 + environments.entrypoint.function - Фунция в файле, которая является точкой входа
 + environments.timeout - Лимит времени исполнения (в секундах)
 + environments.service_account_id - ID Сервисного аккаунта, доступ к которой будет предоставлен функции (IAM которого будет прокинут в ctx)
 + environments.additional_files - Массив файлов, которые не берутся сборщиком, но должны быть в итоговом бандле функции

 ### Непосредственный запуск 
 Т.к. на данный момент поддерживается только деплой, в CLI только одна команда:
 + deploy
 + + Пример:
   + ```bash
     ycf-cli deploy
     ```
 + + Пример c указанием окружения:
   + ```bash
     ycf-cli deploy --env dev
     ```
