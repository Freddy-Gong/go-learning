package knowledege

import (
	"encoding/json"
	"fmt"
)

//自定义类型
type myInt int

//类型别名
type yourInt = int

func Typee() {
	var n myInt
	n = 100
	fmt.Printf("%T", n) //knowledege.myInt

	var s yourInt
	s = 100
	fmt.Printf("%T", s) //int

	var c rune //rune类型就是一个类型别名 int32
	c = '中'
	fmt.Printf("%T", c) //int32
}

//结构体
/* 结构体的结构
type 结构体名 struct {
	属性名 属性类型
	.........
}
*/
type Person struct {
	name   string
	age    int
	hobby  []string
	gender string
}

func Struc() {
	p := Person{
		name:   "freddy",
		age:    12,
		hobby:  []string{"music", "study"},
		gender: "man",
	}
	fmt.Println(p)
	//匿名结构体：多用于临时场景
	var s struct {
		x string
		y int
	}
	s.x = "name"
	s.y = 10
	fmt.Printf("type:%T, value:%v\n", s, s)

	//结构体是值类型 go的函数传参永远都是拷贝 传值就拷贝值 传地址就拷贝地址
	//如果想在函数内改变外面的变量 就需要传指针
	func(x *Person) {
		x.age = 10
	}(&p)
	fmt.Println(p)

	//如何创建一个指针类型结构体
	//使用new 因为new返回的就是指针
	var p2 = new(Person)
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)

	//使用值列表的形式初始化 值的顺序要和结构体定义时字段的顺序一致
	//这种形式和上面的带有属性名的初始化方式不能混合使用
	//同时我们早初始化之前加上取地址符号
	p3 := &Person{
		"小王子",
		12,
		[]string{
			"xxx",
			"xxx",
		},
		"man",
	}
	fmt.Println(p3)

	//结构体内存模式
	//结构体是占用的一块连续的内存
	//Go中起到好处的内存对齐
}

//方法 是一种作用于特定类型的函数
//接受者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
type dog struct {
	name string
}

//给dog这个结构体定义了一个方法，这样dog的结构体的全部实例 就可以调用这个方法了
//和面向对象类似了
func (d dog) wang() {
	fmt.Println("汪汪汪", d.name)
}

//值接受者和指针接收者
func (p Person) growth() {
	p.age++ //实例本身的值不会变
}
func (p *Person) growthTrue() {
	p.age++ //实例的值会改变
}

//因为go函数的特性，以及结构体是值类型，所以如果用值接受者就不会改变实例的值
//使用指针接收者才会改变实例的值
//保持一致性 如果有一个方法是指针接收者 其他的方法也应该是指针接收者

//如何给基础类型添加方法 比如说我想给int类型添加一个方法
//但是系统默认是不允许的 这个自定义类型就起作用了 例子
type BigString string

func (b BigString) print() {
	fmt.Println(b)
}

//以上代码就相当于给string类型添加了一个print方法 揭晓来我们使用一下
func test() {
	s1 := BigString("asdefw") //将string转化为BigString类型
	s1.print()
}

//用点类似于 面向对象的 继承和多态

//匿名字段 字段没有名字 默认会把类型当成名字
//用的很少 并且容易命名冲突
type person struct {
	string
	int
}

func tes() {
	p := person{
		"he",
		100,
	}
	fmt.Println(p.string)
}

//嵌套结构体 大结构体用小结构体 easy
type Company struct {
	name   string
	addr   string
	deploy []Person
	Person //匿名嵌套结构体 这样Company就可以直接调用Person中的属性
	//Company.Person.hobby == Company.hobby
}

var c = Company{
	name: "有限公司",
	addr: "122.111.11.11",
	deploy: []Person{
		{
			name: "xx",
			age:  12,
		},
		{
			name: "ad",
			age:  55,
		},
	},
}

//模拟实现继承
type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s会动", a.name)
}

type cat struct {
	feet   uint8
	animal //使用了匿名嵌套结构体 所以cat就可以直接调用animal的属性和方法了
	//间接实现了继承 intersting
}

func (c cat) miao() {
	fmt.Println("miao~", c.name)
}

//结构体和JSON相互转换
func JSON_STRUCT() {
	p := Person{
		name:   "freddy",
		age:    12,
		hobby:  []string{"music", "study"},
		gender: "man",
	}
	b, err := json.Marshal(p) //序列化
	if err != nil {
		fmt.Println(err)
		return
	}
	//json.Marshal返回的结果是一个int64 所以需要转成string
	fmt.Println(string(b)) // {} 会发现结果是一个空
	//？？因为Person里面的字段是小写的 也就是说是对外不可见
	//而序列化是在json包中进行的 所以json包拿不到字段 所以打印的结果是{}
	//解决把Person的字段变成首字母大写
	//但是这时又出现一个问题，因为前端没有这种习惯 所以基本上不会出现首字母大写的情况
	//所以这时又引入了一个tag的功能 例子
	type Student struct {
		Name string `json:"name" db:"name" ini:"name"`
		Age  int    `json:"age"`
	}
	s1 := Student{
		Name: "xx",
		Age:  12,
	}
	js1, err := json.Marshal(s1)
	fmt.Printf(string(js1))

	//反序列化 以上tag的功能又能够在反序列化的时候将对应JSON的字段
	//转为对应结构体的字段
	s2 := `{"name":"fredd","age":11}`
	var ss Student
	//反序列化第一个参数是字节 第二个参数是一个对应类型的变量
	//用这个变量来承接转化后的值 所以要传地址
	json.Unmarshal([]byte(s2), &ss)
	fmt.Println(ss)
	fmt.Printf("%#v", ss)
}
