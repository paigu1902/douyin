package config

type Config struct {
	Redis    RedisConfig    `yaml:"redis"`
	Mysql    MysqlConfig    `yaml:"mysql"`
	Nacos    NacosConfig    `yaml:"nacos"`
	OSS      OSSConfig      `yaml:"oss"`
	RabbitMQ RabbitMQConfig `yaml:"rabbitmq"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       int    `yaml:"db"`
	Password string `yaml:"password"`
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type NacosConfig struct {
	Host string `yaml:"host"`
	Port uint64 `yaml:"port"`
}

type RabbitMQConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Hostname string `yaml:"hostname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
}

type OSSConfig struct {
	BucketName      string `ymal:"bucketName"`
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
}
