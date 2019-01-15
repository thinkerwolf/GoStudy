package function

import (
	"fmt"
)

// 函数作为参数
func FuncAsParm(y int, f func(int, int) int) {
	r := f(100, 999)
	r += y
	fmt.Printf("The FuncAsParm result is %d\n", r)
}


func AddParm(a, b int) int {
	return a + b
}