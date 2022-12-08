package sample

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	/*
		哈希表是一种巧妙并且实用的数据结构。
		它是一个无序的 key/value 对的集合，
		其中所有的 key 都是不同的，然后通过给定的 key 可以在常数时间复杂度内检索、更新或删除对应的 value。

		map 的 key 必须是支持 == 比较运算符的数据类型。问：有哪些基础类型不能做 map 的 key ？
	*/

	// 初始化一个 map的两种写法
	a := make(map[int]string)
	a[1] = "one"
	a[2] = "two"
	a[3] = "three"

	b := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	delete(b, "one")
	// 以上这些操作是安全的，即使这些 key 不在 map 中也没有关系；如果查找失败将返回 value 类型对应的零值

	// 遍历 map

	for key, value := range a {
		fmt.Printf("[for loop 1]addr:%p,key:%d,value:%s\n", &key, key, value)
	}

	for key := range a {
		fmt.Printf("[for loop 2]addr:%p,key:%d,value:%s\n", &key, key, a[key])
	}

	for _, value := range a {
		fmt.Printf("[for loop 3] value:%s\n", value)
	}

	// map 是指针传递的，使用之前必须初始化
	// 错误示例：
	var nilMap map[int]int
	nilMap[1] = 1
}
