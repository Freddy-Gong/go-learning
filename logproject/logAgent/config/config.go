package config

type AppConf struct {
	KafakaConf `ini:"kafaka"`
	EtcdConf   `ini:"tailog"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key     string `ini:"collect_log_key"`
}

type KafakaConf struct {
	Address string `ini:"address"`
	Size    int    `int:"chan_max_siz"`
}

//unused
type TailogConf struct {
	FileName string `ini:"filename"`
}
