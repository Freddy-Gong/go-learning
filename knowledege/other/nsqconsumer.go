package knowledege

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

//MyHandler是一个消费者模型
type MyHandler struct {
	Title string
}

//HandleMessage 是需要实现的处理消息的方法
func (m *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	fmt.Printf("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	return
}

//初始化消费者
func initConsumer(topic string, channel string, address string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	consumer := &MyHandler{
		Title: "freddy",
	}
	c.AddHandler(consumer)
	//if err := c.ConnectToNSQD(address);err != nil { 直接连NSQD
	if err := c.ConnectToNSQLookupd(address); err != nil { //连接lookup进行查询
		return err
	}
	return nil
}

func Consumer() {
	err := initConsumer("topic_demo", "first", "127.0.0.1:4161")
	if err != nil {
		fmt.Println(err)
		return
	}
	c := make(chan os.Signal)        //定义一个信号的通道
	signal.Notify(c, syscall.SIGINT) //转发键盘终端信号到c
	<-c                              //阻塞
}
