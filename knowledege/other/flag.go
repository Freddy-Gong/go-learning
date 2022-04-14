package knowledege

import (
	"flag"
	"fmt"
	"os"
)

//flag 用来控制命令行参数

//os.Args 获取命令行参数
func Arg() {
	fmt.Println(os.Args)
	fmt.Println(os.Args[1])
}

//配置命令行参数 -name=asd
func FFlag() {
	//创建一个标识位
	// name := flag.String("name", "王也", "姓名")
	age := flag.Int("age", 10, "x")
	maried := flag.Bool("maried", false, "")
	var name string
	flag.StringVar(&name, "name", "xx", "xx") //这样就不用在后面
	//根据地址读值了
	//使用flag 必须要先解析
	flag.Parse()
	//fmt.Println(*name)
	fmt.Println(name)
	fmt.Println(*age)
	fmt.Println(*maried)

	fmt.Println(flag.Args())  //返回命令行参数后的其他参数 以[]string类型
	fmt.Println(flag.NArg())  //返回命令行参数后的其他参数
	fmt.Println(flag.NFlag()) //返回使用的命令行参数个数
}
