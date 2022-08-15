package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/zhangshuai268/spg-go-framework/generator"
	"os"
	"strconv"
)

var opts struct {
	Generator bool `short:"g"`
	Version   bool `short:"v" long:"version"`
	Update    bool `short:"u" long:"update"`
	Service   bool `short:"s" long:"service"`
	Store     bool `short:"f" long:"factory"`
}

func main() {
	getwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = flags.ParseArgs(&opts, os.Args)
	if opts.Version {
		fmt.Println("v1.1.2")
	} else if opts.Update {
		generator.GenerateModels(getwd+"/config.json", getwd+"/internal/config/config.go")
	} else if opts.Generator {
		//初始化go.mod
		err := generator.RunCommand()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("初始化go.mod,下载相关依赖")
		err = generator.FrameworkGenerator()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("生成目录结构成功")
		err = generator.ConfigGenerator()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("生成配置文件成功")
		err = generator.StoreGenerator()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("生成数据库目录成功")
		fmt.Println("请输入开发系统数量：")
		var n string
		_, _ = fmt.Scanln(&n)
		var atoi int
		for {
			atoi, err = strconv.Atoi(n)
			if err == nil {
				break
			} else {
				fmt.Println("请输入数字：")
				_, _ = fmt.Scanln(&n)
			}
		}
		for i := 0; i < atoi; i++ {
			var name, port string
			fmt.Println("请输入开发系统名称：")
			_, _ = fmt.Scanln(&name)
			for {
				letter := isLetter(name)
				if letter {
					break
				} else {
					fmt.Println("系统名称不合法：")
					_, _ = fmt.Scanln(&name)
				}
			}
			fmt.Println("请输入开发系统运行端口号：")
			_, _ = fmt.Scanln(&port)
			for {
				_, err = strconv.Atoi(port)
				if err == nil {
					break
				} else {
					fmt.Println("请输入正确端口号：")
					_, _ = fmt.Scanln(&port)
				}
			}
			//docker文件生成
			err = generator.DockerGenerator(name, port)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			//swagger文件生成
			err = generator.SwaggerGenerator(name, port)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			//主要代码目录文件生成
			err = generator.InternalGenerator(name)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			//启动代码文件生成
			err = generator.CmdGenerator(name, port)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	} else if opts.Store {
		fmt.Println("请输入store名称：")
		var name string
		_, _ = fmt.Scanln(&name)
		for {
			letter := isStore(name)
			if letter {
				break
			} else {
				fmt.Println("系统名称不合法：")
				_, _ = fmt.Scanln(&name)
			}
		}
		err := generator.FactoryGenerator(name)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else if opts.Service {
		fmt.Println("请输入service名称：")
		var name string
		_, _ = fmt.Scanln(&name)
		for {
			letter := isStore(name)
			if letter {
				break
			} else {
				fmt.Println("系统名称不合法：")
				_, _ = fmt.Scanln(&name)
			}
		}
		err := generator.ServiceGenerator(name)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

}

func isLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && r != '-' && r != '_' {
			return false
		}
	}
	return true
}

func isStore(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && r != '_' {
			return false
		}
	}
	return true
}
