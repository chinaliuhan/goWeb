package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200"
	//处理IP地址,得到得到一个TCPAddr类型的tcp ip地址信息
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	//监听端口
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	//同样,死循环处理端口传来的数据
	for {
		//接收端口传来的数据
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//使用goroutine来用handleClient函数处理接收到的数据并返回数据
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	//随便生成的时间格式
	daytime := time.Now().String()
	//发送数据
	conn.Write([]byte(daytime))
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
