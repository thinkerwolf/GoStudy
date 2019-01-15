package function

import (

)

// 函数可以返回多个数，但是必须指定返回值名称
func Add(a int, b int) (add int, minus int) {
	return a + b, a - b
}

func Mutiply(a int, b *int) int {
	*b = 100
	return a * (*b)
}