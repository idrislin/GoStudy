package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
	Goroutine
	在Go语言中，每一个并发的执行单元叫作一个 goroutine。

	当一个程序启动时，其主函数即在一个单独的 goroutine 中运行，我们叫它 main goroutine。
	新的 goroutine 会用 `go` 语句来创建。

	主函数返回时，所有的 goroutine 都会被直接打断，程序退出。
	除了从主函数退出或者直接终止程序之外，
	没有其它的编程方法能够让一个 goroutine 来打断另一个的执行。
	但是有一种方式来实现这个目的，通过 goroutine 之间的通信来让被请求的 goroutine 自行结束执行
*/

func TestGoroutine(t *testing.T) {
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep over 1")
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep over 2")
	}()
}

/*
	通过上面代码可以看出，两个 goroutine 都没能打印任何值，
	这是因为 主函数退出，所有的 goroutine 都被打断，程序退出了。
	那么如何才能让主函数等待其他 goroutine 执行成功呢？
	tips for loop, channel, waitGroup
*/

func TestGoroutineWithWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("sleep over 1")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("sleep over 2")
	}()

	wg.Wait()
}

/*
	课堂扩展：cp 输出结果是什么? 怎样才能按顺序输出 ABC？
	tips: range 浅拷贝 & slice 线程安全相关考虑 sync.Mutex
*/

func TestDataRace(t *testing.T) {
	list := []string{"A", "B", "C"}

	cp := make([]*string, 0)

	var wg sync.WaitGroup
	for _, v := range list {
		wg.Add(1)

		go func() {
			defer func() {
				wg.Done()
			}()

			cp = append(cp, &v)
		}()
	}

	wg.Wait()
	for i := range cp {
		fmt.Printf("%v,", *cp[i])
	}
}
