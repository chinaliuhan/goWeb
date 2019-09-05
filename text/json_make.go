package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type ServerSlice struct {
	Servers []Server
}

type Server1 struct {
	//tag的作用
	//如果一个域不是以大写字母开头的，那么转换成json的时候，这个域是被忽略的。
	//如果没有使用json:"name"tag，那么输出的json字段名和域名是一样的。
	//字段的tag是"-"，那么这个字段不会输出到JSON
	ID int `json:"-"`

	// ServerName2 的值会进行二次JSON编码
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`

	//tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
	ServerIP string `json:"serverIP,omitempty"`
}

func main() {
	//为结构体赋值,该结构体中的变量Servers,将成为JSON的一个下标
	var s ServerSlice
	//向结构体追加新的数据,因为是对s.Servers进行append,所以新的数据将作为s.Servers的值
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	//再次追加数据
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	//将结构体生成为JSON
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	//输出结果{"Servers":[{"ServerName":"Shanghai_VPN","ServerIP":"127.0.0.1"},{"ServerName":"Beijing_VPN","ServerIP":"127.0.0.2"}]}

	//todo 上面的方式生成的JSON的数据只能是大写开头的,要想生成全小写必须要使用struct tag
	s1 := Server1{
		ID:          3,
		ServerName:  `Go "1.0" `,
		ServerName2: `Go "1.0" `,
		ServerIP:    ``,
	}
	b1, _ := json.Marshal(s1)
	os.Stdout.Write(b1)
	//{"serverName":"Go \"1.0\" ","serverName2":"\"Go \\\"1.0\\\" \""}

}
