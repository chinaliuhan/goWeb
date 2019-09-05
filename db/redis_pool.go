package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

var Pool redis.Pool

func init() {
	Pool = redis.Pool{
		MaxIdle:     16,  //最大空闲连接数
		MaxActive:   32,  //最大的激活连接数
		IdleTimeout: 120, //空闲连接等待时间
		//连接方法
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	//通过连接池发起连接
	conn := Pool.Get()
	//执行Redis命令
	res, err := conn.Do("set", "name", "nihao")
	//打印结果
	fmt.Println(res, err)
	//执行Redis命令
	result, err := redis.String(conn.Do("get", "name"))
	fmt.Printf("%v", result)

}

