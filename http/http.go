package http

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"net/http"
	"strings"
)

//这两个参数w和r,不用刻意追求为什么这么写,因为外面的http.HandleFunc需要传入的一个回调函数
func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                         //解析参数,默认是不会解析的,如果不调用该函数解析的话,后面的参数都会拿不到
	fmt.Println("path: ", r.URL.Path)     //打印URL中的path
	fmt.Println("scheme: ", r.URL.Scheme) //打印http请求中的协议部分
	fmt.Println(r.Form["name"])           //打印get参数中的name参数
	for k, v := range r.Form { //遍历get参数表单,打印表单中的内容
		fmt.Println("query-string-key: ", k)
		fmt.Println("query-string-val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hellow fucker") //写入到w的时输出到客户端
}

func fakeServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //调用解析请求参数
	fmt.Println("host: ", r.URL.Host)
	fmt.Fprintf(w, "fake server")

}
func main() {
	http.HandleFunc("/", sayHelloName)      //设置访问的路由为/,并且调用回调函数
	http.HandleFunc("/fake", fakeServer)    //设置访问的路由为/,并且调用回调函数
	err := http.ListenAndServe(":999", nil) //设置监听的端口

	if err != nil {
		log.Fatal("ListenAndService", err)
	}
}
