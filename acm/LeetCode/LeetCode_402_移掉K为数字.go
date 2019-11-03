package main

import "fmt"

func removeKdigits(num string, k int) string {
	slice:=make([]byte,0)
	slice = []byte(num)
	for i:=1;i<=k ;i++  {
		index:=0
		for x:=index+1;x< len(slice);x++  {
			if slice[x]<slice[index] {
				break
			}else{
				index++
			}
		}
		slice = append(slice[0:index],slice[index+1:]...)

	}
	index:=0
	for i:=0;i< len(slice);i++  {
		if slice[i]=='0' {
			index++
		}else{
			break
		}
	}
	slice = slice[index:]
	if len(slice)==0 {
		slice = append(slice, '0')
	}

	return string(slice)

}

func main() {
	fmt.Println(removeKdigits("1432219",3))
}
/*
	总结
	1. 这题目我是通过自己模拟，发现了一些规律 只要每次都移除字符串每个升序子串中最大的一位数字就可以了~
*/