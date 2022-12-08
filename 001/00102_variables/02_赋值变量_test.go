package sample

import (
	"fmt"
	"os"
	"testing"
)

func TestAssignment(t *testing.T) {
	x := 2
	fmt.Println(x)
	x *= 2
	fmt.Println(x)
	x++
	fmt.Println(x)

	var a, b = 1, 2
	fmt.Println(a, b)
	// 可直接交换两个变量的值
	a, b = b, a
	fmt.Println(a, b)
	{
		// C++ 里面交换值做法(GO 写法)，需要定义中间变量
		temp := a
		a = b
		b = temp
	}

	// 无需处理、不会被引用的的变量，可以用下划线替代
	_, err := os.Open("main.go")
	if err != nil {
		fmt.Println(err)
	}
}
