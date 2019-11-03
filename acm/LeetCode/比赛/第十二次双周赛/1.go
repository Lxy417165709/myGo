package main

import "fmt"

func transformArray(arr []int) []int {
	for d:=0;d<100 ;d++  {
		tmp:=make([]int,1)
		for i:=1;i<len(arr)-1;i++{


			if arr[i]>arr[i-1] && arr[i]>arr[i+1]{
				tmp = append(tmp, -1)
			}else{
				if arr[i]<arr[i-1] && arr[i]<arr[i+1]{
					tmp = append(tmp, 1)
				}else{
					tmp = append(tmp, 0)
				}
			}


		}
		for i:=1;i<len(arr)-1;i++{
			arr[i]+=tmp[i]
		}
	}


	return arr
}
func main() {
	fmt.Println(transformArray([]int{
		1,6,3,4,3,5,
	}))
}
