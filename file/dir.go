package main

import (
	"os"
)

func main() {
	//创建单个目录,不能创建多级目录
	//err := os.Mkdir("nihao/hahahha/", 0777)
	//if err != nil{
	//	panic(err)
	//}
	//可以创建多级目录
	err := os.MkdirAll("nihao/tahao/", 0777)
	if err != nil {
		panic(err)
	}
	//只能删除空目录
	//err := os.Remove("nihao")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//可以删除多级目录,即使有内容也一样可以删除
	err = os.RemoveAll("nihao")
	if err != nil {
		panic(err)
	}
}
