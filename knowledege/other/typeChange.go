package knowledege

import (
	"fmt"
	"strconv"
)

//字符串和数字 布尔相互转化
func Change() {
	//数字转化为字符串 方法一
	re1 := fmt.Sprintf("%d", 12)
	fmt.Printf("%#v", re1)
	//方法二 int转string
	re2 := strconv.Itoa(12)
	fmt.Printf("%#v", re2)
	//int64转string
	strconv.FormatInt(12, 10)

	//string转int
	re3, _ := strconv.Atoi("10")
	fmt.Println(re3)
	//string转int64
	strconv.ParseInt("12", 10, 64)
	//数字专string FormatInt string转数字 ParseInt
	//其他类型转string也是Format开头的函数 string转其他类型可是parse开头的函数

}
