package knowledege

import (
	"fmt"
	"time"
)

func Time() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	//时间戳转化为日期 第二个参数为偏移量 纳秒
	re := time.Unix(1297319, 0)
	fmt.Println(re)
	//时间间隔
	fmt.Println(time.Second)

	//now + 1 小时
	la := now.Add(24 * time.Hour)
	fmt.Println(la.Date())
	//now.Sub 求两个时间的差
	//now.Equal 求两个时间想不想等
	//now.Before now.After

	//定时器 本质上是一个channel
	//timer := time.Tick(time.Second)
	//用for循环来监听channel的信息
	// for t := range timer {
	// 	fmt.Println(t.Date())
	// }

	//时间格式化
	//年   月日 时分秒
	//2006 1 2 3 4 5
	//把时间对象转化为字符串
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))  //AM PM表示法
	fmt.Println(now.Format("2006-01-02 15:04:05.000")) //取到毫秒

	//按照对应的格式解析字符串类型的时间 返回Time
	t, _ := time.Parse("2006-01-02", "1998-01-19")
	fmt.Println(t.Unix())
	//按照某个时区进行解析
	lo, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}
	t, _ = time.ParseInLocation("2006-01-02", "1998-01-19", lo)
	fmt.Println(t.Unix())
}
