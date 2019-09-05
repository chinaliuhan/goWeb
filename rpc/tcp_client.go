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

	//判断参数
	if len(os.Args) != 2 {
		fmt.Println("请输入URL地址和端口号,格式 host:port")
	}

	//获取参数,发送连接
	service := os.Args[1]
	client, err := rpc.Dial("tcp", service)
	if err != nil {
		panic(err)
	}

	//填充参数, 调用远程RPC服务,进行乘法操作
	args := Args{4, 2}
	var reply int
	err = client.Call("Arith.Mul", args, &reply)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	//声明变量,调用远程RPC服务,进行除法操作
	var quot Quotient
	err = client.Call("Arith.Div", args, &quot)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
