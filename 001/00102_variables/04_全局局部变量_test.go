package sample

import (
	"fmt"
	"testing"
)

// global 是一个全局变量，整个文件内部都可以引用修改其对应的值
var global int

// test 可以打印 a，也可以打印 global
func test() {
	// private 是一个局部变量，只能在 test 函数内部使用
	private := 10
	fmt.Println(private)
	fmt.Println(global)
}

// main 不能直接打印 a， 但可以打印 global
func TestScope(t *testing.T) {
	test()
	fmt.Println(private)
	fmt.Println(global)
}
