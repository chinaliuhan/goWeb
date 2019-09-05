package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	//如果不知道JSON内部的数据类型,就只能将JSON解析为接口了
	//因为interface{}是可存储任意类型的
	//Go类型和JSON类型的对应关系如下：
	//bool 代表 JSON booleans,
	//float64 代表 JSON numbers,
	//string 代表 JSON strings,
	//nil 代表 JSON null.
	//
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	//声明一个接口
	var f interface{}
	//将JSON解析到接口中,,所谓接口
	err := json.Unmarshal(b, &f)
	if err != nil {
		panic(err)
	}
	fmt.Println("打印F",f)
	//必须通过断言的方式将interface赋值,然后才能使用
	m := f.(map[string]interface{})
	//因为JSON是有很多类型,比如可能是int也有可能是字符串,如果直接用下标取值的话我们会因为不知道数据类型出错
	//所以我们通过下面的形式来遍历一遍
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is array")
			for i, j := range vv {
				fmt.Println(i, j)
			}
		default:
			fmt.Println(k, "不知道是什么类型")
		}
		//打印数据类型
		fmt.Println(reflect.TypeOf(m))
		//直接取值,可能会造成变量类型有误而出错,但是如果和对方商量好的话,是没问题的
		fmt.Println(m["Name"])
	}
}
