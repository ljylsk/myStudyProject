package main // 声明 main 包，表明当前是一个可执行程序
import (
	"fmt"
)

// 导入内置 fmt 包

func main() { // main函数，是程序执行的入口
	// var num int = 100
	// num2 := 100
	// fmt.Println(num, num2) // 在终端打印 Hello World!

	// var c rune //主要用于存放字符,本质上存放ASCII码
	// c = '慕'
	// fmt.Printf("c=c%", c)

	num := 100
	fmt.Printf("%T %d \n", num, num)
}
