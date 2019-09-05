package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析URL传递的参数,对于POST则解析请求包的主图,request body
	//注意,如果没有调用parseform,方法,下面将无法获取表单数据
	fmt.Println(r.Form) //打印提交的数据, r.Form里面包含了所有请求的参数
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["name"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value", strings.Join(v, ""))
	}
	//将数据打印到客户端页面
	fmt.Fprintf(w, "Hellow fucker")
}

//简单的登录操作
//主要教我们怎么获取静态页面
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method) //获取请求的方法
	//判断是否是GET请求
	if r.Method == "GET" {
		//这个应该是读取静态页面
		t, err := template.ParseFiles("login.gtpl")
		fmt.Println(err)
		//打印页面内容
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm() //解析URL传递的参数,对于POST则解析请求包的主图,request body
		//请求的时登录数据, 所以执行登录的逻辑判断
		//r.Form里面包含了所有请求的参数
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

//防止输入框重复提交
//这个函数,主要是教我们怎么向静态页面打印变量
func loginMul(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		//获取静态变量
		t, _ := template.ParseFiles("loginMul.gtpl")
		//想页面打印变量
		t.Execute(w, token)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}

//文件上传
func upload(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("this upload action")
	fmt.Println(r.Method) //打印请求方式
	//判断请求方式
	if r.Method == "GET" { //get请求
		//表示要获取静态文件
		t, err := template.ParseFiles("upload.gtpl")
		fmt.Println(err)
		//创建md5
		md5Hash := md5.New()
		//获取当前Unix时间戳
		date := time.Now().Unix();
		//拼装md5值
		io.WriteString(md5Hash, strconv.FormatInt(date, 10))
		token := fmt.Sprintf("%x", md5Hash.Sum(nil))

		t.Execute(w, token)

	} else { //其他请求
		//文件上传操作
		r.ParseMultipartForm(2048)
		//通过文件名,获取上传的文件,和php的$_FILE 一样爽
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			panic(err)
		}
		//关闭资源
		defer file.Close()
		//打印header
		fmt.Fprintf(w, "%v", handler.Header)
		second := time.Now().Nanosecond()
		fileName := strconv.Itoa(second)
		//openFile和open的地方在于open只能用来读取文件,handler.Filename获取文件名
		//os.O_WRONLY | os.O_CREATE | O_EXCL           【如果已经存在，则失败】
		//os.O_WRONLY | os.O_CREATE                         【如果已经存在，会覆盖写，不会清空原来的文件，而是从头直接覆盖写】
		//os.O_WRONLY | os.O_CREATE | os.O_APPEND  【如果已经存在，则在尾部添加写】
		//这里这个handler.Filename,会获取到包括路径在内的文件名,所以我改成了时间
		f, err := os.OpenFile("./uploads/"+fileName+".png", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	//调用设置路由,放置回调函数
	http.HandleFunc("/", sayHelloName)      //设置访问的路由
	http.HandleFunc("/login", login)        //设置访问的路由
	http.HandleFunc("/loginmul", loginMul)  //设置访问的路由
	http.HandleFunc("/upload", upload)      //设置访问的路由
	err := http.ListenAndServe(":999", nil) //设置监听的端口
	if err != nil {
		log.Fatal("listenAndServer:", err)
	}
}
