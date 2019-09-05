package main

import (
	"fmt"
	"regexp"
)

func main() {
	src := []byte(`
		call hello alice
		hello bob
		call hello eve
	`)
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {

		fmt.Println(string(s[0]))
		// Expand 要配合 FindSubmatchIndex 一起使用。FindSubmatchIndex 在 src 中进行
		// 查找，将结果存入 match 中。这样就可以通过 src 和 match 得到匹配的字符串。
		// template 是替换内容，可以使用分组引用符 $1、$2、$name 等。Expane 将其中的分
		// 组引用符替换为前面匹配到的字符串。然后追加到 dst 的尾部（dst 可以为空）。
		// 说白了 Expand 就是一次替换过程，只不过需要 FindSubmatchIndex 的配合。
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
}
