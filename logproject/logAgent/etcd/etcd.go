package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

//需要收集的日志的配置信息
type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
	Id    int    `json:"id"`
}

var (
	cli *clientv3.Client
)

//初始化ETCD的函数
func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	return
}

//从etcd中根据key获取配置
//可能存在多个配置项
func GetConf(key string) (logentries []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		return
	}
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &logentries)
		if err != nil {
			return
		}
	}
	return
}

//etcd watch
func WatchConf(key string, newConfCh chan<- []*LogEntry) {
	ch := cli.Watch(context.Background(), key)
	//从通道尝试去值
	for wresp := range ch {
		for _, evt := range wresp.Events {
			//把evt的值传给tail
			var newConf []*LogEntry
			if evt.Type != clientv3.EventTypeDelete {
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("unmarshal failed,err:%v\n", err)
					continue
				}
			}
			fmt.Printf("new COnf :%v\n", newConf)
			newConfCh <- newConf
		}
	}
}
