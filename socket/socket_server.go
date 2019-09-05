package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"net/http"
)

func Echo(ws *websocket.Conn) {
	var err error

	//死循环,不间断处理
	for {
		//定义一个变量,后面读取websocket内容的时候要用到他的指针
		var reply string

		//读取websocket的,replay为返回的内容
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("客户端发送的信息: " + reply)

		msg := "服务器接收到数据,这里是服务器返回的信息"
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	//通过一个http服务来做
	http.Handle("/", websocket.Handler(Echo))

	//开启服务
	if err := http.ListenAndServe(":999", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
