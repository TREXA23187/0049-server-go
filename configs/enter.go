package configs

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
	Jwt    Jwt    `yaml:"jwt"`
	Redis  Redis  `yaml:"redis"`
	MQ     MQ     `yaml:"mq"`
	GRPC   GRPC   `yaml:"grpc"`
	QiNiu  QiNiu  `yaml:"qi_niu"`
}
