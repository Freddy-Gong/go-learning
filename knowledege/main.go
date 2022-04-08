package main

import (
	"fmt"
	knowledege "learn/knowledege/other"
	"strings"
)

//在Go语言程序执行时导入包语句会自动出发包内部init函数的调用。需要注意的是
//init函数没有参数也没有返回值 init函数在程序运行时自动被调用执行 不能在代码中主动调用他
func init() {
	fmt.Println("我是自己执行的")
}

//包中执行时机 全局声明->init()->main() 并且会首先执行引入进的包的main函数

//全局声明变量 可以不使用
//但是局部变量必须使用
//Go语言中推荐使用驼峰式命名规则
var name string
var age int
var isOk bool

//批量声明
var (
	name1 string //""
	age1  int    //0
	isOk1 bool   //false
)

func main() {
	//赋值
	name = "zhangsan"
	age = 18
	isOk = true

	//批量赋值
	name1, age1, isOk1 = "lisi", 19, true

	//输出
	fmt.Printf("name:%s\n", name) //%s是一个占位符，表示字符串 \n是换行符
	fmt.Println(name, age, isOk)
	fmt.Println(name1, age1, isOk1)

	//声明变量同时负值
	var s1 string = "hello"
	fmt.Println(s1)
	//类型推导
	var s2 = "world"
	fmt.Println(s2)
	//简短变量声明 只能在函数内部使用
	s3 := "go"
	fmt.Println(s3)
	//同一个作用域中不能重复声明同名的变量
	//匿名变量 _
	knowledege.Mtest()
}

//常量 定义之后就不能进行修改了 const
const pi = 3.1415926

//批量声明常量
const (
	STATUSOK = 200
	NOTFOUND = 404
)

//批量声明常量时，如果某一行声明后没有复制，那么默认值就是上一行的值
const (
	n1 = 100
	n2
	n3
)

//iota 常量计数器
//每新增一行const声明，iota就会自动加1
const (
	a1 = iota //0
	a2        //1
	a3        //2
	a4        //3
)

//其它例子
const (
	b1 = iota //0
	b2        //1
	_         //2
	b3        //3
)
const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)
const (
	d1, d2 = iota + 1, iota + 2 //1,2
	d3, d4 = iota + 1, iota + 2 //3,4
)
const (
	_  = iota
	KB = 1 << (10 * iota) //1 << (10 * 1) 1024
	MB = 1 << (10 * iota) //1 << (10 * 2) 1024 * 1024
	GB = 1 << (10 * iota) //1 << (10 * 3) ......
	TB = 1 << (10 * iota) //1 << (10 * 4)
	PB = 1 << (10 * iota) //1 << (10 * 5)
)

//##基本数据类型
//整数 int8 int16 int32 int64 uint8 uint16 uint32 uint64 int uint uintptr
//int表示有符号整数，uint表示无符号整数（正负号），uintptr表示无符号指针
//八进制 十六进制
func num() {
	//十进制
	num1 := 10
	fmt.Printf("num1:%d\n", num1) //10 %d表示十进制占位
	fmt.Printf("num1:%b\n", num1) //十进制转二进制 1010 %b表示二进制
	fmt.Printf("num1:%o\n", num1) //十进制转八进制  %o表示八进制占位
	fmt.Printf("num1:%x\n", num1) //十进制转十六进制  %x表示十六进制占位
	//八进制 以0开头
	num2 := 077
	fmt.Printf("num2:%d\n", num2) //63 八进制转十进制
	//十六进制 以0x开头
	num3 := 0xFF
	fmt.Printf("num3:%d\n", num3) //255 十六进制转十进制
	//查看变量的类型
	fmt.Printf("num1 type:%T\n", num1) //%T表示变量类型
	//如何声明一个int8类型的变量
	i4 := int8(9) //明确指定类型 否则默认为int
	fmt.Printf("i4 type:%T\n", i4)
}

//浮点数 float32 float64 go中默认小数都是float64
//复数 complex64 complex128
//布尔 bool go中不允许将整型转换为布尔型 布尔型无法参与运算，也不能和其他类型去转换 默认值时false
//fmt占位符总结
func ffmt() {
	n := 100
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n) //%v表示以变量的值来输出
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	s := "hello"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", n)
	fmt.Printf("%#v\n", n) //%#v表示以变量的值来输出
}

