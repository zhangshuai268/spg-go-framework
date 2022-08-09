package generator

import "os"

const (
	configExample = "{\n  \"api\" : {\n    \"api_port\" : \"8080\",\n    \"api_secret\": \"sZWJiJVxhQofXNMpDGnOkVBOQoTWyTHm\"\n  },\n  \"mysql\" : {\n    \"driver\": \"mysql\",\n    \"user\": \"root\",\n    \"pass_word\": \"123456\",\n    \"host\": \"127.0.0.1\",\n    \"port\": \"3306\",\n    \"db_name\": \"localhost\",\n    \"charset\": \"utf8mb4\",\n    \"show_sql\": true,\n    \"parseTime\": \"true\",\n    \"loc\": \"Asia/Shanghai\"\n  }\n}"
	configInit    = "package config\n\nimport (\n\t\"encoding/json\"\n\t\"errors\"\n\t\"io/ioutil\"\n\t\"os\"\n\t\"path/filepath\"\n)\n\nvar Conf *Config\n\nfunc InitConfig() (*Config, error) {\n\n\tdir, err := os.Getwd()\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\tconfig := new(Config)\n\n\tpath := filepath.Join(dir, \"internal/config/config.json\")\n\tif os.Getenv(\"STAGE\") != \"\" {\n\t\tpath = filepath.Join(dir, \"internal/config/config_\"+os.Getenv(\"STAGE\")+\".json\")\n\t}\n\tfile, err := os.Open(path)\n\tif err != nil {\n\t\treturn nil, errors.New(\"打开配置文件错误\" + path + err.Error())\n\t}\n\n\tconfByte, err := ioutil.ReadAll(file)\n\tif err != nil {\n\t\treturn nil, errors.New(\"读取配置文件错误\" + err.Error())\n\t}\n\n\terr = json.Unmarshal(confByte, config)\n\tif err != nil {\n\t\treturn nil, errors.New(\"读取配置文件错误\" + err.Error())\n\t}\n\n\tConf = config\n\n\treturn Conf, nil\n}\n"
)

// FrameworkGenerator 生成目录结构
func FrameworkGenerator() error {
	err := os.MkdirAll("api/docker", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll("api/swagger", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll("cmd", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll("internal/api", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll("internal/config", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll("internal/crontab", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll("internal/model", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll("internal/pkg", os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// ConfigGenerator 生成配置相关文件
func ConfigGenerator() error {
	//生成json实例文件
	fce, err := os.Create("./internal/config/config_example.json")
	if err != nil {
		return err
	}
	fce.WriteString(configExample)
	//转移json文件
	err = os.Rename("./config.json", "./internal/config/config.json")
	if err != nil {
		return err
	}
	//生成config_init.go文件
	fci, err := os.Create("./internal/config/config_init.go")
	if err != nil {
		return err
	}
	fci.WriteString(configInit)
	//根据json文件生成config结构体
	GenerateModels("./internal/config/config.json", "./internal/config/config.go")
	return nil
}

func internalGenerator(name string) error {
	//生成docker
	err := os.MkdirAll("internal/api/"+name, os.ModePerm)
	if err != nil {
		return err
	}

}

func dockerGenerator(name string) error {
	//生成docker
	err := os.MkdirAll("api/docker/"+name, os.ModePerm)
	if err != nil {
		return err
	}
	//生成docker文件
	fcd, err := os.Create("api/docker/" + name)
	if err != nil {
		return err
	}
	fcd.WriteString()
}
