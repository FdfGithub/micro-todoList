package model

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	db.LogMode(gin.Mode() != "release")
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetConnMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}