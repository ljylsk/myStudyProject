package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05")) //  2023-08-02 18:46:44
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan")) // 2021-06-13 01:10:18.143 PM Sun Jun
	fmt.Println(now.Format("2006/01/02 15:04"))                   // 2021/06/13 13:10
	fmt.Println(now.Format("15:04 2006/01/02"))                   // 13:10 2021/06/13
	fmt.Println(now.Format("2006/01/02"))                         // 2021/06/13
}
