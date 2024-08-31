package cache

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"strconv"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

// Init 读取Redis配置
func Init() {
	file, err := ini.Load("./conf/conf.ini")
	if err != nil {
		log.Error(err)
		panic(err)
	}
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

// Redis 连接redis并且ping
func Redis() {
	db, err := strconv.ParseUint(RedisDbName, 10, 62)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		DB:   int(db),
	})
	_, err = client.Ping().Result()
	if err != nil {
		log.Error(err)
		panic(err)
	}
	RedisClient = client
	log.Info("connect redis success")
}
