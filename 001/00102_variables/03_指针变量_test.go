package sample

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	// 声明一个 int类型的空指针
	var p *int
	// 初始化指针，可以使用内置函数 new()
	p = new(int)

	// 赋值: 通过引用其他变量的地址赋值
	x := 0
	p = &x
	fmt.Println(*p)
	// 赋值: 对 p 指向的变量直接赋值
	*p = 2
	fmt.Println(*p)

	// 在取指针指向的值时，一定要先判断指针是否为空
	var a *int
	if a != nil {
		fmt.Println(*a) // 假设不判空， 这里会 panic
	}
}
