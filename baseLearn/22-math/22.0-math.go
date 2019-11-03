package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// 1. 常用数学函数
	// 1.0 取整
	fmt.Println(math.Floor(-5.7)) // 向下取整 (即小于等于)
	fmt.Println(math.Ceil(5))     // 向上取整 (即大于等于)
	// -6
	// 5

	// 1.1 四舍五入
	fmt.Println(math.Round(7.884)) // 四舍五入
	// 8

	// 1.2 取绝对值
	fmt.Println(math.Abs(-7.5555)) //取绝对值
	// 7.5555

	// 1.3 取最大/最小值
	fmt.Println(math.Max(5.8, 100)) // 取最大值
	fmt.Println(math.Min(5.8, 100)) // 取最小值
	// 100
	// 5.8

	// 1.4 取次方
	fmt.Println(math.Pow(5, 1.2))
	// 6.898648307306074

	// 1.5 分离整数和小数
	x, y := math.Modf(-1.789)
	fmt.Println(x, y)
	// -1 -0.7889999999999999


	// 2.随机数
	// 2.0 不置随机种子(每次结果都一样)
	fmt.Println(rand.Int63n(100))
	fmt.Println(rand.Int63n(100))
	fmt.Println(rand.Int63n(100))
	// 10
	// 51
	// 21

	// 2.1 置随机种子(每次结果都随机)
	rand.Seed(time.Now().Unix())	// 把秒时间戳作为种子
	fmt.Println(rand.Int63n(100))
	fmt.Println(rand.Int63n(100))
	fmt.Println(rand.Int63n(100))
	// 84
	// 98
	// 54


}

/*
	总结
	1. 取随机数要置随机种子,不然每次随机都是一样的~
	2. rand.Int63n(100)表示取[0,100)中的整数
*/