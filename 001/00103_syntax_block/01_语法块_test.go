package sample

import "fmt"

// 语法块 是由花括号所包含的一系列语句，就像函数体或循环体花括号包裹的内容一样。
// 语法块 内部声明的名字是无法被外部代码访问的，这个块也是内部声明的变量（局部变量）的作用域。
func scope() {
	j := 1
	for i := 0; i < 5; i++ {
		j := 2 // 此处定义 j，会影响到外面的 j吗？
		fmt.Println(j)
	}
	fmt.Println(i) // 会报错吗 ?
	fmt.Println(j) // 值是多少 ?
}
