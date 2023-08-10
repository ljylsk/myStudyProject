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
}

func main() {
	/*
		如何去连接数据库
	*/
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	// 1.创建表结构
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// err = mil 证明连接正常，否则连接失败
	if err != nil {
		panic(err) //异常处理，有问题抛出异常
	}
	// fmt.Println(db)
	// fmt.Println(err)

	//第三步：自动创建表结构
	// db.AutoMigrate(User{})

	// 第四步：增删改查
	// 4.1 增：向User表添加一条数据
	db.Create(&User{
		// Id:1 , // 因为Id设置了自增，所以无需传递
		Username: "lisi",
		Password: "123456",
	})

	//4.2 改：修改表的某一个字段
	// db.Model(User{
	// 	Id: 1,
	// }).Update("password", "123456")

	//4.3 查询
	// 过滤单条数据:First
	// U := User{Id: 1}
	// db.First(&U)
	// fmt.Printf("%#v", U)

	// 查询所有数据
	// users := []User{} //定义一个User结构体的切片来接收
	// db.Find(&users)
	// fmt.Printf("%#v", users)

	// //  4.4 删除
	// // 根据主键删除
	// db.Delete(&User{Id: 1})

	// // 条件删除
	// db.Where("username = ?", "lisi").Delete(&User{})

}
