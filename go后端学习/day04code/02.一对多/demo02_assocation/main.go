package main

import (
	"encoding/json"
	"fmt"

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
	CreditCards []CreditCard //CreditCards只是一个名字（随便写）
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
	fmt.Println(db)

	// 第三步；查询关联表数据
	u := User{Username: "zhangsan"}
	db.First(&u) //查询单条数据
	// fmt.Printf("%v", u.Username)

	//"CreditCards"：关联外键的字段名
	err = db.Model(&u).Association("CreditCards").Find(&u.CreditCards)
	if err != nil {
		fmt.Println("err", err)
	}
	strUser, _ := json.Marshal(&u)
	fmt.Println(string(strUser))
}
