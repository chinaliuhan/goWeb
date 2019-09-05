package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	password string
}

func main() {
	////创建一个模板
	t := template.New("fieldname example")
	//解析模板文件,t.Parse()处理的可以不是一个文件
	//要导出的字段必须是大写的,否则显示为空
	t, _ = t.Parse("hello {{.UserName}} {{.password}}!")
	p := Person{UserName: "this hello",password:"this password"}
	t.Execute(os.Stdout, p)

	//如果passowrd 为大写 输出  hello this hello
	//如果password为小写 输出 hello this hello this password!
}