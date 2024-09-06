package conf

import (
	"chat/cache"
	"chat/model"
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	MongoDbClient *mongo.Client
	AppMode       string
	HttpPort      string
	Db            string
	DbHost        string
	DbPort        string
	DbUser        string
	DbPassWord    string
	DbName        string
	RedisDb       string
	RedisAddr     string
	RedisPw       string
	RedisDbName   string
	MongoDBName   string
	MongoDBAddr   string
	MongoDBPwd    string
	MongoDBPort   string
)

func Init() {
	file, err := ini.Load("./conf/conf.ini")
	if err != nil {
		log.Error(err)
		panic(err)
		return
	}
	LoadServer(file)
	LoadMysql(file)
	LoadMongoDb(file)
	MongoDb()
	dsn := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8", "&parseTime=true"}, "")
	model.Database(dsn)
	cache.Init()
	cache.Redis()
}

func MongoDb() {
	clientOptions := options.Client().ApplyURI("mongodb://" + MongoDBAddr + ":" + MongoDBPort)
	var err error
	mongoDbClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Error(err)
		panic(err)
		return
	}
	MongoDbClient = mongoDbClient
	log.Info("Connect MongoDB Success")
}

func LoadMongoDb(file *ini.File) {
	MongoDBName = file.Section("MongoDB").Key("MongoDBName").String()
	MongoDBAddr = file.Section("MongoDB").Key("MongoDBAddr").String()
	MongoDBPwd = file.Section("MongoDB").Key("MongoDBPwd").String()
	MongoDBPort = file.Section("MongoDB").Key("MongoDBPort").String()
}

func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}
