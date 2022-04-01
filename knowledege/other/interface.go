package knowledege

import (
	"fmt"
)

func init() {
	fmt.Println("knowledege包中的init函数 在interface.go中")
}

//接口是一种类型
//不管是参数 变量 返回值 我们都可以把接口当成一种类型来使用
//是对结构体方法的约束
//只要具有某种方法的结构体就都满足接口的条件

type Cat struct {
}
type Dog struct {
}
type PPerson struct {
}

func (c Cat) speak() {
	fmt.Println("喵喵喵")
}
func (d Dog) speak() {
	fmt.Println("汪汪汪")
}
func (p PPerson) speak() {
	fmt.Println("啊啊啊")
}

//我们给三种类型中的每一种类型都定义了一个speak方法
//我很希望调用hit的时候，变量x可一调用speak方法 但无所谓x的struct类型是什么
//这时我们就需要接口来对参数进行方法的约束
type Speak interface {
	speak()
}

//我们把x的类型定义为speak接口
func Hit(x Speak) {
	c := Cat{}
	d := Dog{}
	p := PPerson{}
	fmt.Println(c, d, p)
	x.speak()
}

//接口的定义
/*
type 接口名 interface {
	方法名1(参数1, 参数2...)(返回值1, 返回值2...)
	方法名2(参数1, 参数2...)(返回值1, 返回值2...)
	...
}
如果一个变量实现了接口中规定的所有方法，那么这个变量就实现了这个接口
可以称为这个接口类型的变量.
*/

//接口类型分为两部分 一个部分为动态类型 一部分为动态值 这样就实现了 接口变量可以储存不同的值

func Dy() {
	var s Speak
	ca := Cat{}
	s = ca
	fmt.Printf("%T\n", s) //s的类型变为 Cat了
	do := Dog{}
	s = do
	fmt.Printf("%T\n", s) //s的类型变为 Dog了
}

//方法的值接受者和指针接收者 对interface的影响 使用区别
type talk interface {
	talk()
}
type walk interface {
	walk()
}
type me struct {
	Name string
}

//使用值接收者
func (m me) talk() {
	fmt.Println("talk")
}

//使用指针接收者
func (m *me) walk() {
	fmt.Println("walk")
}

func point() {
	//使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存
	var t1 talk
	m1 := me{Name: "nax"}
	t1 = m1
	fmt.Printf("%T\n", t1)
	t1 = &m1
	fmt.Printf("%T\n", t1)
	//使用指针接收者实现接口只能存节后题指针类型的变量
	var w1 walk
	m2 := me{Name: "axs"}
	w1 = &m2
	fmt.Printf("%T\n", w1)
}

//接口和结构体的关系
//同一个结构体可以实现多个接口
//接口还可以嵌套
type mover interface {
	move()
}
type eater interface {
	eat()
}
type animals interface {
	mover
	eater
}
type ca struct {
}

//cat实现了mover接口
func (c *ca) move() {
}

//cat实现了eater接口
func (c *ca) eat() {
}

//空接口 interface{} 和TS中的 any 一样的
var m1 map[string]interface{}

//m1的值就可以是任意类型 奇怪go中没有联合类型 。。
//类型断言
func xx(a interface{}) {
	fmt.Println(a)
	//我想知道空接口接受的参数的类型

	//断言a的类型是string
	str, ok := a.(string)
	if !ok {
		fmt.Println("参数不是string类型")
	}
	fmt.Println(str)
	//断言a的类型是int类型
	number, ok := a.(int)
	if !ok {
		fmt.Println("参数不是int类型")
	}
	fmt.Println(number)

	//以上可以使用case形式重写
	switch v := a.(type) {
	case string:
		fmt.Println("是一个字符串,", v)
	case int32:
		fmt.Println("是一个数字,", v)
	case bool:
		fmt.Println("是一个布尔,", v)

	}
}
