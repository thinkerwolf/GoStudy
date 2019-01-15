package concurrent

import (
	"fmt"
	"runtime"
)

// Go语言的多线程术语是 goroutine 相当于java的Thread，但是与Thread有本质区别

func say(str string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println("Hello ", i, " : ", str)
	}
}

func GoRoutine() {
	go say("world")   // 另起一个goroutine执行
	say("Go")    // 在main goroutine执行
}
