package sample

import (
	"log"
	"testing"
)

func TestBoolean(t *testing.T) {
	/*
	 package builtin

	  // bool is the set of boolean values, true and false.
	  type bool bool

	  // true and false are the two untyped boolean values.
	  const (
	      true  = 0 == 0 // Untyped bool.
	      false = 0 != 0 // Untyped bool.
	  )
	*/

	// 布尔值可以和 `&&`（AND）和 `||`（OR）操作符结合，并且有短路行为：
	// 如果运算符左边值已经可以确定整个布尔表达式的值，那么运算符右边的值将不再被求值，
	// 因此下面的表达式是安全的：
	var a *int
	if a != nil && *a == 2333 {
		log.Println(*a)
	}
	log.Println("safely exiting...")
}
