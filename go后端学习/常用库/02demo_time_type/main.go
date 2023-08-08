// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now() //获取当前时间
// 	fmt.Printf("current time:%v\n", now)

// 	year := now.Year()     //年
// 	month := now.Month()   //月
// 	day := now.Day()       //日
// 	hour := now.Hour()     //小时
// 	minute := now.Minute() //分钟
// 	second := now.Second() //秒

// 	// 打印结果为：2021-05-19 09:20:06
// 	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
// }

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() //
	// 10分钟前
	//获取当前时间，并将当前时间减去10分钟，然后打印出减去10分钟后的时间。
	m, _ := time.ParseDuration("-1m")
	m1 := now.Add(m)
	fmt.Println("m1", m1)
	// 8个小时前
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(8 * h)
	fmt.Println(h1)
	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println(d1)

	// 10分钟后
	mm, _ := time.ParseDuration("1m")
	mm1 := now.Add(mm)
	fmt.Println(mm1)
	// 8小时后
	hh, _ := time.ParseDuration("1h")
	hh1 := now.Add(hh)
	fmt.Println(hh1)
	// 一天后
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	fmt.Println(dd1)

	// Sub 计算两个时间差
	subM := now.Sub(m1)
	fmt.Println(subM.Minutes(), "分钟") // 1 分钟
	sumH := now.Sub(h1)
	fmt.Println(sumH.Hours(), "小时") // 8 小时
	sumD := now.Sub(d1)
	fmt.Printf("%v 天\n", sumD.Hours()/24) // 1 天
}
