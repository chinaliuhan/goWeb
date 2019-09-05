package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	//为结构体开辟空间
	arith := new(Arith)
	//注册RPC服务
	rpc.Register(arith)
	//激活TCP服务
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)
	//监听端口
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		//处理接收到的数据
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//将接收到的连接转发给RPC
		jsonrpc.ServeConn(conn)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
