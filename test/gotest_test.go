package gotest

import (
	"testing"
)

func Test_Div_1(t *testing.T) {
	if i, e := Div(6, 2); i != 3 || e != nil { //try a unit test on function
		t.Error("除法函数测试没通过") // 如果结果不是想要的,就直接报错就行
	} else {
		t.Log("第一个测试通过了") //这里可以记录一些我们期望的信息
	}
}

func Test_Div_2(t *testing.T) {
	//t.Error("就是不让通过")

	//将除数设置为0肯定会报错
	if _, e := Div(6, 0); e == nil {
		t.Error("2测试失败") // 如果不是如预期的那么就报错
	} else {
		t.Log("2测试通过", e) //记录一些你期望记录的信息
	}
}
