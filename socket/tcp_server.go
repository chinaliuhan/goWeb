package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":7777"
	//激活句柄,处理链接,得到一个TCPAddr类型的tcp ip地址信息
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	//监听端口
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	//死循环不断的处理端口传来的数据
	for {
		//接收通过端口接听到的数据
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//随便订的时间戳
		daytime := time.Now().String()
		//发送数据给客户端
		conn.Write([]byte(daytime))
		conn.Close()
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
