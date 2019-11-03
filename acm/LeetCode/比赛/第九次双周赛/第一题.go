package 第九次双周赛

import "fmt"

func smallestCommonElement(mat [][]int) int {

	if len(mat)==0 {
		return -1
	}
	m:=make(map[int]int)
	for i:=0;i< len(mat);i++  {
		mx:=make(map[int]int)
		for t:=0;t< len(mat[i]);t++  {
			if mx[mat[i][t]]==1 {
				continue
			}
			m[mat[i][t]]++
		}
	}
	for i:=0;i< len(mat[0]);i++  {
		if m[mat[0][i]]==len(mat) {
			return mat[0][i]
		}
	}
	return -1
}


func main() {
	fmt.Println(smallestCommonElement([][]int{
		{0,2,3,4,5},{7,8,9,10,11},
	}))
}
