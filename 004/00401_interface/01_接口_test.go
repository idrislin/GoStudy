package main

import (
	"fmt"
	"log"
	"testing"
)

/*
	接口
	接口类型是对其它类型行为的抽象和概括；
	因为接口类型不会和特定的实现细节绑定在一起，
	通过这种抽象的方式我们可以让我们的函数更加灵活和更具有适应能力,
	aka [Duck Typing](https://medium.com/@matryer/golang-advent-calendar-day-one-duck-typing-a513aaed544d)
*/

/*
	思考：我们现在有一个公司，里面有不同的工种，但他们都有一个共同特点：工作，
	那么我们就可以定义一个包含工作的接口，凡是工作了的，都算是此公司员工。
*/

// Employee 接口，内含一个函数声明 Work(),
// 所有实现了这个 Work() 方法的类型，都被认为实现了接口 Employee
type Employee interface {
	Work()
}

// Company 是一个包含 接口 Employee的 结构体，
// 所有实现了 Employee的结构体都可以成为 Company的成员
type Company struct {
	Employees []Employee
}

// Coder 实现了 Work() 方法，所以也是 Employee
type Coder struct {
	hairCount int
}

func (c *Coder) Work() {
	c.hairCount -= 1000
}

// UIDesigner 实现了 Work() 方法，所以也是 Employee
type UIDesigner struct {
	hairCount int
}

func (u *UIDesigner) Work() {
	u.hairCount += 1000
}

func TestCompany(t *testing.T) {
	oldBrotherMa := Coder{
		hairCount: 20000,
	}
	oldBrotherCheng := UIDesigner{
		hairCount: 20000,
	}
	mvalley := Company{
		Employees: []Employee{
			&oldBrotherMa,
			&oldBrotherCheng,
		},
	}
	for _, employee := range mvalley.Employees {
		employee.Work()
	}
	log.Println(oldBrotherMa.hairCount)
	log.Println(oldBrotherCheng.hairCount)
}

/*
 	类型断言
	类型断言是一个使用在接口值上的操作。
	语法上它看起来像 `x.(T)` 被称为断言类型，这里 x 表示一个接口的类型和 T 表示一个类型。
	一个类型断言检查它操作对象的动态类型是否和断言的类型匹配。
*/

func TestAssert(t *testing.T) {
	var x interface{}
	x = func() interface{} {
		return 1
	}()

	if intValue, ok := x.(int); ok {
		intValue++
	}
	fmt.Println(x)
}

/*
	课后作业: 编写一个类型实现 `sort.Interface` 接口
*/
