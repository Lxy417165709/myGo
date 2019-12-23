package main

func sortedSquares(A []int) []int {
	l, r := 0, len(A)-1
	ans := make([]int, 0, len(A))
	for l <= r {
		result := compare(A[l], A[r])
		if result == 0 {
			ans = append(ans, A[l]*A[l])
			l++
		} else {
			if result == 1 {
				ans = append(ans, A[l]*A[l])
				l++
			} else {
				ans = append(ans, A[r]*A[r])
				r--
			}
		}
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return ans
}

func calculateWeight(num1 int) int {
	return num1 * num1
}

func compare(num1, num2 int) int {
	if calculateWeight(num1) == calculateWeight(num2) {
		return 0
	} else {
		if calculateWeight(num1) > calculateWeight(num2) {
			return 1
		} else {
			return -1
		}
	}
}

/*
	总结
	1. 双指针的思路是: 定义l,r := 0,len(A)-1  通过比较A[l],A[r] 指向的值的平方来进行指针的前移。
	2. 双指针感觉也是一个模板，
		可以把计算l,r指向的数的权值抽象为一个函数，即上面的 calculateWeight函数。
		再把比较函数也抽象出来，即 上面的compare函数。
	3. 双指针感觉可以搞成一个接口喔

*/
