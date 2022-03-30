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
	}()) //5 这里是因为传给函数的时候函数里面的x是拷贝的x，所以不会改变传入的x
}
