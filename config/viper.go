package config

import (
	"log"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func ReadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	log.Println(os.Getwd())
	viper.AddConfigPath("config") //("../../config")这个是部署在docker的路径
	if err := viper.ReadInConfig(); err != nil {
		log.Println("failed to read config file")
		panic(err)
	}

	InitConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		klog.Infof("config file changed: %v\n", e.String())
	})

	viper.WatchConfig()

	log.Println("successfully running...!")
}

func InitConfig() {

	SqlUserName = viper.GetString("SqlUserName")
	SqlPassword = viper.GetString("SqlPassword")
	DataBase = viper.GetString("DataBase")
	MysqlIP = viper.GetString("MysqlIP") //"192.168.239.126" //"192.168.1.107.3306" //ipconfig->WLAN->ipv4

	RabbitmqUserName = viper.GetString("RabbitmqUserName")
	RabbitmqPassword = viper.GetString("RabbitmqPassword")
	RabbitmqIP = viper.GetString("RabbitmqIP")     //"192.168.136.128:5672"
	ExchangeName = viper.GetString("ExchangeName") //"direct"

	RedisAddr = viper.GetString("RedisAddr")         //"172.23.21.149:6379" //"redis:6379" //"192.168.1.100:6379"
	RedisPassword = viper.GetString("RedisPassword") //"root"
	RedisDB = viper.GetInt("RedisDB")                //0

	EndPoint = viper.GetString("EndPoint")               //"172.23.21.149:9000"
	AccessKeyID = viper.GetString("AccessKeyID")         //"minioadmin"
	SecretAccessKey = viper.GetString("SecretAccessKey") //"minioadmin"
	SSL = viper.GetBool("SSL")                           //false
	BucketName = viper.GetString("BucketName")           //"videoweb"

	EtcdAddr = viper.GetString("etcd.addr")

	JaegerAddr = viper.GetString("jaeger.addr")

	MilvusAddr = viper.GetString("milvus.addr")
}
