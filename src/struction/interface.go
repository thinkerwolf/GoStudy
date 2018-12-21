package struction

import (
	"fmt"
)

type Man interface {
	SayHi()
	Sing(s string)
}

// Human 实现 Man接口
func (h Human) SayHi() {
	fmt.Println(h.name, " is saying hello")
}

// Human 实现 Man接口
func (h Human) Sing(s string) {
	fmt.Println(h.name, " is saying a song named ", s)
}

func TestInterface() {
	var man Man
	man = Human{name: "Bruce", age: 46, weight: 180}

	man.SayHi()
	man.Sing("Yesterday")

}
