package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "we are chinese"

	//编译正则表达式
	re, _ := regexp.Compile("[a-z]{2,4}")

	//查找第一个匹配的
	one := re.Find([]byte(str))
	fmt.Println(string(one)) //如果不用string()会输出utf8编号,find正则返回的都是utf8编号

	//返回一个多维数组,查找符合正则的所有的slice,n小于0标识返回全部符合条件的内容的一个slice,否则返回指定的长度
	all := re.FindAll([]byte(str), -1)
	fmt.Println(all) //因为返回的时一个slice,所以不能用string()了,想转换到话,需要遍历或者用下标

	//查找符合条件的内容的的index下标,返回开始位置和结束位置的一个slice
	index := re.FindIndex([]byte(str))
	fmt.Println(index)

	//返回一个多维数组,查找符合条件的所有内容的下标,返回开始和结束位置的一个slice
	//n小于0标识返回全部符合条件的内容的一个slice,否则返回指定的长度
	allIndex := re.FindAllIndex([]byte(str), -1)
	fmt.Println(allIndex)

	//重新编译一个新的正则表达式
	re1, _ := regexp.Compile("we(.*)ch(.*)")
	//返回一个多维数组，第一个元素是匹配的全部元素，第二个元素是第一个()里面的，第三个是第二个()里面的
	submatch := re1.FindSubmatch([]byte(str))
	fmt.Println(submatch)

	//查找符合条件的内容的子匹配的的index下标,不过这里是匹配()内的,返回开始位置和结束位置的一个slice
	submatchindex := re1.FindSubmatchIndex([]byte(str))
	fmt.Println(submatchindex)

	//查找所有符合条件的子匹配,返回一个多维数组
	//n小于0标识返回全部符合条件的内容的一个slice,否则返回指定的长度
	submatchall := re1.FindAllSubmatch([]byte(str), -1)
	fmt.Println(submatchall)

	//查找所有符合条件的子匹配的index下标,返回一个多维数组
	//n小于0标识返回全部符合条件的内容的一个slice,否则返回指定的长度
	submatchallindex := re1.FindAllSubmatchIndex([]byte(str), -1)
	fmt.Println(submatchallindex)
}
