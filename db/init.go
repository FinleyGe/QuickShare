package db

import (
	. "QuickShare/config"
	. "QuickShare/log"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	devLog := logger.New(
		log.New(LogMultiWriter, "\r\n[GORM] ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			Colorful:                  false,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Info,
		})
	proLog := logger.New(
		log.New(LogMultiWriter, "\r\n[GORM] ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			Colorful:                  false,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Warn,
		})

	username := Config.Database.Username
	password := Config.Database.Password
	address := Config.Database.Address
	dbName := Config.Database.DBName
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, dbName)
	var err error

	if Config.Env == "dev" || Config.Env == "test" {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: devLog,
		})
	} else if Config.Env == "prod" {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: proLog,
		})
	}
	if err != nil {
		log.Panicln("Database Connection Error!", err)
	}

}
