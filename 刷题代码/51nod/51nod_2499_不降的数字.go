package _1nod

import "fmt"

var arr = [15]int{}

func toArr(n int) int{
	length:=0
	for ;n!=0;{
		arr[length] = n%10
		length++
		n = n / 10
	}
	return length
}

func toInt(sli []int) int{
	ans:=0
	for i:=0;i<len(sli) ;i++  {
		ans = ans * 10 + sli[len(sli) - 1 - i]
	}
	return ans
}
func main() {

	n:=0
	fmt.Scan(&n)
	length := toArr(n)
	//fmt.Println(arr,length)
	flag:=0
	for ;flag==0;{
		tag := 1
		for i:=length-1;i>= 1 ;i--  {
			if arr[i] > arr[i-1]{
				arr[i]--
				for t:=0;t<=i-1 ;t++  {
					arr[t] = 9
				}
				if i == length-1 && arr[i] == 0{
					length --
				}
				tag = 0
				break

			}
		}
		if tag == 1{
			flag = 1
		}
	}
	fmt.Println(toInt(arr[0:length]))

}

/*
	总结
	1. 思维题，其实没那么复杂的，不会的动手模拟模拟就知道了~

*/