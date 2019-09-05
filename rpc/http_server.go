package main

import (
	"github.com/pkg/errors"
	"net/rpc"
	"net/http"
	"fmt"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

//数学计算
type Arith int

//乘法
func (t *Arith) Mul(args *Args, reply *int) error {
	*reply = args.A * args.B

	return nil

}

//除法
func (t *Arith) Div(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("Divide参数为0")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil

}
func main() {
	//为该结构体开辟内存空间
	arith := new(Arith)
	//注册RPC
	rpc.Register(arith)
	//调用RPC服务
	rpc.HandleHTTP()
	//开启服务
	err := http.ListenAndServe(":999", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
