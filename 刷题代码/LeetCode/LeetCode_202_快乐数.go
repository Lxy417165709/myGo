package main

func isHappy(n int) bool {
	slow, fast := n, n
	for true {
		slow = getNum(slow)
		fast = getNum(getNum(fast))
		if slow == fast {
			return slow == 1
		}
	}
	return n == 1	// 按照逻辑来说，这句是不会执行的
}

func getNum(n int) int {
	ans := 0
	for n != 0 {
		ans += (n % 10) * (n % 10)
		n /= 10
	}
	return ans
}

/*
	总结
	1. 第一次的时候直接循环100次，判断结果是否为1，如果是，则为快乐数，否则不是
	2. 看了官方题解后，发现可以用快慢指针，然后根据这个思想就写出AC代码了
*/
