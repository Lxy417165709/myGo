package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 目的: 输出变量a
func problemOne() {

	a = 10 // mark
	fmt.Println(a)

	// undefined: a
	// 出现问题: 提示a未定义
	// 解决方法: a = 10    ---->    a := 10
}

// 目的: 先定义一个空指针，再让这个指针指向传入的链表表头head
func problemTwo(head *ListNode) {

	nullPointer := nil // mark
	nullPointer = head
	fmt.Println(nullPointer)

	// use of untyped nil
	// 出现问题: 编译报错 --- 使用无类型的nil
	// 解决方法:
	//		nullPointer := nil    ---->    var nullPointer *ListNode	(Go语言变量默认值为该类型的"零"值)
}

// 目的: 定义一个函数计算a + b，并返回结果
func problemThree(a, b int) { // mark

	return a + b
	// too many arguments to return
	// 出现问题: 编译报错 --- 返回了太多参数
	// 解决方法: func problemThree(a, b int)    ---->    func problemThree(a, b int) int
}

func main() {
	// problemOne()
	// problemTwo(&ListNode{5, nil})
	// problemThree(1,5)
}
