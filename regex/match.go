package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	//判断命令行输入的参数,需要通过这种方式运行,go run ./match.go 1111 不能用IDE直接运行,否则会没有参数
	if len(os.Args) == 1 {
		fmt.Println("请输入参数")
		os.Exit(1)
	} else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("数字")
	} else {
		fmt.Println("不是数字")
	}

	////MatchString判断输入的字符串是否符合标准
	////var ip string = "127.0.0.1"
	//var ip string = "127.0.0.1.11"
	//if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
	//	fmt.Println("IP地址有误")
	//} else {
	//	fmt.Println("IP地址正确")
	//}
}
