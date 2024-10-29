package model

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var (
	DB *gorm.DB
)

func Database(connString string) {
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		TranslateError: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Error(err)
		panic(err)
		return
	}
	if err != nil {
		log.Error(err)
		panic(err)
		return
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //设置打开最大连接
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db
	log.Info("Connect database success")
	migration()
}
