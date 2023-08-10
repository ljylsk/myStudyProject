package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
  第二步：定义表结构（模型定义）
*/

type User struct {
	// gorm:"primary_key" 主键索引，标记当前这个Id是自增的
	Id       int64 `json:"id" gorm:"primary_key"`
	Username string
	Password string
	Email    string
	Age      int
	Gender   string
}

func main() {
	/*
		如何去连接数据库
	*/
	// 第一步：创建gorm连接
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// err = mil 证明连接正常，否则连接失败
	if err != nil {
		panic(err) //异常处理，有问题抛出异常
	}

	//第三步：自动创建表结构
	db.AutoMigrate(User{})
}
