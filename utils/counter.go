package utils

// Counter 计数器
func Counter() func() int {
	var i int
	return func() int {
		i += 1
		return i
	}
}
