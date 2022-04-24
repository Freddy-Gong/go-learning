package config

type AppConf struct {
	KafakaConf `ini:"kafaka"`
	TailogConf `ini:"tailog"`
}

type KafakaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type TailogConf struct {
	FileName string `ini:"filename"`
}
