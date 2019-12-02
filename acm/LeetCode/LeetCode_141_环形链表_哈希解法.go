package main


func hasCycle(head *ListNode) bool {
	mp := make(map[*ListNode]int)
	for head!=nil{
		if _,ok:=mp[head];ok{
			return true
		}
		mp[head]=1
		head=head.Next
	}
	return false
}

/*
	总结
	1. 哈希解法的思路就是每次都标记访问过的节点，如果再次被访问，则说明有环
	2. 如果没环，那么head.Next.Next....必定有个是nil
*/
