package main

func removeDuplicates(nums []int) int {
	writer,reader:=0,0
	preNum,preCnt:=-1000000000000000,0
	for reader=0;reader<len(nums);reader++{
		nowNum:=nums[reader]
		if nowNum==preNum{
			// 假如是3次，则这里改3就可以了。
			if preCnt<2{
				nums[writer] = nowNum
				writer++
			}
			preCnt++
		}else{
			preNum = nowNum
			preCnt = 1
			nums[writer] = nowNum
			writer++
		}
	}
	return writer
}

/*
	总结
	1. 这题我是使用读写指针法AC的，方法就是根据读指针指向的数，再与先前的信息进行比对，如果满足要求则writer写
	   否则写指针不动。
	2. 读指针始终前移
*/