//字符串 go中字符串必须用双引号包裹 "absc"，单引号包裹的时字符 'a',单独的字母，汉字，符号表示一个字符
//字节 1字节=8bit(8个二进制位)
//1个字符'A' = 1字节 = 8bit
//1个utf8字符'少' = 3字节 = 24bit
func stringg() {
	hello := "hello"
	world := " world"
	str := hello + world                      //可以这样对字符串进行拼接
	str1 := fmt.Sprintf("%s%s", hello, world) //也可以这样对字符串进行拼接
	fmt.Println(len(str), str1)               //获取字符串长度
	//分割 返回一个[]string	字符串切片
	ret := strings.Split(str, " ")
	fmt.Println(ret)
	//判断是否包含 返回bool
	fmt.Println(strings.Contains(str, "hello"))
	//判断前缀
	fmt.Println(strings.HasPrefix(str, "hello"))
	//判断后缀
	fmt.Println(strings.HasSuffix(str, "world"))
	//判断字串出现的位置
	fmt.Println(strings.Index(str, "h"))
	fmt.Println(strings.Index(str, "hello"))
	//按照某种形式把字符串切片拼接成字符串
	fmt.Println(strings.Join(ret, "-"))
}

//byte 字节类型 和rune类型
//组成每个字符串的元素叫做"字符"，可以通过遍历或者单个获取字符串元素获得字符
//字符用单引号'包裹起来
//GO中字符有以下两种
//1.rune类型 代表一个UTF-8字符
//2.uint8类型，或者叫byte类型 代表ASCII码的一个字符
func byteg() {
	s := "hello杭州"
	n := len(s) //len()函数返回byte字节的数量 11
	fmt.Println(n)

	// for i := 0; i < n; i++ {
	// 	fmt.Println(s[i])        //打印出的是ASCII码的值
	// 	fmt.Printf("%c\n", s[i]) //遍历输出字符串
	// 	//遇到汉字就会乱码了因为UTF-8编码比ASCII码多，所以找不到
	// }
	//而使用for range的方式便利 就可以得到UTF-8的字符
	for _, c := range s {
		fmt.Println(c)
		fmt.Printf("%c\n", c) //%c 表示输出字符
	}
	//总结对于英文 就用byte类型 对于中文和其他语言就用rune类型 rune实际是一个int32
}

//修改字符串
//字符串是不能直接修改的，需要先转成切片
func strin() {
	s1 := "白萝卜"
	//s1[0] = '红' 这种操作是不允许的 因为如果这样去是按照byte位取的，而一个汉字是三个字节，就没法改
	s2 := []rune(s1) //把字符串强制转换成了一个rune切片
	s2[0] = '红'
	fmt.Println(string(s2)) //再把rune切片转换成字符串
	c1 := "红"               //字符串 string
	c2 := '红'               //字符 int32 rune
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
}

//条件判断
func ifelse() {
	age := 19
	if age > 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}

	if sex := "man"; sex == "man" {
		fmt.Println("oo")
	} else {
		fmt.Println("xx")
	}
	//fmt.Println(sex) 在if else作用域内声明的变量 外面无法得到
}

//循环
func forf() {
	//基本格式
	for i := 0; i < 10; i++ {
		fmt.Println("1")
	}
	//变1
	j := 3
	for ; j < 10; j++ {
		fmt.Println("1")
	}
	//变2	这个其实可以充当while循环
	for j < 10 {
		fmt.Println("2")
		j++
	}
	//for range 便利数组 切片 字符串 map channel
	//数组 切片 字符串返回索引和值
	//map 返回键值对
	//channel 返回通道内传来的值
	s := []int{1, 2, 3, 4, 5}
	s1 := "hello杭州"
	for i, v := range s {
		fmt.Println(i, v)
	}
	for i, v := range s1 {
		fmt.Println(i, v)
		fmt.Printf("%d %c\n", i, v)
	}
	//break continue switch 都和其他语言相同
}

//位运算
func bit() {
	//& 与
	//| 或
	//^ 异或
	//<< 左移
	//>> 右移
	//&^ 按位取反
	a := 1                            //0000 0001
	b := 2                            //0000 0010
	fmt.Printf("%d %b\n", a&b, a&b)   //0000 0000 按位与
	fmt.Printf("%d %b\n", a|b, a|b)   //0000 0011 按位或
	fmt.Printf("%d %b\n", a^b, a^b)   //0000 0011 按位异或 两位不一样则为1一样则为0
	fmt.Printf("%d %b\n", a<<1, a<<1) //0000 0010 将二进制位左移一位
	fmt.Printf("%d %b\n", a>>1, a>>1) //0000 0000 将二进制位右移一位
	fmt.Printf("%d %b\n", a&^b, a&^b) //0000 0001 a&b^b 按位取反
}
