package sample

import (
	"fmt"
	"testing"
)

type Weekday int

// 常量表达式的值在编译期计算，而不是在运行期。每种常量的潜在类型都是基础类型：boolean、string 或数字。
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	pi            = 3.14159265358979323846264338327950288419716939937510582097494459
	approximatePi = 3.14159
)

func TestConstants(t *testing.T) {
	fmt.Println(pi)
}
