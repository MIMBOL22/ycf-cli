package integration

/*
	Задумка в следующем:
	У любой интеграции есть функции сборки и деплоя, которые
	Вызывают внутри себя другие подфункции.

	Возможно, в будущем появится функция тестирования.
*/

type Integration struct {
	Type         string `yaml:"type"`
	Dir          string `yaml:"dir"`
	BucketName   string `yaml:"bucket_name"`
	Cmd          string `yaml:"cmd"`
	DistDir      string `yaml:"dist_dir"`
	FunctionList []Function
	GatewayList  []Gateway
	FrontendList []Frontend
}
