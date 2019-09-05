package main

import (
	"fmt"
	"strings"
)

func main() {

	//判断字符串是否在另一个字符串中，返回bool值
	fmt.Println(strings.Contains("we are chinese", "chin")) //true
	fmt.Println(strings.Contains("we are chinese", "ss"))   //false
	fmt.Println(strings.Contains("we are chinese", ""))     //true

	//字符串连接,将slice通过指定符号拼接成字符串
	s := []string{"aaa", "sss", "ddd"}
	fmt.Println(strings.Join(s, ",")) //aaa,sss,ddd

	//在字符串中查找s指定字符串所在的位置，返回位置值，找不到返回-1
	fmt.Println(strings.Index("we are chinese", "ch")) //7

	//重复生成字符串sss 3次,返回生成后的字符串
	fmt.Println("aaa" + strings.Repeat("sss", 3))

	//在字符串中，把we替换为We，2表示替换的次数如果小于0表示全部替换(类似于正则中的是否贪婪替换)
	fmt.Println(strings.Replace("we we are chinese", "we", "We", 1)) //We are chinese
	fmt.Println(strings.Replace("we we are chinese", "we", "We", 2)) //We We are chinese
	fmt.Println(strings.Replace("we are chinese", "ch", "Ch", -1))   //we are Chinese

	//把字符串按照指定符号进行分割成slice,如果指定符号为""会将所有字符全部指定为slice的一个元素,包括空格字啊内
	fmt.Printf("%q\n", strings.Split("a,b,c", ",")) //["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split(" abc ", ""))  //[" " "a" "b" "c" " "]
	fmt.Printf("%q\n", strings.Split("", "haha"))   //[""]

	//在字符串首尾,去掉指定的字符
	fmt.Println(strings.Trim("-hahah nihao woshi -", "-")) //hahah nihao woshi

	//去掉字符串中的所有空格,并按照空格分割,返回slice
	fmt.Println(strings.Fields("nihao wo shi ZhonGUo ren nine ")) //[nihao wo shi ZhonGUo ren nine]

}
