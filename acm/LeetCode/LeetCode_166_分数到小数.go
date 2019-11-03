package main

import (
	"fmt"
	"math"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	m := make(map[int]int)
	x := float64(numerator) / float64(denominator) // 判断正负
	flag := 0
	if x == 0 {
		return "0"
	} else {
		if x < 0 {
			flag = 1
		}
	}

	numerator = int(math.Abs(float64(numerator)))
	denominator = int(math.Abs(float64(denominator)))

	zs := numerator / denominator	// 整数部分
	numerator = numerator % denominator
	yushu := -1
	ans := ""

	for i := 1; yushu != 0; i++ {
		yushu = numerator % denominator

		s := strconv.Itoa(numerator / denominator)	// 保证s长度==1
		ans += s
		if m[yushu] == 0 {
			m[yushu] = i
		} else {
			ans = strconv.Itoa(zs) + "." + ans[1:m[yushu]] + "(" + ans[m[yushu]:] + ")"
			break
		}
		numerator = yushu * 10
	}

	if yushu == 0 {
		if ans[1:] != "" {
			ans = strconv.Itoa(zs) + "." + ans[1:]
		}else{
			ans = strconv.Itoa(zs)
		}
	}

	if flag == 1 {
		ans = "-" + ans
	}
	return ans
}

func main() {
	fmt.Println(fractionToDecimal(1, 2))
}

/*
	总结
	1. 该题目考察代码逻辑，使用了寻找相同余数法(使用哈希查找)
	2. 虽然通过了，但是代码结构还是可以优化
	3. 计算的时候先把分数变为假分数，提出整数部分
*/
