package sample

import (
	"log"
	"os"
	"testing"
)

func TestDeclare(t *testing.T) {
	// 一般语法
	var className string = "golang"
	log.Printf((className))
	// 简略写法, go 根据值猜测变量的基础类型，下面的 x 类型为 string
	name := "gogo"
	log.Printf((name))

	// 声明一组变量
	var i, j, k int
	i, j, k = 0, 0, 0
	log.Printf("i = %s, j = %s, k = %s", i, j, k)

	var b, t, s = true, 2.3, "dddd"
	b, t, s = true, 2.3, "dddd"
	log.Printf("b = %s, f = %s, s = %s", b, t, s)

	// 通过函数返回值初始化
	var f, err = os.Open(s)
	f, err := os.Open(s) // 简略写法必须声明至少一个新变量，否则会编译失败

	// 简略写法必须声明至少一个新变量，否则会编译失败
	goFile, err := os.Open("main.go")
	csvFile, err := os.Open("sample.csv")

}
