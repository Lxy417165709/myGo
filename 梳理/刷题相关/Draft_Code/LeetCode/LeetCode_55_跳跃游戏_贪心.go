package main

func canJump(nums []int) bool {
	if len(nums)==0{
		return false
	}
	end := 0
	for i:=len(nums)-1;i>=0;i--{
		if i+nums[i]>=end{
			end = i
		}
	}
	return nums[0]>=end
}

/*
	总结
	1. 这题，我用end表示最终位置，通过贪心算法不断更新它。最后，到第一个位置时，如果能到达end，则为true,否则为false。
*/
