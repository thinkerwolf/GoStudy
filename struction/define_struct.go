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
	Name   string
	Age    int
	Weight int
}

type Student struct {
	Human      // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}

type GoodStudent struct {
	Student
	good bool
}

func NewHuman(name string, age int, weight int) *Human {
	return &Human{Name: name, Age: age, Weight: weight}
}

// Human 实现 Man接口
func (h Human) SayHi() {
	fmt.Println(h.Name, " is saying hello")
}

// Human 实现 Man接口
func (h Human) Sing(s string) {
	fmt.Println(h.Name, " is saying a song named ", s)
}

func TestInterface() {
	var man Man
	man = Human{Name: "Bruce", Age: 46, Weight: 180}

	man.SayHi()
	man.Sing("Yesterday")

}

func AnnomyStruct() {
	// 我们初始化一个学生
	mark := GoodStudent{Student{Human{"Mark", 25, 120}, "Computer Science"}, true}

	// 我们访问相应的字段
	fmt.Println("His name is ", mark.Name)
	fmt.Println("His age is ", mark.Age)
	fmt.Println("His weight is ", mark.Weight)
	fmt.Println("His speciality is ", mark.speciality)
	// 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	// 修改他的年龄信息
	fmt.Println("Mark become old")
	mark.Age = 46
	fmt.Println("His age is", mark.Age)
	// 修改他的体重信息
	fmt.Println("Mark is not an athlet anymore")
	mark.Weight += 60
	fmt.Println("His weight is", mark.Weight)
}
