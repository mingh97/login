package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"loginRegister/global"
	"loginRegister/model"
	"time"
)

func initDB() {
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("connect database failed, err:%v\n", err)
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("connect database failed, err:%v\n", err)
	}

	global.DB = db
	log.Println("connect database success")

	err = global.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("auto migrate failed, err:%v\n", err)
	}
	log.Println("auto migrate success")
}
