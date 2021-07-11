package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 1. 创建双向链表
	mylist := list.New()
	fmt.Println(&mylist)       // 这个是mylist的地址
	fmt.Printf("%p\n", mylist) // 这个应该是mylist指向的地址
	// 0xc042074018
	// 0xc04206a1e0

	// 2. 查看双向链表信息及长度
	fmt.Println(mylist, mylist.Len())
	// &{{0xc04206a1e0 0xc04206a1e0 <nil> <nil>} 0} 0

	// 3. CRUD
	// 3.0 向list增加元素
	mylist.PushBack("a")  // 添加到最后
	mylist.PushFront("b") // 添加到最前
	fmt.Println(mylist, mylist.Len())
	// &{{0xc04206a270 0xc04206a240 <nil> <nil>} 2} 2

	// 3.1 向双向链表插入元素
	mylist.InsertAfter("c", mylist.Front()) // 插入到第一个元素后
	mylist.InsertBefore("d", mylist.Back()) // 插入到最后一个元素前
	fmt.Println(mylist, mylist.Len())
	// &{{0xc04206a270 0xc04206a240 <nil> <nil>} 4} 4

	// 3.2 查看双向链表元素值
	// 3.20 查看首末元素值
	fmt.Println(mylist.Front().Value) // 第一个元素的值
	fmt.Println(mylist.Back().Value)  // 最后一个元素的值

	// 3.21 查看指定索引元素值
	index := 1 // 查看索引为index的元素的值
	if index >= 0 && index < mylist.Len() {
		x := mylist.Front()
		for i := 0; i < index; i++ {
			x = x.Next()
		}
		fmt.Println(x.Value)
	}
	// b
	// a
	// c

	// 3.22 遍历
	for x := mylist.Front(); x != nil; x = x.Next() {
		fmt.Print(x.Value, " ")
	}
	fmt.Println()
	// b c d a

	// 3.3 删除双向链表的元素
	mylist.Remove(mylist.Front()) // 删除首元素
	for x := mylist.Front(); x != nil; x = x.Next() {
		fmt.Print(x.Value, " ")
	}
	fmt.Println()
	fmt.Println(mylist, mylist.Len())
	// c d a
	// &{{0xc04206a2d0 0xc04206a240 <nil> <nil>} 3} 3

	// 3.4 修改双向链表的元素值
	// 3.40 简单修改
	mylist.Front().Value = "lixueyue" // 修改首元素的值
	for x := mylist.Front(); x != nil; x = x.Next() {
		fmt.Print(x.Value, " ")
	}
	fmt.Println()
	// lixueyue d a

	// 3.41 移动元素顺序
	mylist.MoveToFront(mylist.Back()) // 把尾元素移前
	mylist.MoveToBack(mylist.Front()) // 把首元素移动到最后
	for x := mylist.Front(); x != nil; x = x.Next() {
		fmt.Print(x.Value, " ")
	}
	fmt.Println()
	// lixueyue d a

	mylist.MoveBefore(mylist.Front(), mylist.Back()) // 把首元素移动到倒数第二位
	for x := mylist.Front(); x != nil; x = x.Next() {
		fmt.Print(x.Value, " ")
	}
	fmt.Println()
	// d lixueyue a

	mylist.MoveAfter(mylist.Back(), mylist.Front()) // 把尾元素移动到第二位
	for x := mylist.Front(); x != nil; x = x.Next() {
		fmt.Print(x.Value, " ")
	}
	fmt.Println()
	// d a lixueyue

	fmt.Println(mylist, mylist.Len())
	// &{{0xc04206a300 0xc04206a2d0 <nil> <nil>} 3} 3
}

/*
	总结
	1. go的双向链表的操作都是要基于该元素的指针
	2. mylist.Back()不是nil,而是最后一个元素的指针
*/
