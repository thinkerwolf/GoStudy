package base

import (
	"fmt"
)

func SwitchFunc() {
	var num1 = 100
	switch num1 {
	case 10:

	case 100:
		fmt.Println(100)
		fallthrough   // 与java中不加break效果一致
	case 101:
		fmt.Println(101)
	default:
		fmt.Println(9999)
	}

}
