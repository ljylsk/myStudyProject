package main

import "fmt"

/*
channel是不同的goroutine进行数据传递，类似于队列
*/

func main() {
	// 1.定义一个channel
	//make可以给切片、map、channel分配内存
	/*
		chan：定义channel的关键字
		int：指定当前这个channel的数据类型
		5：定义这个channel的大小
	*/
	// ch := make(chan int, 5)

	// //2.向channel存入数据
	// ch <- 10

	// //3.从channel中取数据
	// v1 := <-ch
	// fmt.Println("v1", v1)

	//通过for range 从channel中取值
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5

	close(ch)
	//当循环一个没有close的channel，而且还为空就报错
	for i := range ch {
		fmt.Println(i)
	}
}
