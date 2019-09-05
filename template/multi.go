package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	//将三个文件同时加载进去
	s1, _ := template.ParseFiles("template/header.html", "template/header.html", "template/header.html")
	//打印header文件
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	//打印内容文件
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	//打印注脚文件
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)
}
