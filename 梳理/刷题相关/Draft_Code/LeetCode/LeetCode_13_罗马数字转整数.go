package main

import "fmt"

func romanToInt(s string) int {
	m := make(map[uint8]int)
	m['I']=1
	m['V']=5
	m['X']=10
	m['L']=50
	m['C']=100
	m['D']=500
	m['M']=1000
	sum:=0

	for index:=0;index< len(s);index++ {
		ch:=s[index]
		if ch=='I' || ch=='X' || ch=='C' {
			flag:=0
			if ch=='I' && index+1< len(s) && (s[index+1]=='V' || s[index+1]=='X') {
				flag=1
			}
			if ch=='X' && index+1< len(s) && (s[index+1]=='L' || s[index+1]=='C') {
				flag=1
			}
			if ch=='C' && index+1< len(s) && (s[index+1]=='D' || s[index+1]=='M') {
				flag=1
			}

			if flag==1 {
				sum=sum + m[s[index+1]]-m[s[index]]
				index++
				continue
			}
		}
		sum = sum + m[s[index]]
	}
	return sum
}

func main() {
	fmt.Println( romanToInt("III"))
	fmt.Println( romanToInt("IV"))
	fmt.Println( romanToInt("IX"))
	fmt.Println( romanToInt("LVIII"))
	fmt.Println( romanToInt("MCMXCIV"))
}

/*
	总结
	1. go语言没有char类型，但可以用unit8代替
*/