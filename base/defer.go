package base

// defer关键字，类似于java中的finally

import (
	"fmt"
)

func Defer() {
	i := 0
	defer fmt.Printf("The Defer value is %d\n", i)
	i++
}