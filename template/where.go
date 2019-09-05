package main

import (
	"os"
	"text/template"
)

type Where struct {
	Any bool
}

func main() {

	//为ture,下面显示if 部分,为false,下面显示else部分
	whe := Where{Any: true}
	//新建一个模板
	tIfElse := template.New("template test")
	//这里的这个条件判断就不说了,一看就明白
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if .Any}} if部分 {{else}} else部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, whe)
}
