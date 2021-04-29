package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Local())  // 获取本地时间
	fmt.Println(time.Now().Unix())   // 获取时间戳
	fmt.Println(time.Now().Date())   // 返回年月日
	fmt.Println(time.Now().Day())    // 返回日
	fmt.Println(time.Now().Year())   // 返回年
	fmt.Println(time.Now().Second()) // 返回秒
	fmt.Println(time.Now().Minute()) // 返回分钟
	fmt.Println(time.Now())          // 返回分钟

	now := time.Now()
	fmt.Println(now.Format("2006:01:02 15:04:05"))

}
