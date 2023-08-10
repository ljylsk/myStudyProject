package main

/*

 */

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//第二步：定义 一对多的表结构

// 2.1 定义一个User表
type User struct {
	gorm.Model
	// 用于指定数据库表中的列名。在这个例子中，该标签将数据库表中的列名设置为 "usernmae"
	Username string `json:"username" gorm:"column:username"`
	// 添加一个外键关联
	CreditCards []CreditCard
	// CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
}

// 2.2定义一个Card表
type CreditCard struct {
	gorm.Model
	Number string
	// UserRefer uint
	UserID uint //  这个就是与User表关联的外键  名字是 结构体+主键（gorm的规定）
}

func main() {
	// 0、连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// err = mil 证明连接正常，否则连接失败
	if err != nil {
		panic(err) //异常处理，有问题抛出异常
	}

	//第三步：创建表结构
	db.AutoMigrate(User{}, CreditCard{})

	//第四步
	//创建一对多
	// user := User{
	// 	Username: "zhangsan",
	// 	CreditCards: []CreditCard{
	// 		{Number: "0001"},
	// 		{Number: "0002"},
	// 	},
	// }

	// db.Create(&user)

	//给zhangsan添加信用卡
	u := User{Username: "zhangsan"}
	db.First(&u)
	// 添加新的关联： db.Model(&user).Association("Languages").Append([]Language{languageZH, languageEN})
	db.Model(&u).Association("CreditCards").Append([]CreditCard{
		{Number: "0003"},
	})

}
