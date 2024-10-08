package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	DB *gorm.DB
)

func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		log.Error(err)
		panic(err)
		return
	}
	db.LogMode(true)
	if err != nil {
		log.Error(err)
		panic(err)
		return
	}
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)       //默认不加复数s
	db.DB().SetMaxIdleConns(20)  //设置连接池，空闲
	db.DB().SetMaxOpenConns(100) //设置打开最大连接
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	log.Info("Connect database success")
	migration()
}
