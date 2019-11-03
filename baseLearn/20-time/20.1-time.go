package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 获取当前时间
	t1 := time.Now()
	fmt.Println(t1)
	// 2019-06-15 13:11:20.6608221 +0800 CST m=+0.022937901

	// 2. 根据时间戳创建时间
	t2 := time.Unix(0, t1.UnixNano())
	fmt.Println(t2)
	// 2019-06-15 13:11:20.6608221 +0800 CST

	// 3. 自定义创建时间
	t3 := time.Date(2050, 12, 12, 15, 16, 17, 0, time.Local)
	fmt.Println(t3)
	// 2050-12-12 15:16:17 +0800 CST

	// 4. 获取日期(年,月,日)
	t4 := time.Now()
	year, month, day := t4.Date()                          // 合并获取
	year1, month1, day1 := t4.Year(), t4.Month(), t4.Day() // 分开获取
	fmt.Println(t4.Date())
	fmt.Println(year, month, day)
	fmt.Println(year1, month1, day1)
	fmt.Println(int(month)) // month的类型实际是Month,但是Month来自于type Month int
	// 2019 June 15
	// 2019 June 15
	// 2019 June 15
	// 6

	// 5. 获取时钟(时,分,秒,纳秒)
	t5 := time.Now()
	hour, minute, second := t5.Clock()
	hour1, minute1, second1 := t5.Hour(), t5.Minute(), t5.Second()
	nano := t5.Nanosecond()
	fmt.Println(t5.Clock())
	fmt.Println(hour, minute, second)
	fmt.Println(hour1, minute1, second1)
	fmt.Println(nano)
	// 13 20 43
	// 13 20 43
	// 13 20 43
	// 980319900

	// 6. 获取时间差
	t6 := time.Now()
	fmt.Println(t6.Unix())     // 秒时间戳
	fmt.Println(t6.UnixNano()) // 纳秒时间戳
	// 1560576244
	// 1560576244011464600

	// 7. time与string的相互转化
	// 7.0 time转string
	t7 := time.Now()
	s := t7.Format("2006:01:02 15:04:05") // 这是固定的~
	fmt.Println(t7)
	fmt.Println(s)
	// 2019-06-15 13:27:00.0846912 +0800 CST m=+0.079786701
	// 2019:06:15 13:27:00

	// 7.1 string转time
	s1 := "2050/02/04 19:5:7"
	//t8,_ := time.Parse("2006/01/02 15:04:05",s1)	// 注意这样写是错误的,因为上面是19:5:7
	t8, _ := time.Parse("2006/1/2 15:4:5", s1) // 月日有无前导0都无所谓
	fmt.Println(t8)
	// 2050-02-04 19:05:07 +0000 UTC


	// 7.2 time运算
	t9 :=time.Now()
	fmt.Println(t9)
	fmt.Println(t9.Add(-1e11))
	fmt.Println(t9.AddDate(1,7,20))
	fmt.Println(int((t9.Add(-1e11)).Sub(t9)))
	// 2019-06-15 13:46:52.9241567 +0800 CST m=+0.065823201
	// 2019-06-15 13:45:12.9241567 +0800 CST m=-99.934176799
	// 2021-02-04 13:46:52.9241567 +0800 CST
	// -100000000000
}

/*
	总结
	1. "2006/01/02 15:04:05" 是go语言时间格式的模板
	2. string转time时,模板月日有无前导零无所谓,但是模板分秒是否需要前导零取决于string的内容中的分秒


*/
