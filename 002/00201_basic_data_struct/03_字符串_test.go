package sample

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	/*
	  字符串是不可改变的字节序列
	  package builtin

	  // string is the set of all strings of 8-bit bytes, conventionally but not
	  // necessarily representing UTF-8-encoded text. A string may be empty, but
	  // not nil. Values of string type are immutable.
	  type string string
	*/
	str1 := "123\n45"
	str2 := `123\n45`

	// 标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv 和 unicode 包。
	bytes.Compare([]byte(str1), []byte(str2))
	strings.Compare(str1, str2)
	num, _ := strconv.Atoi("123")
	fmt.Printf("%d", num)

	// 问： 字符串是不可改变的字节序列, 那 strings.ReplaceAll 是怎么改变字符串内容的？
	newStr := strings.ReplaceAll(str1, "r", "")
	fmt.Printf("old: [%p,%s], new:[%p,%s]\n", &str1, str1, &newStr, newStr)
}
