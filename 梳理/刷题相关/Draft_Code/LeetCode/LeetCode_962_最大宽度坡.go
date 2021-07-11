package main
import "fmt"

func maxWidthRamp(A []int) int {
	if len(A)==0 {
		return 0
	}
	stack:= make([]int,0)
	stack = append(stack,0)

	for i:=1;i<len(A);i++ {
		num :=get(&stack)
		if A[i]<=A[num] {
			stack = append(stack,i)
		}
	}
	fmt.Println(stack)
	ans:=0
	for i:=len(A)-1;i>=0;i-- {
		for judge(stack) && A[i]>=A[get(&stack)]{
			ans = max(ans,i-get(&stack))
			stack = stack[:len(stack)-1]
		}
	}
	return ans

}
func max(a, b int) int{
	if a>b{
		return a
	}
	return b
}
func judge(stack []int) bool{
	return len(stack)!=0
}
func get(stack *[]int) int{
	return (*stack)[len(*stack)-1]
}
func main() {
	
}
/*
	总结
	1. 这是第一道手撕题 0.0.
	2. get那改用了指针传递，但是内存还是节省不了多少 0.0
*/