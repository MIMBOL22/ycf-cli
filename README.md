 # YCF-CLI
 ## CLI для работы с cервисом Yandex Cloud Functions
На данный момент поддерживается только деплой функций

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
