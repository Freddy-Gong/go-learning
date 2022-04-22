package knowledege

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

func initproducer(str string) (err error) {
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(str, config)
	if err != nil {
		fmt.Printf("create producer failed, err:%v\n", err)
		return err
	}
	return nil
}

func Producter() {
	nsqAddress := "127.0.0.1:4150"
	err := initproducer(nsqAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin) //从标准输入中读数据
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		data = strings.TrimSpace(data)
		if strings.ToUpper(data) == "Q" {
			break
		}
		//向 topic_demo 推送数据
		err = producer.Publish("topic_demo", []byte(data))
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
