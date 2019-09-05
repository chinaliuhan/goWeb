package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

//上传函数
//下面的例子详细展示了客户端如何向服务器上传一个文件的例子，客户端通过multipart.Write把文件的文本流写入一个缓存中，然后调用http的Post方法把缓存传到服务器。
func uploadFile(filename string, targetUrl string) error {
	//创建控件
	bodyBuf := &bytes.Buffer{}
	//创建一个可写入资源,吐槽一下,multipart这个包,设计的太二笔了...
	bodyWrite := multipart.NewWriter(bodyBuf)

	//从文件中读取数据, 创建表单文件名
	fileWrite, err := bodyWrite.CreateFormFile("uploadfile", filename)
	if err != nil {
		return err
	}

	//打开文件,.open的形式打开,只能用作读取
	fh, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fh.Close()

	//ioCopy,从fh复制到fileWrite，直到到达EOF或发生错误,返回拷贝的字节喝遇到的第一个错误.
	_, err = io.Copy(fileWrite, fh)
	if err != nil {
		return err
	}

	//返回http的请求需要的类型
	contentType := bodyWrite.FormDataContentType()
	bodyWrite.Close()
	//开始上传
	response, err := http.Post(targetUrl, contentType, bodyBuf)
	panic(err)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Println(response.Status)
	fmt.Println(string(responseBody))
	return nil
}

func main() {
	targetUrl := "http://localhost:999/upload"
	//filename := "/Users/liuhao/Desktop/arraycomslice.png"
	filename := "/Users/liuhao/Documents/图/0BC4C4D581D1895A6BD859FDE53FE72A.jpg"
	result := uploadFile(filename, targetUrl)
	fmt.Println(result)

}
