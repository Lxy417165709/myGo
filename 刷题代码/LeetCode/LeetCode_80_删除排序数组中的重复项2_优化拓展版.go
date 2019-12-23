package main

func removeDuplicates(nums []int) int {
	k:= 2 // 数字最多出现的次数
	writer,reader := k,k
	for ;reader<len(nums);reader++{
		if nums[writer-k] != nums[reader]{
			nums[writer] = nums[reader]
			writer++
		}
	}
	return writer
}

/*
	总结
	1. 这次做法和上次差不多，我是看了官方解答后才采用的
	2. 思路是，把当前读的数与当前写的指针-2的数进行比较，如果是相同，说明写入后会变成3个相同的数
		连在一起，所以此时我们不写入。如果不相同，说明我们写入后不会超过2个，所以执行写入，写指针再进行前移。
*/
