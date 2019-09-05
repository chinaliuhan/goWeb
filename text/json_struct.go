package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type ServersLice struct {
	Servers []Server
}

func main() {

	//解析json到结构体中,这种情况需要事先知道JSON的所有字段
	var s ServersLice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	//该函数可以解析json字符串到变量s中,这里的变量s是一个结构体
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	fmt.Println(s.Servers[0].ServerName)
	fmt.Println(s.Servers[0].ServerIP)

}
