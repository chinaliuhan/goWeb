package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {

	//新建一个模板,这里的内容任意填
	t := template.New("hahah example")

	//拼装一个模板结构
	//因为下面会用t调用Execute来把模板结构和结构体拼装到一起
	//所以这里这个模板结构,就很容易明白是怎么回事了,先声明一个机构提,再拼装一个模板结构,用模板结构调用执行函数,将结构体和模板结构拼装
	// 结构已经很显然了,range .Emails代表遍历这个属性,{{.}代表在遍历时获取每一个字段,注意每一个range都需要一个end,{{.}}也要放到两者之间
	//而,下面的with只是另一种写法
	t, _ = t.Parse(`hello {{.UserName}}!
			{{range .Emails}}
				an email {{.}}
			{{end}}
			{{with .Friends}}
			{{range .}}
				my friend name is {{.Fname}}
			{{end}}
			{{end}}
			`)
	//向结构体中填充数据,将结构体补全, 最终将两个结构体合二为一
	//想结构体中的字段填充内容
	f1 := Friend{Fname: "nihao"}
	f2 := Friend{Fname: "wohao"}
	p := Person{
		UserName: "json",
		//这里这个emails是一个数组,所以展示的时候,就需要变遍历,需要再模板中使用range
		Emails: []string{"json@163.com", "ason@gmail.com"},
		//同样,这的friends也是一个数组
		Friends: []*Friend{&f1, &f2},
	}

	//执行模板混合, 将模板结构和结构体的数据拼装,同时打印数据
	//t.Execute指的就是,t是一个模板的结构,Execute就是拼装函数, 而p则是我们事先准备好的结构体
	t.Execute(os.Stdout, p)
}
