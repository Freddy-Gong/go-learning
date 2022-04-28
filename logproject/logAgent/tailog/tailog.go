package tailog

import (
	"context"
	"fmt"
	"learn/logproject/logAgent/etcd"
	"learn/logproject/logAgent/kafaka"
	"time"

	"github.com/hpcloud/tail"
)

type tailId int
type tailTask struct {
	path     string
	topic    string
	id       int
	instance *tail.Tail
	//为了能够停止goroutine
	ctx    context.Context
	cancel context.CancelFunc
}
type tailLogMgr struct {
	logEntry    *[]*etcd.LogEntry
	taskMap     map[tailId]*tailTask
	newConfChan chan []*etcd.LogEntry
}

var tailMgr = tailLogMgr{
	logEntry:    &logentries,
	taskMap:     tailMap,
	newConfChan: make(chan []*etcd.LogEntry, 0),
}
var tailMap = make(map[tailId]*tailTask, 10)
var logentries = make([]*etcd.LogEntry, 0, 10)

func Register(logEntry *etcd.LogEntry) (err error) {
	logentries = append(logentries, logEntry)
	var task tailTask
	task.init(logEntry.Id, logEntry.Path, logEntry.Topic)
	_, ok := tailMap[tailId(task.id)]
	if !ok {
		tailMap[tailId(task.id)] = &task
	} else {
		return fmt.Errorf("id 重复")
	}
	return
}

func (t *tailTask) init(id int, path, topic string) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tailObj, err := tail.TailFile(path, config)
	if err != nil {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	t.id = id
	t.path = path
	t.topic = topic
	t.instance = tailObj
	t.ctx = ctx
	t.cancel = cancel
	go t.run()
}

func (t *tailTask) run() {
	// for line := range tailog.ReadChan() {
	// 	kafaka.SendToKafaka("web_log", line.Text)
	// }
	//可以优化成下面的写法
	for {
		select {
		//读取日志
		case line := <-t.instance.Lines:
			//发送到kafaka
			//需要等待下面kafaka运行完之后 才能再次for循环 如何实现异步呢？
			//答：使用通道
			kafaka.SendToChan(t.topic, line.Text)
		case <-t.ctx.Done():
			fmt.Println("tailtask 退出了")
			return
		default:
			time.Sleep(time.Second)
		}
	}
}

//监听自己的newConfChan 有了新的配置过来就做对应的处理

func (t *tailLogMgr) run() {
	for newConf := range t.newConfChan {
		for _, conf := range newConf {
			v, ok := t.taskMap[tailId(conf.Id)]
			if ok {
				if v.path == conf.Path && v.topic == conf.Topic {
					continue
				} else {
					v.path = conf.Path
					v.topic = conf.Topic
				}
				//如果存在id应该进行对比
			} else {
				//不存在就新增
				Register(conf)
			}
		}
		//1. 配置新增
		//2. 配置删除 如果发现有删除操作 则就操作t.cancel就可以关掉了
		//3. 配置变更
		fmt.Println(newConf)
	}
}

//向外暴露一个函数，把newConfChan暴露出去
func NewConfChan() chan<- []*etcd.LogEntry {
	return tailMgr.newConfChan
}
