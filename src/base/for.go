package base
import (
	"fmt"
)

func ForFunc() {
	// for range结构
	str := "Go is a beautiful lanauage"
	for p, val := range str {
		fmt.Printf("The index is %d, The value is %c\n", p, val)
	}
	
	
}

