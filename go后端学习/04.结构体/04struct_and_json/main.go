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
	s := Student{
		ID:      1,
		Name:    "zhangsan",
		Address: "bj",
		Age:     24,
	}

	/*
		将结构体转化为json数据
		第一个值：sByte 是一个[]byte对象（将结构体转化出来的数据）
		第二个值：err
	*/

	// fmt.Printf("%T %#v", s, s)
	sByte, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err", err)
	}
	//将string(sByte)类型转换为string类型
	fmt.Println(string(sByte))
}
