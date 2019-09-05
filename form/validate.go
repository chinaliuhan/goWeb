package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"time"
)

//日期时间
func ValidateDate() {
	//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time {} // 返回指定时间
	//time.November 为11月,这里正常的输出应该是2009-11-10 23:00:00,但是下面.local()了一下,获取的时当地时间
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	//折算成当地的时间,获取的时当地时间
	fmt.Printf(" %s\n", t.Local())
}

//身份证
func ValidateIdCard(idCard string) bool {
	//验证15位身份证，15位的是全部数字
	//if m, _ := regexp.MatchString(`^(\d{15})$`, idCard); !m {
	//	return false
	//}

	//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
	if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, idCard); !m {
		return false
	}

	return true
}

//电子邮件地址
func validateEmail(email string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return false
	} else {
		return true
	}
}

func validateXss(w http.ResponseWriter, r *http.Request) {
	var xssStr string = "<script>alert(111)</script>"
	fmt.Println("username:", template.HTMLEscapeString(xssStr)) //输出到服务器端
	fmt.Println("password:", template.HTMLEscapeString(xssStr))
	template.HTMLEscape(w, []byte(xssStr)) //直接输出到客户端浏览器页面
}

func validateXxxDecode() {
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
	//t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	//if err != nil {
	//	panic(err)
	//}
	////var out = new(bytes.Buffer)
	//err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
	////建立一个模板
	//t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	//if err != nil {
	//	panic(err)
	//}
	//err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
	//if err != nil {
	//	panic(err)
	//}
}

func main() {
	//ValidateDate()
	//println(ValidateIdCard("111111111111111111"))
	//println(validateEmail("nihao@nihao.com"))

	validateXxxDecode()

	//调用设置路由,放置回调函数
	//http.HandleFunc("/", validateXss)      //设置访问的路由
	//err := http.ListenAndServe(":999", nil) //设置监听的端口
	//if err != nil {
	//	log.Fatal("listenAndServer:", err)
	//}

}
