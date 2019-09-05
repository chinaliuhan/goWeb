package main

import (
	"strconv"
	"fmt"
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {

	//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	fmt.Println("Append函数:-----------")
	//创建数组
	str := make([]byte, 0, 100)

	//[]byte中添加int类型（int-->[]byte）,值是4567,10进制,返回的是utf8编号,显示的时候需要string()一下
	str = strconv.AppendInt(str, 4567, 10)
	fmt.Println(string(str))

	// []byte中添加bool类型 （bool-->[]byte）
	str = strconv.AppendBool(str, false)
	fmt.Println(string(str))

	// []byte中添加string类型(包含双引号)  （string-->[]byte）
	str = strconv.AppendQuote(str, "ssss")
	fmt.Println(string(str))
	//AppendQuoteRune 将 Unicode 字符转换为“单引号”引起来的字符串，
	//并将结果追加到 dst 的尾部，返回追加后的 []byte
	//“特殊字符”将被转换为“转义字符”
	str = strconv.AppendQuoteRune(str, '哈')
	fmt.Println(str)

	fmt.Println("Format函数:-----------")

	//将布尔值转换为字符串 "true" 或 "false"
	a := strconv.FormatBool(false)
	//将浮点数123.23转换为字符串值
	// 123.23要转换的浮点数, g格式标记（b、e、E、f、g、G,), 12精度（数字部分的长度，不包括指数部分）,64指定浮点类型（32:float32、64:float64）
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	//将 int 型整数 1234 转换为字符串形式
	//10为进制
	c := strconv.FormatInt(1234, 10)
	//将 uint 型整数 12345 转换为字符串形式
	//10为进制
	d := strconv.FormatUint(12345, 10)
	//Itoa 相当于 FormatInt(i, 10), Itoa()仅限十进制
	e := strconv.Itoa(1024)

	fmt.Println(a, b, c, d, e)

	fmt.Println("Parse函数:-----------")

	//将字符串转换为布尔值
	aa, err := strconv.ParseBool("false")
	checkError(err)
	//ParseFloat 将字符串转换为浮点数
	// 将123.12转换为浮点型
	// 64：指定浮点类型（32:float32、64:float64）
	bb, err := strconv.ParseFloat("123.12", 64)
	checkError(err)

	// ParseInt 将字符串转换为 int 类型
	// s：要转换的字符串
	// base：进位制（2 进制到 36 进制）
	// bitSize：指定整数类型（0:int、8:int8、16:int16、32:int32、64:int64）
	cc, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	// ParseUint 功能同 ParseInt 一样，只不过返回 uint 类型整数
	dd, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	// Atoi 相当于 ParseInt(s, 10, 0)
	ee, err := strconv.Atoi("1024")
	checkError(err)

	fmt.Println(aa, bb, cc, dd, ee)

}
