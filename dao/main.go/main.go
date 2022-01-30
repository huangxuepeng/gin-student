package main

import (
	"fmt"
	"log"
	"os"
	"student/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB //全局变量
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/huangxuepeng?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "/r/n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //Log level
			Colorful:      false,       //禁用彩色打印
		},
	)

	//全局模式
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	// //迁移 schema
	// if err = DB.AutoMigrate(&model.Student{}); err != nil {
	// 	zap.S().Infof("数据库表关联失败!")
	// }
	Create()
	// Save()
}

func Create() {
	fmt.Println("插入数据:")
	var user model.Student
	user.Name = "黄雪朋"
	user.Mobile = "18088630924"
	user.Password = "1234567890"
	DB.Select("Name", "Mobile", "Password").Create(&user)
}

func chaxun() {
	var user model.Student
	DB.First(&user, 10)
	fmt.Println(user.Name)
}

func Save() {
	var user model.Student
	DB.Model(&user).Where("Mobile=?", "18088630924").Update("name", "黄雪朋")
}
