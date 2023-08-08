package main

import "fmt"

//1.定义一个结构体
type Person struct {
	name string
	age  int
}

//3.值类型接受者,绑定方法
func (p Person) printInfo() {
	fmt.Printf("姓名:%v 年龄:%v\n", p.name, p.age) // 姓名:小王子 年龄:25
}

//3.指针类型接收者
func (p *Person) setInfo(name string, age int) {
	p.name = name
	p.age = age
}
func main() {
	//2.实例化结构体  3种方法
	p1 := Person{
		name: "小王子",
		age:  25,
	}
	p1.printInfo() // 姓名:小王子 年龄:25
	p1.setInfo("张三", 20)
	p1.printInfo() // 姓名:张三 年龄:20
}
