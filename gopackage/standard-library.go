/**
 GO 常用标准库解析
archive/tar 和 /zip-compress：压缩(解压缩)文件功能。

fmt-io-bufio-path/filepath-flag:
fmt: 提供了格式化输入输出功能。
io: 提供了基本输入输出功能，大多数是围绕系统功能的封装。
bufio: 缓冲输入输出功能的封装。
path/filepath: 用来操作在当前系统中的目标文件名路径。
flag: 对命令行参数的操作。

strings-strconv-unicode-regexp-bytes:
strings: 提供对字符串的操作。
strconv: 提供将字符串转换为基础类型的功能。
unicode: 为 unicode 型的字符串提供特殊的功能。
regexp: 正则表达式功能。
bytes: 提供对字符型分片的操作。
index/suffixarray: 子字符串快速查询。

math-math/cmath-math/big-math/rand-sort:
math: 基本的数学函数。
math/cmath: 对复数的操作。
math/rand: 伪随机数生成。
sort: 为数组排序和自定义集合。
math/big: 大数的实现和计算。
container-/list-ring-heap: 实现对集合的操作。
list: 双链表。
ring: 环形链表。

*/

package gopackage

import (
	"container/list"
	"fmt"
	"unsafe"

	"github.com/GoStudy/struction"
)

func AList() {
	el := list.New()
	el.PushFront("23123")
	el.PushBack(12)

	// 遍历链表
	for e := el.Front(); e != nil; e = e.Next() {
		fmt.Println("Link element value:", e.Value)
	}
}

func UnSafe() {
	human := struction.Human{Name: "Jack", Age: 12, Weight: 50}
	fmt.Println("Human.Age offset:", unsafe.Offsetof(human.Age))
	fmt.Println("Human.Weight offset:", unsafe.Offsetof(human.Weight))
	fmt.Println("Human.Weight Size:", unsafe.Sizeof(human.Weight))
}
