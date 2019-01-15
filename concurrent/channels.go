package concurrent

// 类似于Unix中的双向通道Channel，可以用来在线程中分享数据
import (
	"fmt"
)

func add(arr []int, channel chan int) {
	var sum int
	for _, a := range arr {
		sum += a
	}
	channel <- sum
}

func Channel() {
	var arr = []int{4, 5, 6, 7, 8, 9}

	c := make(chan int)

	go add(arr[0:len(arr)/2], c)
	go add(arr[len(arr)/2:], c)

	x, y := <-c, <-c
	fmt.Println("Channel ", x, "+", y, "=", (x + y))
}
