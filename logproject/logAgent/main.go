package logagent

import (
	"fmt"
	config "learn/logproject/logAgent/config"
	"learn/logproject/logAgent/etcd"
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
	err := kafaka.Init([]string{cfg.KafakaConf.Address}, cfg.KafakaConf.Size)
	if err != nil {
		return
	}
	fmt.Println("init kafaka success.")
	//2.初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("init etcd success.")
	//2.1 从etcd中拉取日志手机的配置信息
	logentries, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Get conf from etcd success.", logentries)
	//2.2 派一个哨兵取件事日志收集的变化
	//3. 收集日志发往kafka
	for _, logEntry := range logentries {
		err = tailog.Register(logEntry)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
