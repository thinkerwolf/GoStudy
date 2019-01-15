package base

import (
	"fmt"
)

func StringPointer() {
	s := "Good Bye"
	var p *string
	p = &s
	*p = "vicle"
	fmt.Printf("Here is the pointer p: %p\n", p) // prints address
    fmt.Printf("Here is the string *p: %s\n", *p) // prints string
    fmt.Printf("Here is the string s: %s\n", s) // prints same string
}
