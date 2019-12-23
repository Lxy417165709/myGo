package main

import (
	"fmt"
	"sort"
)
type arr [][]int
func (a arr)Len() int{
	return len(a)
}
func (a arr)Swap(i,j int){
	a[i],a[j] = a[j],a[i]
}
func (a arr)Less(i,j int)bool{
	return a[i][0]<a[j][0]
}
func checkStraightLine(coordinates [][]int) bool {
	sort.Sort(arr(coordinates))
	a :=coordinates[0]
	b :=coordinates[1]

	if a[0]-b[0]==0 {
		x:=a[0]
		for i:=1;i< len(coordinates);i++  {
			if coordinates[i][0]!=x {
				return false
			}
		}
		return true
	}else{
		k := float64(a[1]-b[1])/float64(a[0]-b[0])

		for i:=1;i< len(coordinates);i++  {
			if a[0]-coordinates[i][0]==0 {
				return false
			}
			k2 := float64(a[1]-coordinates[i][1])/float64(a[0]-coordinates[i][0])
			if k!=k2 {
				return false
			}
		}
		return true
	}



}

func main() {
	fmt.Println(checkStraightLine([][]int{
		{-1,0},{0,0},{4,0},{5,0},{6,0},{7,1},
	}))
}
