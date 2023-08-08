package main

import "fmt"

//结构体用法：用来模拟其他语言中的类
//函数调用方式：直接通过函数名字调用
//结构体方法调用：第一：初始化结构体  第二：通过结构体这个对象，方法名调用

//第一步：定义结构体（定义一个类）
type Person struct {
	name string
	age  int8
}

//第三步；给结构体绑定 方法和接收者
/*
(p Person)：接收者
*/
func (p Person) printInfo() {
	fmt.Printf("姓名:%v 年龄:%v", p.name, p.age) // 姓名:小王子 年龄:25
}

func main() {
	//第二步：实例化结构体（实例化一个类）
	p1 := Person{
		name: "小王子",
		age:  25,
	}

	//第四步：调用结构体的方法（通过类调用方法）
	p1.printInfo() // 姓名:小王子 年龄:25
}
