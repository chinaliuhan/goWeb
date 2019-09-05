package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200"
	//通过传入IP地址,得到UDPAddr类型的UDP IP信息
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	//监听端口,返回一个网络套接字,其实就是一个UDPConn类型的数据
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	//死循环,不见得监听
	for {
		//调用逻辑循环函数
		handleClient(conn)
	}
}
func handleClient(conn *net.UDPConn) {
	//开辟内容空间
	var buf [512]byte
	//从UDP连接中读取连接中的内容
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	daytime := time.Now().String()
	//想UDP发送数据
	conn.WriteToUDP([]byte(daytime), addr)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
