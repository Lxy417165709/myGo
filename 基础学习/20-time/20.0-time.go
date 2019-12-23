package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("开始了开始了!!")
	time.Sleep(1e9)	// 延迟1s
	fmt.Println("第一阶段~")
	time.AfterFunc(3e9, func() {
		fmt.Println("第二阶段...~")
	})
	time.Sleep(5e9)	// 延迟5s
	fmt.Println("第三阶段~")
	// 开始了开始了!!
	// 第一阶段~
	// 第二阶段...~
	// 第三阶段~

	// 总体的执行:  先"开始了开始了!!",
	// 				再1s后"第一阶段",
	// 				再3s后"第二阶段",
	// 				再2s后"第三阶段"

}
/*
	总结
	1. time模块的单位是纳秒,1秒 == 1e9纳秒

 */