package kafaka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

type logData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer //声明一个全局的连接kafaka的生产者client
	logDataChan chan *logData
)

// Init 初始化client
func Init(addrs []string, size int) (err error) {
	config := sarama.NewConfig()
	//配置kafaka的连接
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	//连接kafaka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed,err:", err)
		return
	}
	logDataChan = make(chan *logData, size)
	go sendToKafaka()
	return
}

//发送消息给kafaka
func sendToKafaka() {
	for {
		select {
		case logdata := <-logDataChan:
			//构造一个发送给kafaka的消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = logdata.topic
			msg.Value = sarama.StringEncoder(logdata.data)
			//发送到kafaka
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, err:", err)
				return
			}
			fmt.Println(pid, offset)
		default:
			time.Sleep(200 * time.Microsecond)
		}
	}
}

//发送消息给chan
func SendToChan(topic, data string) {
	logDataChan <- &logData{
		topic,
		data,
	}
}
