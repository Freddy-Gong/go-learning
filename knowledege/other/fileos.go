package knowledege

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//操作文件
//读文件
//os.Open()函数能够打开一个文件，返回一个*File和一个err 对得到的文件实例调用
//close()方法能够关闭文件
func Read() {
	//这个路径要同步到，运行这个函数的路径
	//我这个函数是在/learn/knowledge/main.go文件下调用的
	//所以写相对路径的时候要相对于这个文件来写
	file, err := os.Open("./other/struct.go")
	if err == io.EOF {
		fmt.Println("读完了")
		return
	}
	if err != nil {
		fmt.Printf("Open file failed!, err:%v\n", err)
		return
	}
	defer file.Close()
	//读文件
	//temp := make([]byte, 128) //首先准备一个容纳文件内容的变量
	// for {                     //循环读取
	// 	n, err := file.Read(temp) //返回读取的字节数 如果文件的大小大雨你所设置的temp的大小的时候。
	// 	//就只会读取切片长度大小的文件内容，如果像读完整个文件 则需要进行循环。
	// 	//每一个循环读取128个字节的内容，循环完成后就可以读完了。
	// 	if err != nil {
	// 		fmt.Printf("read from file failed, err:%v", err)
	// 		return
	// 	}
	// 	fmt.Printf("读了%d个字节\n", n)
	// 	fmt.Println(string(temp[:n])) //因为读取的文件的大小不一定和temp的长度一致 所以需要用n来进行截取
	// 	if n < 128 {                  //当读取的内容没有填满这个切片时 就代表文件已经读完了
	// 		return
	// 	}
	// }

	//另一种更优雅的方法 bufio读取文件
	//先创建一个用来从文件中读取内容的对象
	// reader := bufio.NewReader(file)
	// for {
	// 	line, err := reader.ReadString('\n') //按行读取
	// 	if err == io.EOF {
	// 		return
	// 	}
	// 	if err != nil {
	// 		fmt.Printf("read line failed, err:%v", err)
	// 		return
	// 	}
	// 	fmt.Print(line)
	// }

	//更简单的方法 ioutil 不用自己操作文件句柄 相当于对一种方法的封装
	ret, err := ioutil.ReadFile("./other/function.go")
	if err != nil {
		fmt.Printf("read file failed,err:%v\n", err)
	}
	fmt.Println(string(ret))
}

//写文件
//os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入的相关功能
//打开方式：os.O_WRONLY 只写 os.O_CREATE 创建文件 os.O_RDONLY 只读
//os.O_RDWR 读写 os.O_TRUNC 清空 os.O_APPEND 追加
func Write() {
	//参数 文件名 打开方式 权限
	//打开方式这里使用的位运算的|,来对每一个打开方式进行添加
	//每一种打开方式的类型都是一个数字 和vue的diff优先级一样的操作
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	//defer fileObj.Close() 不可以写在这里
	if err != nil {
		fmt.Printf("open file failed err:%v\n", err)
		return
	}
	//这个defer一定要写在if后面 因为如果err确实存在的话 fileObj就是一个空值
	//如果把defer写在前面就会调用一个空值的close方法，就会panic 所以要写在后面
	defer fileObj.Close()
	//write
	fileObj.Write([]byte("xxxx\n"))
	//writeString
	fileObj.WriteString("xasd\n")

	//另一种方法写文件
	//创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	//调用方法把文件写到缓存里
	wr.WriteString("hello world!\n")
	//将缓存中的内容写入文件
	wr.Flush()

	//最简单的写文件方式
	//先定义一个字符串
	str := "hello util"
	//然后直接写到文件中
	err = ioutil.WriteFile("./xx.txt", []byte(str), 0644)
	if err != nil {
		fmt.Println("wirte file failed, err:", err)
		return
	}
}

//获取待用空格的用户输入
func Input() {
	var s string
	fmt.Scan(&s)
	fmt.Println(s)
	//这种获取输入的方式 无法获取被空格隔开的字符，只能分配给多个变量

	//可以通过这种方式来获取含有空格的输入
	reader := bufio.NewReader(os.Stdin) //读取标准输入的内容
	s, _ = reader.ReadString('\n')
	fmt.Println(s)
}
