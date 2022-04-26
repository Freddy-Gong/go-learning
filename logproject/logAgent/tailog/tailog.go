package tailog

import (
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
}

var tailMap = make(map[tailId]tailTask, 10)

func Register(logEntry *etcd.LogEntry) (err error) {
	var task tailTask
	task.init(logEntry.Id, logEntry.Path, logEntry.Topic)
	_, ok := tailMap[tailId(task.id)]
	if !ok {
		tailMap[tailId(task.id)] = task
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
	t.id = id
	t.path = path
	t.topic = topic
	t.instance = tailObj
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
		default:
			time.Sleep(time.Second)
		}
	}
}
