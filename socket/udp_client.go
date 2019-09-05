package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	//通过传入IP地址,得到UDPAddr类型的UDP IP信息
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	//建立一个UDP连接,返回UDPConn类型的对象
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	//发送信息
	_, err = conn.Write([]byte("anything"))
	checkError(err)
	//创建内存空间
	var buf [512]byte
	//读取UDP连接中的数据
	n, err := conn.Read(buf[0:])
	checkError(err)
	//打印数据
	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
