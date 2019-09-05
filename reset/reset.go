package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "你好!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//获取请求参数
	fmt.Fprintf(w, ps.ByName("name"))
}

func getuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//获取请求参数
	uid := ps.ByName("uid")
	fmt.Fprintf(w, uid)
}

func modifyuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//获取请求参数
	uid := ps.ByName("uid")
	fmt.Fprintf(w, uid)
}

func deleteuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//获取请求参数
	uid := ps.ByName("uid")
	fmt.Fprintf(w, uid)
}

func adduser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//获取请求参数
	uid := ps.ByName("uid")
	fmt.Fprintf(w, uid)
}

func main() {
	//新建路由
	router := httprouter.New()
	//根据不同的请求方式, 设置匹配的URL连接
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/user/:uid", getuser)
	router.POST("/adduser/:uid", adduser)
	router.DELETE("/deluser/:uid", deleteuser)
	router.PUT("/moduser/:uid", modifyuser)

	//监听服务,同时记录日志
	log.Fatal(http.ListenAndServe(":999", router))
}
