package sample

import (
	"builtin"
	"fmt"
	"testing"
)

type myInt builtin.int

func TestNumber(t *testing.T) {
	// 1. 整型分为有符号和无符号两种类型，具体定义可看源码 [builtin/builtin.go]
	var _ int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint
	/*
		像 Java 里面的 `int` 类型的底层定义是 C++里面的 `int`，这是因为 Java的编译器是用 C++写的；
		我们平时要自定义类型时候，像 [myInt] ，就会是使用其他基础类型来声明的，
		那么为什么 GO 里面的 ` int` 的底层定义是 GO 的类型 ` int` （即自身来定义自身）?

		[语言的自举](https://golang.design/under-the-hood/zh-cn/part3tools/ch11compile/bootstrap/)
	*/

	// 2. 浮点数有两种精度，float32 和 float64，具体定义可看源码 [builtin/builtin.go]
	var _ float32, float64

	// 3. 数字之间可以互相转换
	var a int32 = 2333
	var b int64
	b = int64(a)
	fmt.Println("[1]",a, ",", b)

	var x float64 = 32767.1214
	var y int
	fmt.Println("[2]",x, ",", y)
	y = int(x)
	fmt.Println("[3]",x, ",", y)
	x = float64(y)
	fmt.Println("[4]",x, ",", y)
	// 问：x 的值发生了什么变化？为什么？
}
