package concurrent

import (
	"fmt"
)

/* select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行。
当多个channel都准备好的时候，select是随机的选择一个执行的。
*/
func SelectChannel() {
	c := make(chan int, 10)
	quit := make(chan int)

	x, y := 1, 1

	selectFun := func() {
		for {
			select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				return
			}
		}
	}

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Select Channel : ", <-c)
		}
		quit <- 1
	}()
	selectFun()
}
