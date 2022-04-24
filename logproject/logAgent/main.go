package logagent

import (
	"fmt"
	config "learn/logproject/logAgent/config"
	"learn/logproject/logAgent/kafaka"
	"learn/logproject/logAgent/tailog"
	"time"

	"gopkg.in/ini.v1"
)

var cfg = new(config.AppConf) //new函数返回的是指针

func main() {
	//0. 加载配置文件
	// cfg, err := ini.Load("./config.ini")
	// address := cfg.Section("kafaka").Key("address").String()
	// logfile := cfg.Section("tailog").Key("path").String()
	// topic := cfg.Section("kafaka").Key("topic").String()
	ini.MapTo(cfg, "./config/config.ini")
	//1. 初始化kafaka连接
	err := kafaka.Init([]string{cfg.KafakaConf.Address})
	if err != nil {
		return
	}
	fmt.Println("init kafaka success.")
	//2.打开日志文件准备收集日志
	err = tailog.Init(cfg.TailogConf.FileName)
	if err != nil {
		return
	}
	fmt.Println("init tailog success.")
	run()
}

func run() {
	// for line := range tailog.ReadChan() {
	// 	kafaka.SendToKafaka("web_log", line.Text)
	// }
	//可以优化成下面的写法
	for {
		select {
		//读取日志
		case line := <-tailog.ReadChan():
			//发送到kafaka
			kafaka.SendToKafaka(cfg.KafakaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}

}
