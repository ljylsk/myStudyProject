package main

import (
	"fmt"
	"time"
)

func main() {
	// test()
	//go 中开启一个协程go关键字  方法者
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main", i)
		time.Sleep(time.Microsecond * 100)
	}
	time.Sleep(time.Second)
	fmt.Println("over")
}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Microsecond * 100)
	}
}
