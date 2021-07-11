package main

import "fmt"

type MyQueue struct {
	slice []int
}

func (q *MyQueue) Push(a int) {
	q.slice = append(q.slice, a)
}

func (q *MyQueue) Pop() int {
	// 注意这没写判长度操作，因为题目说合法了
	ans := q.slice[0]
	q.slice = q.slice[1:]
	return ans
}
func (q *MyQueue) Size() int {
	return len(q.slice)
}

type MyStack struct {
	queue [2]MyQueue
	mark  int // 0,1取值
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue[this.mark].Push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {

	// 注意这没写判长度操作，因为题目说合法了

	for ; this.queue[this.mark].Size() != 1; {
		this.queue[this.mark^1].Push(this.queue[this.mark].Pop())
	}
	ans := this.queue[this.mark].Pop()
	//this.queue[this.mark^1].Push(ans)	与Top就这一句不一样
	this.mark ^= 1
	return ans

}

/** Get the top element. */
// 这还可以优化，我们可以先把top保存起来 0.0.
func (this *MyStack) Top() int {
	for ; this.queue[this.mark].Size() != 1; {
		this.queue[this.mark^1].Push(this.queue[this.mark].Pop())
	}
	ans := this.queue[this.mark].Pop()
	this.queue[this.mark^1].Push(ans)
	this.mark ^= 1
	return ans
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.queue[this.mark].Size() == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Pop())
	obj.Push(100)

	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())

}

/*
	总结
	1. 该题我使用了2个队列， Push O(1),Pop O(n)，官方有题解 0.0.
	2. 官方还有用一个队列实现的 0.0.. 思想是 Push时把队列翻转 0.0.
	3. 我的做法优化上面有写，比如Top那里

*/
