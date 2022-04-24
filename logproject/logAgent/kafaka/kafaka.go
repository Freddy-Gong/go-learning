package kafaka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

var client sarama.SyncProducer //声明一个全局的连接kafaka的生产者client
// Init 初始化client
func Init(addrs []string) (err error) {
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
	return
}

//发送消息给kafaka
func SendToKafaka(topic, data string) {
	//构造一个发送给kafaka的消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	//发送到kafaka
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Println(pid, offset)
}
