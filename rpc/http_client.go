package main

import (
	"os"
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("请输入参数")
		os.Exit(1)
	}

	//获取第一个参数作为请求地址
	serverAddress := os.Args[1]

	//发送请求,通过TCP的形式
	client, err := rpc.DialHTTP("tcp", serverAddress+":999")
	if err != nil {
		panic(err)
	}

	//设置请求参数
	args := Args{8, 2}

	//声明一个变量,然后使用指针形式调用该变量作为返回值的赋值
	var reply int
	//调用远程的RPC服务
	err = client.Call("Arith.Mul", args, &reply)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	//调用远程的RPC服务
	err = client.Call("Arith.Div", args, &quot)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

	//打印结果
	//~/go/src/web/rpc on  master! ⌚ 15:07:25
	//$ go run ./http_client.go localhost
	//Arith: 8*2=16
	//Arith: 8/2=4 remainder 0

}
