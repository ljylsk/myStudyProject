package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}
type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name, "开始工作")
}
func (p Phone) Stop() {
	fmt.Println("phone 停止")
}
func main() {
	phone1 := Phone{ // 一：实例化值类型
		Name: "小米手机",
	}
	var p1 Usb = phone1 //phone1 实现了 Usb 接口 phone1 是 Phone 类型
	p1.Start()

	phone2 := &Phone{ // 二：实例化指针类型
		Name: "苹果手机",
	}
	var p2 Usb = phone2 //phone2 实现了 Usb 接口 phone2 是 *Phone 类型
	p2.Start()          //苹果手机 开始工作
}

//值类型拷贝了很多份对象
