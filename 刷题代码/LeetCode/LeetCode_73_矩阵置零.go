package main
var inf = 1<<50

func setZeroes(matrix [][]int)  {
	if len(matrix)==0{
		return
	}
	m,n:=len(matrix),len(matrix[0])
	for i:=0;i<m;i++{
		for t:=0;t<n;t++{
			if matrix[i][t]==0{
				for u:=0;u<m;u++{
					// matrix[u][t]!=0 很重要，不然的话该位置的0就不能影响其他行列了
					if u!=i && matrix[u][t]!=0{
						matrix[u][t] = inf
					}
				}
				for v:=0;v<n;v++{
					// matrix[i][v]!=0 很重要，不然的话该位置的0就不能影响其他行列了
					if v!=t && matrix[i][v]!=0{
						matrix[i][v] = inf
					}
				}
			}
		}
	}

	for i:=0;i<m;i++{
		for t:=0;t<n;t++{
			if matrix[i][t]==inf{
				matrix[i][t] = 0
			}
		}
	}
}
func main() {

}
/*
	总结
	1. 为了满足题意，我扫描的时候把需要清0的位置设置为inf,之后再扫描一次，
	   把inf变为0就可以了
	2. 官方有更好的方法，我这个方法对应官方的方法2，方法3就是延时标记
	   当有0时，把该行和该列的第一个元素都置为0，扫描完毕后，再遍历一次进行清0处理
*/