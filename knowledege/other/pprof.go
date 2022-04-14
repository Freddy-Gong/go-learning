package knowledege

//给系统运行情况照一个快照 性能快照
import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

//只是一段有问题的代码
func logicCode() {
	var c chan int //nil
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan, value:%v\n", v)
		default:
			//没有下面这行 就会在这个for循环空跑
			//加上这行就会让cpu休息 降低这个函数的cpu占用率
			//time.Sleep(500 * time.Microsecond)
		}
	}
}

func PProf() {
	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof, "cpu", false, "trun cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer func() {
			pprof.StopCPUProfile()
			file.Close()
		}()
	}
	for i := 0; i < 6; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)
	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(file)
		defer file.Close()
	}
}
