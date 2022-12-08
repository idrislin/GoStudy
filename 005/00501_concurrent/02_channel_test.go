package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

/*
	Channel
	channel 是 goroutine 之间的通信机制。
	和 map 类似，channel 也对应一个 make 创建的底层数据结构的引用。
	当我们复制一个 channel 或用于函数参数传递时，我们只是拷贝了一个 channel 引用，
	因此调用者和被调用者将引用同一个 channel 对象。

	channel 还支持 close 操作，用于关闭 channel，随后对基于该 channel 的任何发送操作都将导致panic异常。
	对一个已经被 close 过的 channel 进行接收操作依然可以接受到之前已经成功发送的数据；
	如果 channel 中已经没有数据的话将产生一个零值的数据。

*/

func TestChannel(t *testing.T) {
	// 声明
	ch := make(chan int, 1)
	// 往 ch 添加一个值
	ch <- 5
	// 取出 ch 中的值，如果没有值，则阻塞
	x := <-ch
	fmt.Println(x)
	/*
		此处会死锁，在main goroutine线中，期待从其他goroutine线放入数据，
		但是其他 goroutine 都已经执行完了(all goroutines are asleep)，那么就永远不会有数据放入管道。
		所以，main goroutine线在等一个永远不会来的数据，那整个程序就永远等下去了。
		此时就会报错 死锁
	*/
	<-ch

	ch <- 4
	// 关闭 ch 后尝试读取跟写入
	close(ch)
	x = <-ch
	fmt.Println(x)
	ch <- 3
}

/*
	 无缓存 channel

	一个基于无缓存 channel 的发送操作将导致发送者 goroutine 阻塞，
	直到另一个 goroutine 在相同的 channel 上执行接收操作，
	当发送的值通过 channel 成功传输之后，两个 goroutine 可以继续执行后面的语句

	反之，如果接收操作先发生，那么接收者 goroutine 也将阻塞，
	直到有另一个 goroutine 在相同的 channels 上执行发送操作
*/

func TestBufferLessChannel(t *testing.T) {
	ch := make(chan bool)

	start := time.Now()

	go func() {
		time.Sleep(10 * time.Second)
		ch <- true
	}()

	// 主协程会在此阻塞 10秒，直到另一个 goroutine 往 ch 里面发送信号
	x := <-ch
	fmt.Println(x, ", time since", time.Since(start).Seconds(), "seconds ago")
}

/*
	有缓存 channel

	带缓存的 channel 内部持有一个元素队列。

	向有缓存 channel 的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素

	如果内部缓存队列是满的，那么发送操作将阻塞直到因另一个 goroutine 执行接收操作而释放了新的队列空间,
	相反，如果 channel 是空的，接收操作将阻塞直到有另一个 goroutine 执行发送操作而向队列插入元素
*/

func TestBufferChanel(t *testing.T) {
	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for {
		// 当 x== 10 时，其他 goroutine 已经结束，继续等待信号会引起死锁，那么要怎么才能不死锁呢？
		x := <-ch
		fmt.Println(x)
		// if x == 10 {
		//	break
		// }
	}
}

/*
	### 课堂练习

	- 简要描述 `goroutine` 和系统线程的区别

	- 简要描述 golang 实现高并发的调度模型

	- 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字，
	另一个 goroutine 打印字母，最终结果为`12AB34CD45EF78GH910IJ`

	- 在 main goroutine 中开启一个 goroutine，并实现在满足一定条件时停止额外的那个 goroutine
*/

/*
	### SELECT
	类似switch, 但是只是用来处理通讯(communication)操作。
	它的case可以是send语句，也可以是receive语句，亦或者default
	select 语句和 switch语句一样，它不是循环，它只会选择一个case来处理，
	如果想一直处理channel，可以在外面加一个无限的for循环。
*/

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func TestFibonacci(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

/*
	#### TIMEOUT
	select有很重要的一个应用就是超时处理。
	如果没有case需要处理，select语句就会一直阻塞着。
	这时候我们可能就需要一个超时操作，用来处理超时的情况
*/

func TestTimeout(t *testing.T) {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
}

/*
	### 系统信号
	包 `os.signal` 提供了操作接受系统信号的方法 Notify & Stop
*/

func TestSignal(t *testing.T) {
	go func() {
		stop := make(chan os.Signal)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

		fmt.Println("sub goroutine waiting for syscall signals")
		select {
		case s := <-stop:
			fmt.Printf("receiving signal[%v] from os", s)
			signal.Stop(stop)
		}
	}()

	fmt.Println("blocking main goroutine")
	select {}
}
