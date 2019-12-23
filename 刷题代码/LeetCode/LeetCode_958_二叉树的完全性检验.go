package main



// func isCompleteTree(root *TreeNode) bool {
//     queue := []*TreeNode{root}
//     ans := []*TreeNode{}
//     for len(queue)!=0{
//         top := queue[0]
//         queue = queue[1:]
//         ans = append(ans,top)
//         if top==nil{
//             continue
//         }
//         queue = append(queue,top.Left)
//         queue = append(queue,top.Right)
//     }
//     // fmt.Println(ans)
//     flag :=false
//     for i:=0;i<len(ans);i++{
//         if ans[i]==nil{
//             flag = true
//         }else{
//             if flag{
//                 return false
//             }
//         }
//     }
//     return true
// }

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	flag := false
	for len(queue)!=0{
		top := queue[0]
		queue = queue[1:]
		if top==nil{
			flag = true
			continue
		}
		if flag{
			return false
		}
		queue = append(queue,top.Left)
		queue = append(queue,top.Right)
	}

	return true
}

/*
	总结
	1. 刚开始第一种思路是，给每个节点标号，之后通过标号的值，看是否连续，即只要1-n (n为树节点总数)
		这样就可以判断是否是完全二叉树了
	2. 第二种是用层序遍历的方式，或许层序遍历总序列，看nil是否都出现在最后，即不出现一个nil,使得它后面还有非nil的节点。
		这样也可以判断是否是完全二叉树
	3. 以上贴了2份代码，第二份把ans切片去除了，减少了空间。
*/