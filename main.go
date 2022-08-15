package main

import (
	"fmt"
	"github.com/zhangshuai268/spg-go-framework/generator"
	"os"
	"strconv"
	"strings"
)

func main() {
	//修改git源
	fmt.Println("请输入git仓库地址：")
	var url string
	_, _ = fmt.Scanln(&url)
	for {
		index := strings.Index(url, ".git")
		if index == -1 {
			fmt.Println("请输入正确的git仓库地址：")
			_, _ = fmt.Scanln(&url)
		} else {
			break
		}
	}
	err := generator.ChangeGit(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	//初始化go.mod
	err = generator.RunCommand()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("初始化go.mod,下载相关依赖")
	err = generator.FrameworkGenerator()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("生成目录结构成功")
	err = generator.ConfigGenerator()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("生成配置文件成功")
	err = generator.StoreGenerator()
	if err != nil {
		fmt.Println(err)
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
			fmt.Println(err)
			continue
		}
		//swagger文件生成
		err = generator.SwaggerGenerator(name, port)
		if err != nil {
			fmt.Println(err)
			return
		}
		//主要代码目录文件生成
		err = generator.InternalGenerator(name)
		if err != nil {
			fmt.Println(err)
			return
		}
		//启动代码文件生成
		err = generator.CmdGenerator(name, port)
		if err != nil {
			return
		}
	}
	_ = os.Remove("./framework_create.exe")
	_ = os.Remove("./README.md")
}

func isLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && r != '-' && r != '_' {
			return false
		}
	}
	return true
}
