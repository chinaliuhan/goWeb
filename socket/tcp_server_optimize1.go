package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	service := "172.16.8.7:1200"
	//设置IP信息,得到tcp ip地址的TCPAddr格式的信息
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	//监听端口
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	//死循环,不停的监听端口
	for {
		//接收端口传来的数据
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		//使用goroutine处理业务逻辑的函数
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	//设置超时时间
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	//设置最大请求数据的长度,防止洪水攻击
	request := make([]byte, 128)
	defer conn.Close()
	clientDate := "这里是客户端发送的数据"
	for {
		//从连接中读取内容
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}

		//如果返回内容时的长度是0,退出进程
		if read_len == 0 {
			fmt.Println("请求的数据为空")
			break

			//如果返回时间,清除两端的空格截取0-最后的内容转换为字符串后==timestamp这个字符串,那么格式化时间,返回时间
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			//否则直接发送当前时间戳
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}

		fmt.Println("客户端请求的数据长度为:", read_len)

		//清空最后一次读取的内容,好让新的数据写入进去
		request = make([]byte, 128)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
