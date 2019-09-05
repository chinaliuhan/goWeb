package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect error :", err)
		return
	}
	defer conn.Close()
	conn.Send("MULTI")
	conn.Send("INCR", "aa")
	conn.Send("INCR", "aa")
	r, err := conn.Do("EXEC")
	fmt.Println(r)
}

