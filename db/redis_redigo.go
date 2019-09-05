package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"reflect"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

var conn redis.Conn

func init() {
	//连接到Redis,网络连接方式,连接地址,选择库号
	connect, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(3))
	checkErr(err)
	conn = connect
}

func set(key string, value string) bool {
	//连接到Redis,网络连接方式,连接地址,选择库号
	//conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(3))
	//checkErr(err)
	//设置数据
	_, err := conn.Do("set", key, value)
	checkErr(err)

	//关闭连接
	return true
}

func get(key string) string {
	//连接到Redis
	//conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(3))
	//checkErr(err)
	//获取数据
	result, err := redis.String(conn.Do("get", key))
	checkErr(err)
	//关闭连接

	reflect.TypeOf(result)

	return result
}

func expire(key string, expire int) bool {
	//conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(3))
	//checkErr(err)
	result, err := conn.Do("expire", key, expire)
	checkErr(err)
	//这里因为result默认是一个interface,所以下面用了断言,来判断是不是一个Int
	//但是当result为Int的时候,返回的是一个int64
	value, ok := result.(int64)
	fmt.Println(ok)
	if ok {
		if value <= 0 {
			return false
		}
	}
	reflect.TypeOf(result)
	return true
}

func mGet(keys ...interface{}) {
	fmt.Println(reflect.TypeOf(keys))
	//这里的这个Keys后面必须要有..., 这三个点,代表将keys打散传入,没有三个点,代表传入一个集合
	//比如下面这个args{},如果要当做多个参数同时传入,就必须要在后面加...否则就是一个参数
	//var args= []interface{}{
	//	"name",
	//	"sex",
	//}
	result, err := redis.Strings(conn.Do("mget", keys...))
	checkErr(err)

	res_type := reflect.TypeOf(result)
	fmt.Printf("res type : %s \n", res_type)
	fmt.Printf("MGET name: %s \n", result)
	//return result
}

func lpush(values ...interface{}) {
	_, err := conn.Do("lpush", values)
	checkErr(err)

	fmt.Println("lpush ok")
}

func lpop(key string) {
	result, err := redis.String(conn.Do("lpop", key))
	checkErr(err)

	fmt.Printf("%s", result)
	fmt.Println("lpop ok")
	reflect.TypeOf(result)
}

func hset() {
	_, err := conn.Do("hset", "student", "name", "wd", "age", 22)
	checkErr(err)
	fmt.Println("hset ok")

}

func hget() {
	result, err := redis.Int64(conn.Do("hget", "student", "age"))
	checkErr(err)
	fmt.Printf("%s", result)

	fmt.Println(reflect.TypeOf(result))
}

//管道
//管道操作可以理解为并发操作，并通过Send()，Flush()，Receive()三个方法实现。
// 客户端可以使用send()方法一次性向服务器发送一个或多个命令，
// 命令发送完毕时，使用flush()方法将缓冲区的命令输入一次性发送到服务器，客户端再使用Receive()方法依次按照先进先出的顺序读取所有命令操作结果。
//Send：发送命令至缓冲区
//Flush：清空缓冲区，将命令一次性发送至服务器
//Recevie：依次读取服务器响应结果，当读取的命令未响应时，该操作会阻塞。

func pipelin() {
	conn.Send("HSET", "student", "name", "wd", "age", "22")
	conn.Send("HSET", "student", "Score", "100")
	conn.Send("HGET", "student", "age")
	conn.Flush()

	res1, err := conn.Receive()
	checkErr(err)
	fmt.Printf("Receive res1:%v \n", res1)
	res2, err := conn.Receive()
	fmt.Printf("Receive res2:%v\n", res2)
	res3, err := conn.Receive()
	fmt.Printf("Receive res3:%s\n", res3)
}
func main() {
	defer conn.Close()

	println(set("name", "ssss"))
	println(set("sex", "1111"))
	//println(get("name"))
	//println(expire("name", 2))
	//mGet("name", "sex")

}
