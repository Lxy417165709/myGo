package main

import (
	"fmt"
	"strconv"
	"strings"
)

func queryString(S string, N int) bool {
	for i:=N;i>=1;i--{
		str := toBinary(i)
		fmt.Println(str)
		if strings.Index(S,str)==-1{
			return false
		}
	}
	return true
}


func toBinary(a int) string{
	str := ""
	for a!=0{
		str = strconv.Itoa(a&1) + str
		a>>=1
	}
	if str==""{
		str="0"
	}
	return str
}
/*

	总结
	1. 这题直接暴力就可以了，因为1000位的string,能够组成满足题目的数字很少(据说是30000多个)，
		所以我们直接从N->1遍历，找到不满足就false，全满足就true
	2. 这题目挺搞的 ... 不过的确需要一些思维
	3. 第一次做的时候，我在纸上列出一些规律，但是这些规律还是无法AC这个题 ...
*/
