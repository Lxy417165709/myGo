package main
func isValidSudoku(board [][]byte) bool {
	flag:=true
	for i:=0;i<9;i++{
		flag = flag && judge(getSlice19(board,i)) && judge(getSlice91(board,i))
	}
	for i:=0;i<3;i++{
		for t:=0;t<3;t++{
			flag = flag && judge(getSlice33(board,i*3,t*3))
		}
	}
	return flag
}
func getSlice19(board [][]byte,line int) []int{

	slice := []int{}
	for i:=0;i<9;i++{
		if board[line][i]!='.'{
			slice = append(slice,int(board[line][i]-'0'))
		}
	}
	return slice
}


func getSlice91(board [][]byte,row int) []int{

	slice := []int{}
	for i:=0;i<9;i++{
		if board[i][row]!='.'{
			slice = append(slice,int(board[i][row]-'0'))
		}
	}
	return slice
}
func getSlice33(board [][]byte,x int ,y int) []int{
	slice := []int{}
	for i:=x;i<x+3;i++{
		for t:=y;t<y+3;t++{
			if board[i][t]!='.'{
				slice = append(slice,int(board[i][t]-'0'))
			}
		}
	}
	return slice
}


func judge(slice []int) bool{
	m:=make(map[int]int)
	for i:=0;i<len(slice);i++{
		m[slice[i]]++
	}
	for k,v := range m{
		if k>=10 || k <= 0 || v>=2 {
			return false
		}
	}
	return true
}

/*
	总结
	1. 这是一个暴力的解法，判断每一行每一列每一块是否满足，全满足则true，否则false
		可优化的地方有 judge函数可以不用map,而是使用一个数来做位运算
	2. 官方有更优美的解法，通过3个9*9矩阵，表示第x行/列/块 数字y已经出现了来进行判断
*/
