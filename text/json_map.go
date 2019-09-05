package main

import (
	"fmt"
	"encoding/json"
)

func main() {

	b := []byte(`{"IP": "127.0.0.1", "name": "sss"}`)

	m := make(map[string]string)

	//将JSON解析到map中
	err := json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}

	fmt.Println("m:", m)

	for k, v := range m {
		fmt.Println(k, ":", v)
	}

	//直接取值,可能会造成变量类型有误而出错,但是如果和对方商量好的话,是没问题的
	fmt.Println(m["IP"])

}
