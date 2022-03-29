package main

import "fmt"

func knowledege() {

}

//Array数组 slice切片
func array() {
	//数组的定义 必须指定数组的存放的元素的类型和容量
	//var arr [3]int
	//数组如果不进行初始化 则默认为0值
	//初始化方式1
	arr := [3]int{1, 2, 3}
	//初始化方式2 可以不指定数组的容量 根据初始值的个数来计算容量
	arr2 := [...]int{1, 2, 3}
	//初始化方式3 根据索引来初始化
	arr3 := [5]int{1: 1, 3: 2} // [0,1,0,2,0]
	fmt.Println(arr, arr3)
	//数组的遍历
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	for i, v := range arr2 {
		fmt.Println(i, v)
	}

	//多维数组初始化
	//[[1,2,3],[4,5,6]]
	a11 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(a11)
	//数组是在内存中一段连续的内存空间 是一个值类型
	//所以在把一个数组负值给另一个变量的时候进行了拷贝
	a12 := a11
	//因为是值类型所以可以使用 == 和 != 来进行比较
	fmt.Println(a12 == a11)
}
func slice() {
	arr3 := [5]int{1: 1, 3: 2}
	//----------------------------------------------------------
	//数组的切片则是一个引用类型，指向数组的某个部分
	//切片就是一个框，框住一块连续的内存空间
	//切片的定义声明
	var s1 []int //[] s1==nil 没有开辟内存空间
	//切片的初始化
	s2 := []int{1, 2, 3}
	fmt.Println(s2)
	s3 := []int{} //[] s3!=nil
	fmt.Println(s1, s3, s1 == nil, s3 == nil, len(s1), len(s3), cap(s1), cap(s3))
	//由数组得到切片
	s4 := arr3[1:3] //基于一个数组切割，左包含右不包含 取出arr3索引位1 2的元素
	fmt.Println(s4)
	s5 := arr3[1:] //从索引1切到最后 切片的容量等于切片第一个元素的位置到数组的末尾 cap(s5)=4
	s6 := arr3[:3] //从索引0切到索引2 cap(s6)=5
	s7 := arr3[:]  //全切了 cap(s7)=5
	fmt.Println(s5, s6, s7, cap(s5), cap(s6), cap(s7))
	//切片的遍历方式和数组相同
	//切片的也可以再次被切片
	s8 := s7[1:3]
	fmt.Println(s8)
	//使用make函数创建切片 可以预先规定容量和长度 make其实就是分配内存
	s9 := make([]int, 3, 5) //切片的容量是5
	fmt.Println(s9, len(s9), cap(s9), s9 == nil)
	//append() 函数可以在切片的末尾追加元素
	//调用append函数必须用原来的切片变量接受返回值
	s9 = append(s9, 1, 2, 3) //append后如果需要扩容，内存地址就会变的 如果不用变量接受返回值 引用就丢了
	fmt.Println(s9, cap(s9))
	s9 = append(s9, s1...) //...表示将切片的元素一个一个添加到切片中 和JS相同
	fmt.Println(s9, cap(s9))

	//copy 深拷贝
	s10 := s9                            //浅拷贝 复制了地址
	s11 := make([]int, len(s9), cap(s9)) //必须要先分配内存 才能进行深拷贝
	copy(s11, s9)
	s9[0] = 1000
	fmt.Println(s9, s10, s11) //s11没有受到s9的影响

	//从切片中删除元素 没有直接的方法
	s12 := []int{1, 2, 3, 4, 5}
	//删除索引为2的元素 ，只能这样通过切片的组合实现
	s12 = append(s12[:2], s12[3:]...)
	fmt.Println(s12, cap(s12)) //但是容量还是底层数组的容量
	//这种操作会对底层数组产生影响 举一个例子
	x1 := [...]int{1, 3, 5}
	s13 := x1[:]
	fmt.Println(s13, len(s13), cap(s13))
	//1. 切片不保存具体的值
	//2. 切片对应一个底层数组
	//3.底层数组都是占用一块连续的内存
	fmt.Printf("%p\n", s13)           //%p 打印内存地址
	s13 = append(s13[:1], s13[2:]...) //这一步是会修改底层数组的值
	fmt.Printf("%p\n", s13)           //没有重新创建
	fmt.Println(s13, len(s13), cap(s13))
	fmt.Println(x1) //[1, 5, 5] append 把中间的三删掉了后
	//用后面的元素依次补上了，因为slice必须是连续的，所以会出现这种情况
}

//指针
func pointer() {
	//1. & 取地址符
	//2. * 取值符
	n := 19
	p := &n
	fmt.Println(p, *p)    //0xc0000a0a0 19
	fmt.Printf("%T\n", p) //*int int类型的指针
}

//new make
func newMake() {
	//var a *int   //指针类型 但是此时是没有分配内存的，所以是nil
	//a = new(int) //new 分配内存 这个时候才有内存地址 才可以执行后续操作
	//make也是分配内存的 它只适用于slice map channel ，且他会返回类型本身，而不返回指针类型
	//new很少用 一般用来给基本数据类型申请内存string int  返回的是对应类型的指针 *string *int
}

//map
func mmap() {
	//ma是引用类型 所以必须先初始化才能使用
	//var m map[string]int      //没有分配内存 nil 这个时候不能使用
	a := make(map[string]int, 10) //分配了内存 应该避免动态扩容
	a["a"] = 1
	a["b"] = 2      //通过键值对的方式赋值和取值
	v, ok := a["a"] //通过ok来判断是否存在对应的值
	if !ok {
		fmt.Println("not found")
	} else {
		fmt.Println(v)
	}
	for k, v := range a {
		fmt.Println(k, v)
	}
	for k := range a { //只便利key
		fmt.Println(k)
	}
	for _, v := range a { //只便利value
		fmt.Println(v)
	}
	//删除键值对
	delete(a, "a")
	delete(a, "c") //当删除一个不存在的键的时候什么也不做
	//按照指定顺序便利map
	//1. 取出map中的所有的kep 存入切片中
	//2. 对切片进行排序
	//3. 按照切片的顺序便利map

	//元素为map的切片
	s := make([]map[string]int, 10)
	//此时我们只是初始化了一个切片，但是没有初始化map元素
	s[0] = make(map[string]int, 10)
	s[0]["a"] = 1 //这时才能进行赋值

	//元素为切片的map
	m := make(map[string][]int, 10)
	//同样此时只是对map进行了分配内存
	m["a"] = []int{1, 2, 3}
}

//函数
func function() {
	/*func 函数名[范型](参数 参数类型)(返回值 返回值类型){
		 函数体
	}
	返回值可以命名也可以不命名 命名的时候可以直接return 返回值
	多个参数类型相同时 可以简写
	可变长参数
	func test(arg1 string, arg2...int){}
	表示arg2可以传多个int值
	可变长参数必须放在最后
	*/
}
