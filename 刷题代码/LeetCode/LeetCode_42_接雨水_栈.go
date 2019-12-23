package 双指针

import "fmt"

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	stack := make([]int, 0)
	sum := 0
	for i := 0; i < len(height); i++ {
		for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {

			topHeight := height[stack[len(stack)-1]]  // 这是此时栈顶的高度

			stack = stack[:len(stack)-1]
			// 此时表示已经没有左边界了
			if len(stack) == 0 {
				break
			}
			leftIndex := stack[len(stack)-1]	// 这是左边界的下标(为了计算宽度，所以采用下标形式)
			w := i - leftIndex - 1
			h := min(height[leftIndex], height[i]) - topHeight
			sum += w * h
		}
		stack = append(stack, i)
	}
	return sum
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	总结
	1. 这是采用栈来实现的。(还是有些不熟练)


*/
