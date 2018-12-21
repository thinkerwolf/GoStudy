package array

import (
	"fmt"
)

/**
切片是引用，所以相比数组更有效率。类似于Java中的List
绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针!!
*/

func SliceFunc() {
	// 初始化切片的方式
	var slice1 = []int{3, 4, 5, 6} // 不要和数组[5]int{3, 4, 5, 6}、[...]int{3, 4, 5, 6}混淆
	fmt.Printf("Slice1 len is %d, cap is %d\n", len(slice1), cap(slice1))
	sliceTest1(slice1)
	fmt.Println(slice1)

	var arrLazy = [...]string{"Tom", "Jack", "Mary", "LiMing"}
	var slice2 = arrLazy[0:3]
	fmt.Println(slice2)
	fmt.Printf("Slice2 len is %d, cap is %d\n", len(slice2), cap(slice2)) // len 3, cap 4

	createSlice()
}

func sliceTest1(slice []int) {
	slice[0] = 990
}

func createSlice() {
	// 创建一个长度50，容量为100的切片
	// make([]type, len, cap)
	var sliceMake = make([]int, 50, 100)

	// make([cap]type[x:y])
	var sliceNew = new([100]int)[0:50]

	fmt.Println("sliceMake : ", sliceMake)
	fmt.Println("sliceNew : ", sliceNew)

	var p *[]int = new([]int)
	fmt.Println(p)

}
