package struction

import (
	"fmt"
)

/* 定义一个新类 , struct类似java的class */
type people struct {
	name string
	age  int
}

func Older(p1, p2 people) (bool, int) {
	if p1.age > p2.age {
		return true, p1.age - p2.age
	}
	return false, p1.age - p2.age
}

func TestStruct() {
	var p1 people
	p1.age = 20
	p1.name = "Jack"

	var p2 people
	p2 = people{name: "Michael", age: 21}

	p3 := people{"William", 32}

	fmt.Println(Older(p1, p2))

	fmt.Println(Older(p2, p3))
}

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}

type GoodStudent struct {
	Student
	good bool
}

func AnnomyStruct() {
	// 我们初始化一个学生	
	mark := GoodStudent{Student{Human{"Mark", 25, 120}, "Computer Science"}, true}
	
	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
	// 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	// 修改他的年龄信息
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age)
	// 修改他的体重信息
	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)
	
	
}
