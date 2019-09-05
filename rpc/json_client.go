package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	//判断传入的参数个数
	if len(os.Args) != 2 {
		fmt.Println("请输入参数")
		log.Fatal(1)
	}

	//获取参数,调动RPC服务
	service := os.Args[1]
	client, err := jsonrpc.Dial("tcp", service)
	if err != nil {
		panic(err)
	}
	//设置提交参数
	args := Args{4, 2}

	//声明当做返回值的变量
	var reply int
	//取变量的内存地址传入函数,并使用函数调用远程RPC服务
	err = client.Call("Arith.Mul", args, &reply)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	//和上面一样,声明变量,将变量内存地址放入函数作为返回值用,同时使用函数调用远程RPC服务
	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
