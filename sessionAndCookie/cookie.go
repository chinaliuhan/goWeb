package main

import (
	"fmt"
	"net/http"
	"time"
)

//设置cookie
func setCookie(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var name string
	var value string
	if len(r.Form["name"]) > 0 {
		name = r.Form["name"][0]
	}
	if len(r.Form["value"]) > 0 {
		value = r.Form["value"][0]
	}
	fmt.Printf("%v---%T\n", name, name)
	fmt.Printf("%v---%T\n", value, value)

	//获取当前时间格式 2017-07-26 15:32:04.251666 +0800 CST m=+5.348925672
	expiration := time.Now()
	fmt.Println(expiration)
	//在原来的基础上增加一年,这个连在一起写就明白了expiration := time.Now().AddDate(1, 0, 0)
	expiration = expiration.AddDate(1, 0, 0)
	fmt.Println(expiration)


	cookie := http.Cookie{Name: name, Value: value, Expires: expiration}
	http.SetCookie(w, &cookie)
	fmt.Fprint(w, "this is set cookie")
}

//获取cookie
func getCookie(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var name string
	if len(r.Form["name"]) > 0 {
		name = r.Form["name"][0]
	}
	fmt.Println(name)
	cookie, _ := r.Cookie(name)

	fmt.Fprint(w, "get cookie", "\n", cookie)
}

//获取所有的cookie
func getTotalCookie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "get total cookie\n")

	for _, v := range r.Cookies() {

		fmt.Fprint(w, v.Name, "---", v.Value)
	}

}

//设置session
func setSession(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "this is set session")
}

//获取session
func getSession(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "this is get session")

}

func main() {
	http.HandleFunc("/setcookie", setCookie)
	http.HandleFunc("/getcookie", getCookie)
	http.HandleFunc("/gettotalcookie", getTotalCookie)

	http.HandleFunc("/setsession", setSession)
	http.HandleFunc("/getsession", getSession)

	err := http.ListenAndServe(":999", nil)
	if err != nil {
		panic(err)
	}

}
