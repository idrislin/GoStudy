package sample

import (
	"fmt"
	"testing"
)

// 结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员
type Person struct {
	ID   string
	Name string
	Age  int
}

// Employee 结构体内可嵌套结构体
type Employee struct {
	Person
	Salary int
}

type Employer struct {
	Person
	Employees []Employee
}

type EmptyStruct struct{}

func TestStruct(t *testing.T) {
	// 声明一个结构体方式一
	var oldBrotherMa Person
	oldBrotherMa.ID = "1"
	oldBrotherMa.Name = "Kcat"

	// 通过指针修改内部成员的值
	agePtr := &oldBrotherMa.Age
	*agePtr++
	fmt.Println(oldBrotherMa)

	// 声明一个结构体方式二
	var yangBrotherMa = Person{
		ID:   "2",
		Name: "K不发音",
	}
	// 直接修改结构体内部成员的值
	yangBrotherMa.Age += 1

	// 结构体内可嵌套结构体
	employeeMa := Employee{
		Person: Person{
			ID:   "3",
			Name: "kcat",
		},
		Salary: 100,
	}
	employeeMa.Person.Age += 1
}
