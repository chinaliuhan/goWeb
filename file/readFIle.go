package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "nihao.txt"
	//打开文件
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	//关闭资源
	defer fl.Close()
	//开辟内存空间
	buf := make([]byte, 1024)
	//死循环
	for {
		//按照开辟的内存空间,循环读取文件,读到文件结束的时候会返回0
		n, _ := fl.Read(buf)
		if 0 == n {
			//退出循环
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
