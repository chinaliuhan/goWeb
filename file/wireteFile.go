package main

import (
	"fmt"
	"os"
)

func newFile() {

	fout := os.NewFile(3, "haha.txt") //输如ini.go的句柄
	//关闭资源
	defer fout.Close()
	//循环写入内容
	for i := 0; i < 10; i++ {
		//写入内容
		fout.WriteString("sss\n")
		//注意这个和上面的不同,这是一个byte
		fout.Write([]byte("ssss\n"))
	}
}

func main() {
	newFile()
	return
	userFile := "nihao.txt"
	//直接创建文件
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	//关闭资源
	defer fout.Close()
	//循环写入内容
	for i := 0; i < 10; i++ {
		//写入内容
		fout.WriteString("sss\n")
		//注意这个和上面的不同,这是一个byte
		fout.Write([]byte("ssss\n"))
	}
}
