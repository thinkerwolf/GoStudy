package concurrent

// 带有缓冲的channel，可以指定往channel中存入的个数。如果存储超过指定个数，会阻塞，直到从channel中取出
import (
	"fmt"
)


func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}

func BufferedChannel() {
	n := 10
	c := make(chan int, n)
	go fibonacci(n + 2, c)
	for i := range c {
		fmt.Println("BufferedChannel : ", i)
	}
	
}