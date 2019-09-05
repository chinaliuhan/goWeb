package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	//获得到参数,[0]为文件路径,[1]为输入的第一个参数,所以如果要带参的话,数组最小也应该是2个元素
	if len(os.Args) != 2 {
		fmt.Println("请输入参数")
		os.Exit(1)
	}

	name := os.Args[1]
	//将IP换转为IP类型,一般我们主要用来判断该IP地址是否是一个格式合法的IP地址
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("IP地址有误")
	} else {
		fmt.Println("输入的IP地址是:", addr.String())
		fmt.Printf("%T - %v", addr, addr)
	}

	os.Exit(0)
}
