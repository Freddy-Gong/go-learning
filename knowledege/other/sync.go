package knowledege

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

//并发安全 锁🔒
//例子 多个线程同时访问一个全局变量
var x = 0
var lock sync.Mutex

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
	wg.Done()
}
func low() {
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x--
		lock.Unlock()
	}
	wg.Done()
}
func Stest() {
	wg.Add(2)
	//多个线程操作同一个数据
	go add()
	go low()
	wg.Wait()
	fmt.Println(x)
}

//虽然互斥锁能够保证在同一时间只有一个goroutine访问数据，但是
//如果读的情况远大于写的情况 这样做就是非常浪费资源 所以引出了读写互斥锁这个概念
//读写互斥锁
//它分为两种 当一个线程获取了一个读锁，其他读的线程也能继续访问
//如果当前线程获取了一个写锁，其他线程无法访问
var rw sync.RWMutex

func read() {
	//lock.Lock()
	rw.RLock()
	fmt.Println(x)
	time.Sleep(200 * time.Microsecond)
	//lock.Unlock()
	rw.RUnlock()
	wg.Done()
}

func write() {
	//lock.Lock()
	rw.Lock()
	x++
	time.Sleep(1000 * time.Microsecond)
	//lock.Unlock()
	rw.Unlock()
	wg.Done()
}
func Wtest() {
	wg.Add(1100)
	now := time.Now()
	for i := 0; i < 100; i++ {
		go write()
	}
	for i := 0; i < 1000; i++ {
		go read()
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(now))
}

//sync.once 确保某一个操作只执行一次 比如关闭channel 加载配置文件
//例子
var once sync.Once //接收一个函数 没有参数没有返回值 所以
//当参数不满足这个条件的时候穿入一个闭包
func f11() {

}
func f22(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	once.Do(func() { close(ch2) })
}
func Otest() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f11()
	go f22(a, b)
	go f22(a, b)
	wg.Wait()
	for result := range b {
		fmt.Println(result)
	}
}

//sync.Map
//Go内置的map不是并发安全的
//如果我们使用多线程同时修改一个map就会出现错误
//所以需要加🔒，但是手动实现很复杂 所以Go中内置了一种并发安全的map
var m = make(map[string]int)

func get(key string) int {
	return m[key]
}
func set(key string, value int) {
	m[key] = value
}
func Mtest() {
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			lock.Lock()
			set(key, n)
			lock.Unlock()
			fmt.Printf("k=%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//不用进行make操作
var m2 sync.Map

//内置 store写值 load读值 loadorstroe 读写值 delete删除 range便利
func Mtest2() {
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)         //必须使用sync.Map内置的方法store来存值
			value, _ := m2.Load(key) //必须使用 Load去读值
			fmt.Printf("k=%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//原子操作 atomic
//从语言层面实现的🔒用来代替上锁和解锁的操作
var y int64

func dadd() {

	// lock.Lock()
	// x += 1
	// lock.Unlock()
	atomic.AddInt64(&y, 1)
	//还有其他的读和比较等操作
}
