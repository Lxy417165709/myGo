package main


func get(a int) int{
	if a == 0{
		return 1
	}

	length := 0
	for a!=0{
		length++
		a/=10
	}
	return length
}

func findNumbers(nums []int) int {
	ans := 0
	for i:=0;i<len(nums);i++{
		if get(nums[i])%2==0{
			ans++
		}
	}
	return ans

}