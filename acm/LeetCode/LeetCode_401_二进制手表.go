package main

import "strconv"

const number = 10
var seq [number]int
func readBinaryWatch(num int) []string {
	seq = [number]int{}
	ansSet = make([]string,0)
	solve(num,0)
	return ansSet
}

var ansSet []string

func solve(n int,begin int){

	if begin>number{
		return
	}
	if n==0{
		hours,minutes:=getHours(),getMinutes()
		if hours!="" && minutes!=""{
			time := hours+":"+minutes
			ansSet = append(ansSet,time)
		}
		return
	}



	for i:=begin;i<number;i++{

		seq[i] = 1
		solve(n-1,i+1)
		seq[i] = 0
	}
}

func getHours() string{
	num:=0
	for i:=0;i<4;i++{
		num = num*2 + seq[i]
	}
	if num>=12{
		return ""
	}
	ans := strconv.Itoa(num)
	return ans

}
func getMinutes() string{
	num:=0
	for i:=4;i<number;i++{
		num = num*2 + seq[i]
	}
	if num>=60{
		return ""
	}
	ans := strconv.Itoa(num)
	if len(ans)==1{
		ans = "0" + ans
	}
	return ans

}



/*
	总结
	1. 这题目我是使用回溯ac的，中间还因为自己大意，没把begin改为i一直找了好久错
	2. 其实可以暴力ac，就是判断每一个时间的1的位数和是否满足n，满足则加入，否则不加入
*/
