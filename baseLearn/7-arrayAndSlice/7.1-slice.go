package main

import "fmt"


func show(a []int){
	fmt.Printf("slice=%v,length=%d,cap=%d\n",a,len(a),cap(a))
}

func main(){
	// 1. 切片定义
	var slice0 = make([]int,10,10)
	slice0[5] = 88
	show(slice0)

	// 2. 切片初始化
	// 2.0 切片初始化方式1
	var slice1 = []int{1,2,3,4}
	show(slice1)

	// 2.1 切片初始化方式2
	var array = [...]int{1,2,3,4,5,6,7,8,9}
	slice2 := array[2:]	//切片截取,和python一样
	show(slice2)


	// 3.切片操作(copy,append,删除)

	// 3.0 copy
	var slice3 []int
	show(slice3)
	copy(slice3,slice2)
	show(slice3)
	// slice=[]
	// slice=[]

	slice3 = make([]int,5,10)
	show(slice3)
	copy(slice3,slice2)
	show(slice3)
	//  slice=[0 0 0 0 0],length=5,cap=5
	//  slice=[3 4 5 6 7],length=5,cap=5

	// 3.1 append
	slice3 = append(slice3,1,2,3,4,5,6,7)
	show(slice3)
	//  slice=[3 4 5 6 7 1 2 3 4 5 6 7],length=12,cap=20

	sli4 := make([]int,2,2)
	fmt.Printf("%p\n",sli4)
	sli4 = append(sli4, 1,2,3,4,5,6,7,8,9)
	fmt.Printf("%p\n",sli4)
	show(sli4)
	sli4 = append(sli4, slice3...)	// 切片追加切片
	show(sli4)
	// 0xc042054260
	// 0xc042046060
	// slice=[0 0 1 2 3 4 5 6 7 8 9],length=11,cap=12
	// slice=[0 0 1 2 3 4 5 6 7 8 9 3 4 5 6 7 1 2 3 4 5 6 7],length=23,cap=24



	// 3.2 删除(go语言没有切片元素删除函数)
	// 3.20 普通删除
	index := 2	// 需要删除元素的索引
	slix := []int{1,2,3,4,5,6,7}
	newSlix := append(slix[0:index],slix[index+1:]...)
	show(slix)
	show(newSlix)
	fmt.Printf("%p %p\n",slix,newSlix)
	// slice=[1 2 4 5 6 7 7],length=7,cap=7	(slix也改变了~因为slice本质是指针)
	// slice=[1 2 4 5 6 7],length=6,cap=7
	// 0xc04207a0c0 0xc04207a0c0

	// 3.21 使用copy实现删除
	slix1 := []int{1,2,3,4,5,6,7}
	newSlix1 := make([]int,index)
	copy(newSlix1,slix)
	newSlix1 = append(newSlix1,slix1[index+1:]...)
	show(slix1)
	show(newSlix1)
	fmt.Printf("%p %p\n",slix,newSlix)
	// slice=[1 2 3 4 5 6 7],length=7,cap=7	(slix没改变了)
	// slice=[1 2 4 5 6 7],length=6,cap=6
	// 0xc04207a0c0 0xc04207a0c0


}

/*
	总结
	1. 切片的复制只能复制length个,大于后就不复制了
	2. 切片的追加如果超过容量,容量会自动*2 (这是不准确的说法)
    3. 切片元素的写入读出操作和数组一样
	4. 长度表示当前切片中的数据个数,容量表示当前切片可占用的空间位子个数
	5. 当切片所占用的空间位子个数改变的时候,切片的指针也会改变(会找到另外一个连续的内存区域,把其首指针赋值给切片)
	6. 切片本质是指针
	7. copy按角标一一复制元素,不会使长度短的一方扩容
*/
