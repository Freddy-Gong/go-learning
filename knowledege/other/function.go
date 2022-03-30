package knowledege

import "fmt"

func Function() {
	/*func 函数名[范型](参数 参数类型)(返回值名 返回值类型){
		 函数体
		 return
	}
	返回值可以命名也可以不命名 命名的时候可以直接return 返回值
	多个参数类型相同时 可以简写
	可变长参数
	func test(arg1 string, arg2...int){}
	表示arg2可以传多个int值
	可变长参数必须放在最后
	*/
	//函数里面不能再定义一个具名函数 但是匿名函数是可以的
	//defer语句  函数结束前才去运行defer后面跟随的语句
	func() { //1 4 3 2
		//defer 存语句是用栈储存的 后进先出
		fmt.Print("1")
		defer fmt.Println("2")
		defer fmt.Println("3")
		fmt.Print("4")
	}()
	//defer语句确切执行时机
	//Go中函数的return不是原子操作，在底层是分为两步的
	//第一步给返回值赋值
	//defer
	//第二步真正RET返回
	//函数中如果存在defer，defer就是在第一步和第二步之间执行的
	fmt.Println("1", func() int {
		x := 5
		defer func() { x++ }()
		return x
	}()) //5 因为返回值没有命名 在第一步就把x的值赋值给了返回值，虽然defer修改了x但是没有修改返回值 所以返回5
	fmt.Println("2", func() (x int) {
		defer func() { x++ }()
		return 5
	}()) //6
	fmt.Println("3", func() (y int) {
		x := 5
		defer func() { x++ }()
		return x
	}()) //5
	fmt.Println("4", func() (x int) {
		defer func(x int) {
			x++
		}(x)
		return 5
	}()) //5 这里是因为传给函数的时候函数里面的x是x的拷贝，所以不会改变传入的x
	fmt.Println("5", func() (x int) {
		defer func(x *int) {
			(*x)++
		}(&x)
		return 5
	}()) //6 这里是因为穿的是地址所以把x给改了

	//defer面试题
	/*
		func calc(index string,a,b int) int {
			ret := a + b
			fmt.Println(index,a,b,ret)
			return ret
		}
		func main(){
			a := 1
			b := 2
			defer calc("1",a,calc("10",a,b)) 在defer之前，因为参数里面有函数，所以要先把函数运行得到的结果填到defer中
			并且defer 回保存当前的变量的值，所以 第一个defer执行的时候 a 还是等于 1 的，而不是调用的时候a的值 也是闭包
			a = 0
			defer calc("2",a,calc("20",a,b))
			b = 1
		}
	*/

	//函数的类型会把自己的参数的类型和返回值的类型都带上
	//函数也可以最为参数传进函数里
	//函数也可以作为返回值 所以要写它的类型
	func(f func(int) int) {
	}(func(x int) int { return 8 })

	//闭包 和js的闭包效果和原理都是一样的
	//闭包其实就是一个函数和它所使用的外部变量的集合，因为函数里面的参数的值，所在定义的时候
	//就已经确定了，而不在于在哪里运行，重要的时在定义的时候
	//闭包的例子1
	ret := f3(f2, 100, 200)
	f1(ret)
	//例子2 函数定义在下面
	f1, f2 := calc(10)
	//base一直在被修改
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}

//闭包的例子1
//加入f1是一个借口，但是你现在需要把f2传进去，但由于类型不匹配
//所以只能通过包装一个f3得到最后的结果。如下
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}
func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}
func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

//闭包例子2
func calc(base int) (func(int) int, func(int) int) {
	add := func(int) int {
		base += 1
		return base
	}

	sub := func(int) int {
		base -= 1
		return base
	}
	return add, sub
}

//panic recover
//panic调用之后会使得程序崩溃 recover回尝试恢复这个程序
//1. recover()必须搭配defer使用
//2. defer一定要在可能引发panic的语句之前定义
func panic_recover() {
	//刚刚打开数据库连接
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("释放数据库连接")
		}
	}()
	panic("出错了")
	fmt.Println("b")
}

//fmt包详解
func ffmt() {
	//输出
	fmt.Print("asc")
	fmt.Printf("%d", 10)
	fmt.Println("asd")
	//获取输入
	fmt.Scan() //从终端中获取输入
	var s string
	fmt.Scan(&s)
	fmt.Println("用户输入的内容是: ", s)
	//fmt.Scanf() //从终端中获取固定格式的内容
	var (
		name  string
		age   int
		class string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Println(name, age, class)
	fmt.Scanln() //自动扫描到换行符，最后读输入
}
