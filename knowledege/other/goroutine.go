package knowledege

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

//goroutine和channel是go中并发的关键概念
//并发 同一时间段内 执行多个任务
//并行 同一时刻 执行多个任务
//Go并发通过goroutine 类似于线程 但是是属于用户态线程，比系统级线程更加轻量
//channel是在多个goroutine之间进行通信
func hello(i int) {
	defer wg.Done()
	fmt.Println("hello", i)

}

//它不能进行复制 所以需要在全局声明或者传指针
var wg sync.WaitGroup

func Gogo() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i) //开启一个单独的goroutine去执行hello函数
	}
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		fmt.Println(i) //这里i是一个闭包 会在执行的时候去外面找i，正确的i
	// 	}()
	// }
	fmt.Println("main")
	time.Sleep(time.Second)
	//main函数如果结束了 那么由main函数启动的goroutine也都结束了 就不会继续执行任务了
	//所以需要等待所有线程结束任务之后 才能返回main函数所以有了waitGroup
	wg.Wait()
}

func random() {
	//随机数种子 因为打包之后都是编译好的了，所以会出现结果相同的情况
	//所以需要一个随时变化的随机数生成依据 用时间生成这个随机数种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10) //0 <= x < 10
		fmt.Println(r1, r2)
	}
}

//GO中并行原理
//1. 可增长的栈 OS线程一般都有固定的栈内存（通常2MB），一个goroutine的栈在其生命周期开始的时候
//只有很小的栈（通常2KB），goroutine的栈不固定，特可以按需增大和缩小，goroutine的栈大小限制可以
//达到1GB
//2. goroutine调度 GMP模型
//G：就是goroutine，里面除了存放本goroutine信息外还有与所在P的绑定等信息
//M：是Go运行时（runtime）对操作系统内核线程的虚拟，M与内核线程一般是一一映射的关系
//一个goroutine最终都是要放到M上执行的
//P：管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界）
//P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停，运行后续的goroutinue
//等等）当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务

//因为GO是默认跑满CPU的，我们可以通过命令控制并发使用机器的核数
func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}
func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}
func RunProcess() {
	//指定使用机器的核数
	runtime.GOMAXPROCS(2)
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}

//goroutine间进行通信channel
//channel的定义
//var channel名 chan 通道中元素的类型
var ch1 chan int //nil

func Chan() {
	//channel是引用类型 是需要使用make进行初始化的 才能使用
	ch1 := make(chan int) //不带缓冲区的通道 通道中不能存值，如果没有变量接收就只能阻塞
	wg.Add(1)
	go func() {
		wg.Done()
		//将通道中的值 接收到x中
		x := <-ch1
		fmt.Println(x)
	}()
	//将一个值 发送到通道中
	//但是这里我们是一个没有缓冲区的channel，所以如果没有变量来接收这个值函数会阻塞
	//但是我们又不能只是单纯的在这行后面添加一个<-ch1
	ch1 <- 10
	//<-ch1 这样是不行的，因为在后面接收，前面还是阻塞的，是运行不到这里的
	//所以我们需要提前开一个goroutine来接收这个10，这样才能保证函数不阻塞
	ch1 = make(chan int, 10) //带缓冲区的通道 这个是可以在通道里存16个字节的值的
	ch1 <- 100               //只发送一个int不会阻塞当前channel所以会继续执行
	x := <-ch1
	fmt.Println(x)
	//关闭通道
	close(ch1)
	wg.Wait()
	//通道中最好不要传比较大的值 如果数据比较大 尽量传指针
}

//channel test
func ChannelTest() {
	//1.启动一个goroutine，生成100个数发送到ch1中
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1) //注意要把通道关闭掉
	}()
	//2.启动一个goroutine，从ch1中取值，计算其平方放到ch2中
	go func() {
		defer wg.Done()
		//如果ch1如果没有关闭，这个for循环就会一直跑就会死锁
		// for v := range ch1 {
		// 	ch2 <- v * 2
		// }
		// close(ch2)
		for {
			x, ok := <-ch1 //把通道关闭后 如果通道里面还有值 则还可以读取 如果没值了再去读的时候就返回false
			if !ok {
				break
			}
			ch2 <- x * 2
		}
		close(ch2) //注意要把通道关闭掉 要不然外面的for循环就会一直等死锁
	}()
	for v := range ch2 {
		fmt.Println(v)
	}
	wg.Wait()

}

//单向通道 多用在函数的参数里面一般都是限制函数中对通道的操作
func ff1(ch1 chan<- int, ch2 <-chan int) {
	ch1 <- 100 //在该函数中ch1只能存值
	//<-ch 这种操作是不允许的
	//ch2 <- 100 这种操作是不允许的
	<-ch2 //在该函数中ch2只能取值
}
