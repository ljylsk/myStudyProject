/*
结构体定义的三种方式：
1.如何定义一个结构体
2.初始化结构体的三种方法
*/

//第一种
// package main

// import "fmt"

// type person struct {
// 	name string
// 	city string
// 	age  int
// }

// func main() {
// 	// 结构体基本实例化
// 	var p1 person //实例化 person = P1
// 	p1.name = "张三"
// 	p1.city = "北京"
// 	p1.age = 18
// 	fmt.Printf("p1=%v\n", p1)  // p1={张三 北京 18}
// 	fmt.Printf("p1=%#v\n", p1) // p1=main.person{name:"张三", city:"北京", age:18}
// }

//第二种
// package main

// import "fmt"

// type person struct {
// 	name string
// 	city string
// 	age  int
// }

// func main() {
// 	//new 方法 返回的是一个指针类型
// 	var p2 = new(person)
// 	p2.name = "张三"
// 	p2.age = 20
// 	p2.city = "北京"
// 	fmt.Printf("p2=%#v \n", p2) //p2=&main.person{name:"张三", city:"北京", age:20}
// }

package main

import "fmt"

type person struct {
	name string
	city string
	age  int
}

func main() {
	// 键值对的初始化
	p4 := person{
		name: "zhangsan",
		city: "北京",
		age:  18,
	}
	// p4=main.person{name:"zhangsan", city:"北京", age:18}
	fmt.Printf("p4=%#v\n", p4)
}

/*
 1.结构体是go语言中实现类的一种方法
 go语言中没有类的概念
*/
