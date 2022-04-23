package knowledege

import (
	"context"
	"fmt"
	"time"
)

var exitChan chan struct{} = make(chan struct{}, 1)

//如何优雅的控制子goroutine的开关
/*
1. 全局变量来控制
2. channel进行通信
*/
func example() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			fmt.Println("freddy")
			time.Sleep(time.Second)
			select {
			case <-exitChan:
				break //这里break只会跳出select 不会跳出for循环
			default:
			}
		}
	}()
	time.Sleep(time.Second * 5)
	exitChan <- struct{}{}
	wg.Wait()
}

//context要解决的问题就是避免层层嵌套带来的麻烦
func Context() {
	//context.Background()根结点
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
	Loop:
		for {
			fmt.Println("freddy")
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				break Loop
			default:
			}
		}
	}(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}

//这就是官方的方法，其实原理和使用channel是一样的。这个context方便的是
//当一个goroutine里面又开了一个goroutine，我们也可以吧ctx穿进去
//我们最外层的cancel函数依然可以是最里面那层的goroutine停止
/*
cancel源码
for child := range c.children {
	child.cancel(fasle,err)
}
*/

func deadline() {
	d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	//尽管ctx会过期，但在任何情况下吊用它的cancel函数都是很好的实践。
	//如果不这样做，可能会使上下文及其父类存活的实践超过必要的时间。
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("xx")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func timeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	defer cancel()
	wg.Add(1)
	go thiswork(ctx)
	time.Sleep(time.Second)
	wg.Wait()
}
func thiswork(ctx context.Context) {
LooP:
	for {
		fmt.Println("db connecting....")
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done(): //50毫秒后自动调用
			break LooP
		default:
		}
	}
	fmt.Println("worker done")
	wg.Done()
}
