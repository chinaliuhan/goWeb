package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

//这里来一个简单的小爬虫
func main() {

	//请求页面
	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	//关闭资源
	defer response.Body.Close()
	//获取页面内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	//将body转换为字符串
	content := string(body)

	//将所有的HTML转换为大写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	//ReplaceAllStringFunc可以接受一个回调函数进去
	src := re.ReplaceAllStringFunc(content, strings.ToUpper)
	fmt.Println(src)

	//去掉script
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	//ReplaceAllString不接收回调函数,只能纯粹的替换
	src1 := re.ReplaceAllString(content, "")
	fmt.Println(src1)

	//去掉style
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src2 := re.ReplaceAllString(content, "")
	fmt.Println(src2)

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src3 := re.ReplaceAllString(content, "\n")
	fmt.Println(src3)

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src4 := re.ReplaceAllString(content, "\n")
	fmt.Println(src4)

}
