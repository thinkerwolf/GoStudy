package function

import (
	"fmt"
	"strings"
)

// 匿名函数就是闭包

func FuncAnony() {
	func(sum int) {
		for i := 1; i <= 1e6; i++ {
			sum += i
		}
		fmt.Printf("The FuncAnony result is %d\n", sum)
	}(9)

	makeJpg := MakeAddSuffix(".jpg")
	makeTxt := MakeAddSuffix(".txt")

	fmt.Println("MakeJpg " + makeJpg("photo"))
	fmt.Println("MakeTxt " + makeTxt("hello"))
}

// 闭包作为返回，返回一个函数
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
