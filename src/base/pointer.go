package base

import (
	"fmt"
)

/** go指针 */
func Pointer() {
	
	var i1 = 5
	//fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
	
	var i1P *int // 定义一个指针参数,与c的指针类似
	i1P = &i1
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, i1P)
	fmt.Printf("The value at memory location %p is %d\n", i1P, *i1P)
	
	
	
}