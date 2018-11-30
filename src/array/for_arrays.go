package array

import (
	"fmt"
)

/**
数组需要注意的地方
1.GO中的数组是值类型，不像Java或者C中是对指针的引用。
2.因为是值类型，所以在传入方法作为参数是会产生拷贝，在方法中修改数组不会修改原来的数组。
3.如果想要在被传入的方法中修改数组的值，需要传递指针进去。
因此GO中数组也被称为数组常量
*/

// 数组初始化方式
func ArrayInitial() {
	//var arrAge = [5]int{34, 32, 56, 18, 19}            // 固定长度，后面内容不可超过5个
	//var arrLazy2 = [...]string{"Tom", "Jack", "Mary"}  //后面内容可以为任意数目
	//var arrKeyValue1 = [5]string{3: "Chris", 4: "Ron"} //指定对应索引位置的值
	//var arrKeyValue2 = [...]string{3: "Chris", 4: "Ron"}
	
}

func ArrayFunc() {
	var arr [5]int

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	// 没有改变
	arrayTest1(arr)
	fmt.Println(arr)

	// 改变了
	arrayTest2(&arr)
	fmt.Println(arr)
}

func arrayTest1(arr [5]int) {
	arr[0] = 9999
}

func arrayTest2(arr *[5]int) {
	(*arr)[0] = 9999
}
