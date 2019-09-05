package main

import (
	"text/template"
	"os"
)

func main() {
	tOk := template.New("first")
	//一切正常,正常输出
	mustOk := template.Must(tOk.Parse(" some static text /* and a comment */"))
	mustOk.Execute(os.Stdout, nil)

	tErr := template.New("check parse error with Must")
	//这里就会提示一个panic: template: check parse error with Must:1: unexpected "}" in operand
	//只要在{{ .Name}后面再加上一个}就可以修复
	mustErr := template.Must(tErr.Parse(" some static text {{ .Name}"))

	mustErr.Execute(os.Stdout, nil)


}
