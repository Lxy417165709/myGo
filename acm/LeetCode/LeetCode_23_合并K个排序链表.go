package main

func mergeKLists(lists []*ListNode) *ListNode {
	dummy:=&ListNode{-1,nil}
	tmp:=dummy

	for true{
		index:=-1
		mn:=0
		for i:=0;i<len(lists);i++{
			if lists[i]==nil{
				continue
			}
			if index == -1 || lists[i].Val<mn{
				mn = lists[i].Val
				index = i
			}
		}
		if index==-1{
			break
		}
		tmp.Next = &ListNode{mn,nil}
		tmp = tmp.Next
		lists[index] = lists[index].Next
	}
	return dummy.Next
}
func main() {

}

/*
	总结
	1. 这是第二次手撕，采取了另外的方法，我每次都扫描每一个链表表头，选择最小的一个，之后插入答案链表，并
		把最小一个元素的链表后移一个位置，直到所有链表都为nil。
	2. 第一次做的时候是类似于插入排序的，将数据与其索引封装成一个Node，然后有保存一个Node数组。将每个链表的
		数据封装后加入Node数组，每次取头部就可以了，当然，插入的时候要做相应的维护。Node数组长度为链表数组的长度
*/
