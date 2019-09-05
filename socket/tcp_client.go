package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	//验证输入的数据长度
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "错误的参数: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	//得到参数
	service := os.Args[1]
	//处理链接,得到一个TCPAddr类型的tcp ip地址信息
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	//创建连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	//发送数据
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	//读取通道中的信息,这样不好,应该用Listen和Accept来做接收
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	//打印数据
	fmt.Println(string(result))
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
