package config

// [mysql]
var (
	SqlUserName string //= "root"
	SqlPassword string //= "123456"
	DataBase    string //= "videoWebsite"
	MysqlIP     string //= "192.168.239.126" //"192.168.1.107.3306" //ipconfig->WLAN->ipv4
)

// [RabbitMQ]
var (
	RabbitmqUserName string //= "admin"
	RabbitmqPassword string //= "123456"
	RabbitmqIP       string //= "192.168.136.128:5672"
	ExchangeName     string //= "direct"
)

// [redis]
var (
	RedisAddr     string //= "172.23.21.149:6379" //"redis:6379" //"192.168.1.100:6379"
	RedisPassword string //= "root"
	RedisDB       int    //= 0
)

// [elasticsearch]
var (
	EsAddr                 string //= "https://localhost:9200"
	EsName                 string //= "elastic"
	EsPassword             string //= "123456"
	CertificateFingerprint string //= "81bdbeefa2b378da4fd29dfd4d8e82e96cd0d92a50b24db5ec30401260fad917"
)

// [MinIo]
var (
	EndPoint        string //= "172.23.21.149:9000"
	AccessKeyID     string //= "minioadmin"
	SecretAccessKey string //= "minioadmin"
	SSL             bool   //= false
	BucketName      string //= "videoweb"
)

var (
	EtcdAddr string
)

var (
	JaegerAddr string
)

var (
	MilvusAddr string
)
