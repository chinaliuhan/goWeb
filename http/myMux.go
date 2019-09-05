package http

import (
	"fmt"
	"net/http"
)

//自定义类型
type MyMux struct {
}

//路由匹配方法,作为MyMux的方法,主要在下面传入函数中,覆盖系统的ServeHTTP用
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		SayHelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

//打印一串字符串
func SayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute")
}

func main() {
	//声明该自定义类型,同时取到该类型的内存地址
	mux := &MyMux{}
	//监听端口
	//将该类型,作为回调函数放入到方法中,相应的MyMux类型的ServeHTTP方法也会被传入进去,或者说是在http.ListenAndServe调用ServeHTTP方法时被调用
	//因为http类型里面也有一个ServeHTTP方法, 上面我们自己写的ServeHTTP是为了,在这类传入进去,覆盖系统的ServeHTTP方法
	http.ListenAndServe(":999", mux)
}
