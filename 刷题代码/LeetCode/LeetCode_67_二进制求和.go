package main

import (
	"fmt"
	"strconv"
)

func rev(a string) string{
	slice := []byte(a)
	length:= len(slice)
	for i:=0;i< length/2;i++  {
		slice[i],slice[length-1-i] = slice[length-1-i],slice[i]
	}
	return string(slice)
}

func addBinary(a string, b string) string {
	if len(a)< len(b) {
		a,b = b,a
	}
	a,b = rev(a),rev(b)
	slice_a,slice_b := []byte(a),[]byte(b)
	slice_c := make([]int, len(slice_a))
	carry := 0

	for i:=0;i< len(slice_a);i++  {

		num_b:=0
		if i<len(slice_b) {
			num_b = int(slice_b[i]-'0')
		}
		slice_c[i] = int(slice_a[i]-'0') + num_b + carry
		carry = slice_c[i]/2
		slice_c[i] = slice_c[i]%2
	}
	if carry==1 {
		slice_c = append(slice_c, 1)
	}
	ans:=""


	for i:=0;i< len(slice_c);i++  {

		ans =   strconv.Itoa(slice_c[i]) + ans
	}
	return ans
}
func main() {
	fmt.Println(addBinary("1010","1011"))
}


/*
	总结
	1. 该题考查普通加法规则~注意进位
*/