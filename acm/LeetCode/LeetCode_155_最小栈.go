package main

import "fmt"

type MinStack struct {
	slice []int
	arr []int
}


/** initialize your data structure here. */
func Constructor() MinStack {

	return MinStack{[]int{},[]int{}}
}


func (this *MinStack) Push(x int)  {
	this.slice = append(this.slice,x)


	this.arr = append(this.arr,x)
	for i:= len(this.arr)-1;i>=1 ;i--  {
		if this.arr[i]>this.arr[i-1] {
			this.arr[i-1],this.arr[i] = this.arr[i],this.arr[i-1]
		}else{
			break
		}
	}


}


func (this *MinStack) Pop()  {

	if len(this.slice)==0 {
		return
	}

	x :=this.Top()
	this.slice = this.slice[:len(this.slice)-1]

	for i:=len(this.arr)-1;i>=0 ;i--  {
		if this.arr[i] == x {
			this.arr = append(this.arr[0:i],this.arr[i+1:]...)
			break
		}
	}
	//fmt.Println(this.arr)


}


func (this *MinStack) Top() int {
	return this.slice[len(this.slice)-1]
}


func (this *MinStack) GetMin() int {
	return this.arr[len(this.arr)-1]
}

func main() {
	obj := Constructor()
	obj.Push(3)
	obj.Push(80)
	obj.Push(-35)
	obj.Push(9)
	fmt.Println(obj.Top(),obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top(),obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top(),obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top(),obj.GetMin())

}

/*
	总结
	1. 该题目要理解栈的结构与操作


*/