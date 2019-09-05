package main

import (
	"github.com/pkg/errors"
	"net/rpc"
	"net"
)

//声明三个结构体

//作为参数的结构体
type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

//声明一个结构体,作为对象使用
type Arith int

//乘法
func (t *Arith) Mul(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

//除法
func (t *Arith) Div(args *Args, quo *Quotient) error {

	if args.B == 0 {
		return errors.New("div 不能是0")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {

	//同样,需要为该结构体(相当于对象)开辟空间
	arith := new(Arith)
	//注册rpc
	rpc.Register(arith)
	//同样需要开启一个socket服务
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":999")
	if err != nil {
		panic(err)
	}
	//监听socket端口
	listenner, err := net.ListenTCP("tcp", tcpAddr)

	//死循环不断的处理
	for {
		//通过监听的端口获取到数据
		conn, err := listenner.Accept()
		if err != nil {
			continue
		}
		//通过rpc调用请求
		rpc.ServeConn(conn)
	}
}
