package config

type Config struct {
	Redis    RedisConfig    `yaml:"redis"`
	Mysql    MysqlConfig    `yaml:"mysql"`
	Nacos    NacosConfig    `yaml:"nacos"`
	RabbitMQ RabbitMQConfig `yaml:"rabbitmq"`
	OSS      OSSConfig      `yaml:"oss"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Db   int    `yaml:"db"`
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
	Port     int    `yaml:"port"`
	Hostname string `yaml:"hostname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type OSSConfig struct {
	BucketName      string `ymal:"bucketName"`
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
}
