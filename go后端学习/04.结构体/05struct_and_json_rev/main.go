// 将string转化为结构体

package main

import (
	"encoding/json"
	"fmt"
)

// 1.定义一个结构体
type Student struct {
	ID      int
	Name    string
	Address string
	Age     int
}

func main() {
	//2.结构体对应的json字符串数据（string）
	s := `{"ID":1,"Name":"zhangsan"}`
	var stu Student
	//3.将string类型的数据转换为一个[]byte类型数据
	byteS := []byte(s)
	json.Unmarshal(byteS, &stu)
	fmt.Printf("%T %#v", stu, stu)
}
