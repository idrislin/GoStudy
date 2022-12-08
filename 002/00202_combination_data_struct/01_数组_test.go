package sample

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	// 数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成
	var a [3]int
	fmt.Println("[1]", a[0], ",", a[len(a)-1])

	for i, v := range a {
		fmt.Printf("[for loop 1] %d %d\n", i, v)
	}

	// 默认情况下，数组的每个元素都被初始化为元素类型对应的零值，对于数字类型来说就是 0。
	// 我们也可以在初始化数组时给数组赋值
	var r [3]int = [3]int{1, 2}
	fmt.Printf("[2] %T, %v\n", r, r)

	q := [...]int{1, 2, 3}
	fmt.Printf("[3] %T, %v\n", q, q)

	p := [...]int{25: 56, 45: 98}
	fmt.Printf("[4] %T, %v\n", p, p)
}

// TestAlterArray 测试修改数组元素
func TestAlterArray(t *testing.T) {
	var a [3]int
	for _, v := range a {
		v++
		// 下面打印什么？
		fmt.Printf("%p,%d\n", &v, v)
	}
	// 上面的代码怎么修改才能打印出 [1,2,3] ？
	fmt.Println(a)

	var array = [3]int{1, 2, 3}
	array = append(array, 4) // 为什么会报错？
}

// TestTypeOf 不同长度类型判断
func TestTypeOf(t *testing.T) {
	var a [4]int
	var b [3]int
	// a,b 是同一种数据类型吗?
	fmt.Println(a == b)

	// 数组长度能不能用未知/可变变量来定义？
	var slice = make([]int, 3, 3)
	var c [len(slice)]int
	fmt.Println(c)
}
