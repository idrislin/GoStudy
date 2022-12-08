package main

import (
	"fmt"
	"golang/utils"
	"log"
	"os"
	"testing"
	"time"
)

// sum returns the result of x + y
/*
	函数声明包括函数名、形参列表、返回值列表（可省略）以及函数体，当然，最好包含对齐功能解释的注释

	go doc 式注释的正确格式如上，"//" + 空格 + 函数名 + 返回值说明

*/
func sum(x, y int) int {
	return x + y
}

/*
	声明不定参数函数：
	在声明可变参数函数时，需要在参数列表的最后一个参数类型之前加上省略符号 `...`，
	这表示该函数会接收任意数量的该类型参数。

	在函数体中，`params` 被看作是类型为 `[]int` 的切片。
*/

// sumAll returns sum of all params
func sumAll(params ...int) (total int) {
	for _, val := range params {
		total += val
	}
	return
}

/*
	调用函数
	每一次函数调用都必须按照声明顺序为所有参数提供实参（参数值）
*/

// sigma returns recursively call of itself plus x
func sigma(x int) int {
	if x <= 0 {
		return 0
	}
	return sigma(x-1) + x
}

/*
	Panic
	Go 的类型系统会在编译时捕获很多错误，但有些错误只能在运行时检查，如数组访问越界、空指针引用等。这些运行时错误会引起 panic 异常。

	一般而言，当 panic 异常发生时，程序会中断运行，并输出日志信息。
*/
func TestPanic(t *testing.T) {
	var a = []int{1, 2, 3}
	fmt.Println(a[3])
}

/*
	Defer
	当 defer 语句被执行时，跟在 defer 后面的函数会被延迟执行。
	直到包含该 defer 语句的函数执行完毕时，defer 后的函数才会被执行，
	不论包含 defer 语句的函数是通过 return 正常结束，还是由于 panic 导致的异常结束。

	注意：当使用 os.Exit 时 defer 将不会 被执行（通过命令 ctrl+c 结束程序，也会导致 defer 不运行）
*/

func openFile() {
	var f, err = os.Open("filename")
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("running defer")
	}()
	// openFile()
	fmt.Println("running main")

	// 在 sleep 期间结束运行，会发现 defer 并没有被执行
	time.Sleep(10 * time.Second)
}

/*
	Recover
	通常来说，不应该对 panic 异常做任何处理，但有时，我们可以从异常中恢复，在程序崩溃前，做一些操作。
	当 web 服务器发生 panic 时，在崩溃前应该将所有的连接关闭。

	注意：
	不管三七二十一恢复所有的 panic 不是可取的做法；
	因为在 panic 之后，无法保证全局变量的状态仍然和我们预期一致。
	作为被广泛遵守的规范，不应该试图去恢复其他包引起的 panic 。
	公有的 API 应该将函数的运行失败作为 error 返回，而不是 panic。
*/

func TestRecover(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	panic("don't panic")
}

/*
	课堂作业：使用 panic 和 recover 编写一个不包含 return 语句但能返回一个非零值的函数
*/

func nonReturn() (result int) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			result = 1
		}
	}()
	panic("don't panic")
}

func TestNonReturn(t *testing.T) {
	fmt.Println(nonReturn())
}

/*
	扩展知识：Go 支持匿名函数， 并能用其构造 闭包。 匿名函数在你想定义一个不需要命名的内联函数时是很实用的
*/

func TestClosure(t *testing.T) {
	counter := utils.Counter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
