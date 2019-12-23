package main
func magicalString(n int) int {
	return get(n)
}


func get(n int) int{
	if n==0{
		return 0
	}

	str:="1"    // 神奇字符串
	ans:=1      // 神奇字符串1的个数
	nowChar:='1'    // 数字
	times:=1   // 数字使用次数
	pointer:=0  // 一个指针，指向a的字符

	for len(str)<n{
		// 表示数字不能用了，于是要切换数字，且指针要指向下一位
		if times==int(str[pointer]-'0'){
			pointer++
			if nowChar=='2'{
				nowChar='1'
			}else{
				nowChar='2'
			}
			times = 0
		}
		if nowChar=='1'{
			ans++
		}
		str = str + string(nowChar)
		times++
	}

	return ans
}

/*
	总结
	1. 这题我是模拟AC的 0.0
	2. 做题之前先在纸上模拟模拟喔~
*/
