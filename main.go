package main

import (
	"fmt"
	"spg-go-framework-exe/generator"
)

func main() {
	err := generator.FrameworkGenerator()
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
	fmt.Println("请输入开发系统数量：")
	var n string
	fmt.Scanln(&n)

}
