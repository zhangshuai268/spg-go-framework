package generator

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

// FrameworkGenerator 生成目录结构
//TODO: FrameworkGenerator
func FrameworkGenerator() error {
	getwd, err := os.Getwd()
	if err != nil {
		return err
	}
	err = os.MkdirAll(getwd+"/api/docker", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(getwd+"/api/swagger", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(getwd+"/cmd", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(getwd+"/internal/api", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(getwd+"/internal/config", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(getwd+"/internal/crontab", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(getwd+"/internal/model", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(getwd+"/internal/pkg/middle", os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(getwd+"/internal/pkg/code", os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// ConfigGenerator 生成配置相关文件
//TODO: ConfigGenerator
func ConfigGenerator() error {
	getwd, err := os.Getwd()
	if err != nil {
		return err
	}
	//生成config_init.go文件
	fci, err := os.Create(getwd + "/internal/config/config_init.go")
	if err != nil {
		return err
	}
	fci.WriteString("package config\n\nimport (\n\t\"encoding/json\"\n\t\"errors\"\n\t\"io/ioutil\"\n\t\"os\"\n\t\"path/filepath\"\n)\n\nvar Conf *Config\n\nfunc InitConfig(paths ...string) (*Config, error) {\n\tvar path string\n\tif len(paths) == 0 {\n\t\tdir, err := os.Getwd()\n\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n\t\tpath = filepath.Join(dir, \"config.json\")\n\t\tif os.Getenv(\"STAGE\") != \"\" {\n\t\t\tpath = filepath.Join(dir, \"config_\"+os.Getenv(\"STAGE\")+\".json\")\n\t\t}\n\t} else {\n\t\tpath = paths[0]\n\t}\n\n\tconfig := new(Config)\n\tfile, err := os.Open(path)\n\tif err != nil {\n\t\treturn nil, errors.New(\"打开配置文件错误\" + path + err.Error())\n\t}\n\n\tconfByte, err := ioutil.ReadAll(file)\n\tif err != nil {\n\t\treturn nil, errors.New(\"读取配置文件错误\" + err.Error())\n\t}\n\n\terr = json.Unmarshal(confByte, config)\n\tif err != nil {\n\t\treturn nil, errors.New(\"读取配置文件错误\" + err.Error())\n\t}\n\n\tConf = config\n\n\treturn Conf, nil\n}\n")
	//根据json文件生成config结构体
	GenerateModels(getwd+"/config.json", getwd+"/internal/config/config.go")
	return nil
}

// StoreGenerator 生成dao层目录和文件
//TODO: StoreGenerator
func StoreGenerator() error {
	//生成store目录
	err := os.MkdirAll("internal/api/store", os.ModePerm)
	if err != nil {
		return err
	}
	//生成工厂模式文件
	//生产文件
	fcd, err := os.Create("internal/api/store/store.go")
	if err != nil {
		return err
	}
	fcd.WriteString("package store\n\ntype datastore struct {\n}\n\nvar DataStore Factory\n\nfunc GetFactory() (Factory, error) {\n\n\tDataStore = &datastore{}\n\n\treturn DataStore, nil\n}\n")
	//工厂文件
	fcf, err := os.Create("internal/api/store/factory.go")
	if err != nil {
		return err
	}
	fcf.WriteString("package store\n\nvar client Factory\n\ntype Factory interface {\n}\n\nfunc Client() Factory {\n\treturn client\n}\n\nfunc SetClient(factory Factory) {\n\tclient = factory\n}\n")
	return nil
}

// InternalGenerator 生成主要代码目录
//TODO: InternalGenerator
func InternalGenerator(name string) error {
	//生成权限控制目录
	err := os.MkdirAll("internal/api/"+name+"/auth", os.ModePerm)
	if err != nil {
		return err
	}
	fca, err := os.Create("internal/api/" + name + "/auth/auth.go")
	if err != nil {
		return err
	}
	fca.WriteString("package auth")
	//生成控制器目录
	err = os.MkdirAll("internal/api/"+name+"/controller", os.ModePerm)
	if err != nil {
		return err
	}
	//生成业务逻辑代码目录
	err = os.MkdirAll("internal/api/"+name+"/service", os.ModePerm)
	if err != nil {
		return err
	}
	fcs, err := os.Create("internal/api/" + name + "/service/service.go")
	if err != nil {
		return err
	}
	path := getCurrentPath()
	fcs.WriteString("package service\n\nimport (\n\t\"" + path + "/internal/api/store\"\n)\n\ntype Service interface {\n}\n\ntype service struct {\n\tfactory store.Factory\n}\n\nfunc NewService(factory store.Factory) Service {\n\treturn &service{\n\t\tfactory: factory,\n\t}\n}\n")
	//生成中间件代码
	fcm, err := os.Create("internal/pkg/middle/middle.go")
	if err != nil {
		return err
	}
	fcm.WriteString("package middle\n\nimport (\n\t\"errors\"\n\t\"fmt\"\n\t\"github.com/dgrijalva/jwt-go\"\n\t\"github.com/gin-gonic/gin\"\n\t\"net/http\"\n)\n\nfunc CORS() gin.HandlerFunc {\n\treturn func(c *gin.Context) {\n\t\tmethod := c.Request.Method\n\t\torigin := c.Request.Header.Get(\"Origin\")\n\t\tif origin != \"\" {\n\t\t\t//c.Writer.Header().Set(\"Access-Control-Allow-Origin\", \"*\")\n\t\t\tc.Header(\"Access-Control-Allow-Origin\", \"*\") // 可将将 * 替换为指定的域名\n\t\t\tc.Header(\"Access-Control-Allow-Methods\", \"POST, GET, OPTIONS, PUT, DELETE, UPDATE\")\n\t\t\tc.Header(\"Access-Control-Allow-Headers\", \"Origin, X-Requested-With, Content-Type, Accept, Authorization\")\n\t\t\tc.Header(\"Access-Control-Expose-Headers\", \"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type\")\n\t\t\tc.Header(\"Access-Control-Allow-Credentials\", \"true\")\n\t\t}\n\t\tif method == \"OPTIONS\" {\n\t\t\tc.AbortWithStatus(http.StatusNoContent)\n\t\t}\n\t\tc.Next()\n\t}\n}\n\nfunc ErrHandler() gin.HandlerFunc {\n\treturn func(c *gin.Context) {\n\t\tdefer func() {\n\t\t\tif err := recover(); err != nil {\n\t\t\t\tfmt.Println(err)\n\t\t\t}\n\t\t\tc.Abort()\n\t\t}()\n\t\tc.Next()\n\t}\n}\n\n// GetToken 生成token\n//  secret: 项目密钥\n//  claims: 自定义jwt加密结构体\nfunc GetToken(secret string, claims jwt.Claims) (string, error) {\n\tappSecret := secret\n\tjwtKey := []byte(appSecret)\n\ttoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)\n\ttokenString, err := token.SignedString(jwtKey)\n\tif err != nil {\n\t\treturn \"\", err\n\t}\n\n\treturn tokenString, nil\n}\n\n// ParseToken 解密token\n//  tokenString: 解密方法\nfunc ParseToken(tokenString string, secret string, claims jwt.Claims) error {\n\n\tjwtKey := []byte(secret)\n\ttoken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {\n\t\treturn jwtKey, nil\n\t})\n\tif err == nil && token != nil {\n\t\tif token.Valid {\n\t\t\treturn nil\n\t\t} else {\n\t\t\treturn errors.New(\"token解析失败\")\n\t\t}\n\t}\n\treturn err\n}\n")
	//生成返回值控制代码
	fcc, err := os.Create("internal/pkg/code/code.go")
	if err != nil {
		return err
	}
	fcc.WriteString("package code\n\nimport \"github.com/gin-gonic/gin\"\n\nfunc BuildReturn(context *gin.Context, state int, message string, data interface{}) {\n\n\tcontext.JSON(200, gin.H{\n\t\t\"status\":  state,\n\t\t\"message\": message,\n\t\t\"data\":    data,\n\t})\n\treturn\n}\n")
	//生成路由文件
	fcr, err := os.Create("internal/api/" + name + "/route.go")
	fcr.WriteString("package " + name + "\n\nimport (\n\t\"github.com/gin-gonic/gin\"\n\t\"" + path + "/internal/api/store\"\n\t\"" + path + "/internal/pkg/middle\"\n)\n\nfunc RouterInit(factory store.Factory) (*gin.Engine, error) {\n\trouter := gin.Default()\n\trouter.Use(gin.Logger())\n\trouter.Use(gin.Recovery())\n\n\t//全局异常监控\n\trouter.Use(middle.ErrHandler())\n\trouter.Use(middle.CORS())\n\n\tstore.SetClient(factory)\n\treturn router, nil\n}")
	return nil
}

// DockerGenerator 生成dockers容器相关目录和文件
//TODO: DockerGenerator
func DockerGenerator(name string, port string) error {
	//生成docker
	err := os.MkdirAll("api/docker/"+name, os.ModePerm)
	if err != nil {
		return err
	}
	//生成docker文件
	fcd, err := os.Create("api/docker/" + name + "/Dockerfile")
	if err != nil {
		return err
	}
	fcd.WriteString("FROM golang:1.18-alpine AS builder\n\n# 设置工作目录\nWORKDIR /app\n\nENV GO111MODULE=on  GOPROXY=https://goproxy.cn,direct GIN_MODE=release\n\nRUN cd\n\nCOPY . /app\n\nRUN go mod tidy\n\nRUN cd /app/cmd/" + name + " \\\n    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o " + name + "\n# 获取 Distroless 镜像，只有 650 kB 的大小，是常用的 alpine:latest 的 1/4\nFROM alpine\n\n# 将上一阶段构建好的二进制文件复制到本阶段中\nCOPY --from=builder /app/cmd/" + name + "/ .\nCOPY --from=builder /app/internal/config/ ./internal/config\nCOPY --from=builder /app/log ./log\n\n# 设置监听端口\nEXPOSE " + port + "\n# 配置启动命令\nENTRYPOINT [\"./" + name + "\"]")
	return nil
}

// SwaggerGenerator 生成swagger相关目录和文件
//TODO: SwaggerGenerator
func SwaggerGenerator(name string, port string) error {
	//生成swagger
	err := os.MkdirAll("api/swagger/"+name+"/doc", os.ModePerm)
	if err != nil {
		return err
	}
	//生成docker文件
	fcd, err := os.Create("api/swagger/" + name + "/doc/doc.go")
	if err != nil {
		return err
	}
	fcd.WriteString("//Package docs " + name + "\n//\n// " + name + "接口文档\n// Schemes: http, https\n// Version: 1.0.0\n// BasePath: /\n// Host: localhost:" + port + "\n// Consumes:\n// - application/x-www-form-urlencoded\n// Produces:\n// - application/json\n// Security:\n// - api_key:\n// SecurityDefinitions:\n//  api_key:\n//   type: apiKey\n//   in: header\n//   name: Authorization\n// swagger:meta\npackage docs")
	return nil
}

// CmdGenerator 生成cmd目录
// TODO: CmdGenerator
func CmdGenerator(name string, port string) error {
	path := getCurrentPath()
	//生成cmd文件
	err := os.MkdirAll("cmd/"+name, os.ModePerm)
	if err != nil {
		return err
	}
	fcc, err := os.Create("cmd/" + name + "/main.go")
	if err != nil {
		return err
	}
	fcc.WriteString("package main\n\nimport (\n\t\"" + path + "/internal/api/" + name + "\"\n\t\"" + path + "/internal/api/store\"\n\t\"" + path + "/internal/config\"\n)\n\nfunc main() {\n\tport := \"" + port + "\"\n\t//初始化代码配置\n\t_, err := config.InitConfig()\n\tif err != nil {\n\t\tpanic(\"配置初始化失败\" + err.Error())\n\t}\n\t//初始化数据层\n\tfactory, err := store.GetFactory()\n\tif err != nil {\n\t\tpanic(\"数据层初始化失败\" + err.Error())\n\t}\n\t//初始化路由\n\trouter, err := " + name + ".RouterInit(factory)\n\terr = router.Run(\":\"+port)\n\tif err != nil {\n\t\tpanic(err)\n\t}\n}\n")
	return nil
}

func FactoryGenerator(name string) error {
	fc, err := os.Create("internal/store/" + name + ".go")
	if err != nil {
		return err
	}
	big := firstUpper(name) + "Store"
	small := strings.ToLower(name[:1]) + name[1:] + "Store"
	fc.WriteString("package store\n\ntype " + big + " interface {\n}\n\ntype " + small + " struct {\n}\n\nfunc New" + big + "(ds *datastore) " + big + " {\n\treturn &" + small + "{}\n}\n")
	return nil
}

func ServiceGenerator(name string) error {
	path := getCurrentPath()
	if path != "service" {
		return errors.New("请进入service目录执行此命令")
	}
	big := firstUpper(name) + "Service"
	small := strings.ToLower(name[:1]) + name[1:] + "Service"
	fc, err := os.Create("./" + name + ".go")
	fc.WriteString("package service\n\ntype " + big + " interface {\n}\n\ntype " + small + " struct {\n}\n\nfunc New" + big + "(s *service) " + big + " {\n\treturn &" + small + "{}\n}\n")
	if err != nil {
		return err
	}
	return nil
}

// RunCommand 执行初始化命令
//TODO: RunCommand
func RunCommand() error {
	//初始化go.mod
	getwd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := getCurrentPath()
	if _, err = os.Stat(getwd + "/go.mod"); err != nil {
		if os.IsNotExist(err) {
			command := exec.Command("go", "mod", "init", path)
			err = command.Run()
			if err != nil {
				return err
			}
		}
	}
	//获取初始包
	command := exec.Command("go", "get", "github.com/golang-collections/collections")
	err = command.Run()
	if err != nil {
		return err
	}
	//获取gin框架
	command = exec.Command("go", "get", "github.com/gin-gonic/gin")
	err = command.Run()
	if err != nil {
		return err
	}
	//获取jwt
	command = exec.Command("go", "get", "github.com/dgrijalva/jwt-go")
	err = command.Run()
	if err != nil {
		return err
	}
	return nil
}

// ChangeGit 修改git源
// TODO: ChangeGit
func ChangeGit(url string) error {
	command := exec.Command("git", "remote", "rm", "origin")
	err := command.Run()
	if err != nil {
		return err
	}
	command = exec.Command("git", "remote", "add", "origin", url)
	err = command.Run()
	if err != nil {
		return err
	}
	return nil
}

func getCurrentPath() string {
	getwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	split := strings.Split(getwd, "\\")
	if len(split) == 0 {
		return ""
	}
	return split[len(split)-1]
}

func firstUpper(s string) string {
	var res string
	if s == "" {
		return ""
	}
	flag := strings.Contains(s, "_")
	if flag {
		split := strings.Split(s, "_")
		for _, str := range split {
			res += strings.ToUpper(str[:1]) + str[1:]
		}
	} else {
		res = strings.ToUpper(s[:1]) + s[1:]
	}
	return res
}
