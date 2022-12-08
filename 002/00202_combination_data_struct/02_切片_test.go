package sample

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	/*
		Slice（切片）代表变长的序列，序列中每个元素都有相同的类型;
		一个 slice 类型一般写作 []T，其中 T 代表 slice 中元素的类型；
		slice 的语法和数组很像，只是没有固定长度而已
	*/

	/*
		一个 slice 由三个部分构成：
		指针: 指向第一个 slice 元素对应的底层数组元素的地址
		长度: 对应 slice 中元素的数目，长度不能超过容量
		容量: 一般是从 slice 的开始位置到底层数据的结尾位置

		package runtime
		type slice struct {
			array unsafe.Pointer
			len   int
			cap   int
		}
	*/

	/*
		slice 的切片操作 s[i:j],
		其中 0 ≤ i ≤ j ≤ cap(s)，用于创建一个新的 slice，引用 s 的从第 i 个元素开始到第 j-1 个元素的子序列。
	*/
	var s = make([]int, 10, 10)
	var ss = s[5:8]
	fmt.Printf("[1] %p,%v\n", &s, s)
	fmt.Printf("[2] %p,%v\n", &ss, ss)

	/*

		初始化 slice 可使用内建函数 make([]int, len, cap) ,
		在底层，make 创建了一个匿名的数组变量，然后返回一个 slice ；
		只有通过返回的 slice 才能引用底层匿名的数组变量;
		slice 只引用了底层数组的前 len 个元素，但是容量将包含整个的数组;
		额外的元素是留给未来的增长用的;
	*/
	var slice = make([]string, 10, 20)
	fmt.Printf("[3] len:%d, cap:%d, a:%v\n", len(slice), cap(slice), slice)

	a := []int{1, 2, 3}
	b := a[:]
	fmt.Printf("[4] addr:%p,elems:%v,len:%d,cap:%d\n", &a, a, len(a), cap(a))
	fmt.Printf("[5] addr:%p,elems:%v,len:%d,cap:%d\n", &b, b, len(b), cap(b))

	// 修改第一个元素的值
	b[0] = 0
	fmt.Printf("[6] addr:%p,elems:%v,len:%d,cap:%d\n", &b, b, len(b), cap(b))
	// 添加一个元素
	b = append(b, 4)
	// 添加两个元素
	b = append(b, 5, 6)
	// 添加另一个切片
	b = append(b, a...)
	fmt.Printf("[7] addr:%p,elems:%v,len:%d,cap:%d\n", &b, b, len(b), cap(b))

	// 遍历切片
	for i, num := range a {
		fmt.Printf("[for loop 1] addr:%p,i:%d,num:%d\n", &i, i, num)
	}

	for i := range a {
		fmt.Printf("[for loop 2] addr:%p,i:%d,num:%d\n", &i, i, a[i])
	}

	for _, num := range a {
		fmt.Printf("[for loop 3] addr:%p,num:%d\n", &num, num)
	}

	// 错误的修改 a的元素的方法
	for _, num := range a {
		num++
		fmt.Printf("[for loop 4] addr:%p,num:%d\n", &num, num)
	}
	fmt.Println("[8]", a)

	// 正确的修改 a元素的方法
	for i := range a {
		a[i]++
	}
	fmt.Println("[9]", a)
}
