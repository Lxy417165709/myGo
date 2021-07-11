package _1nod

import (
	"fmt"
	"math"
)

var m = [105][30050]int{}
var l = [105]int{}
var ele = [30050]int{}




func getPosition (arr []int,num int) int{

	L := 0
	R := len(arr)-1

	for ;L<=R;{
		M := (L+R)/2
		if arr[M] > num{
			R = M - 1
		}else{
			L = M + 1
		}
	}

	return L
}
func get(num int,x int) int{

	//fmt.Println(x,num)
	ans:= 50000

	for i:=num+1;i<=100 ;i++  {
		y := getPosition(m[i][:l[i]],x)

		if y != l[i]{

			ans = int(math.Min(float64(ans),float64(m[i][y])))
		}
	}
	if ans!=50000{
		return ans - x
	}else {
		return 0
	}

}
func main() {

	n:=0
	fmt.Scan(&n)
	for i:=1;i<=n;i++{
		a:=0
		fmt.Scan(&a)
		ele[i] = a
		m[a][l[a]] = i
		l[a]++
	}

	for i:=1;i<=n;i++{
		if  i==n{
			fmt.Println(get(ele[i],i))
		}else {
			fmt.Print(get(ele[i],i)," ")
		}
		
	}
}
