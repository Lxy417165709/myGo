package main
func minMutation(start string, end string, bank []string) int {
	mark = make(map[int]int)

	cgStr:="ACGT"

	flag:=0
	for i:=0;i<len(bank);i++{
		if end==bank[i]{
			flag = 1
			break
		}
	}
	if flag==0{
		return -1
	}
	queue :=make([]string,0)
	queue = append(queue,start);
	level := 0
	for len(queue)!=0{
		size:=len(queue)
		for sz:=0;sz<size;sz++{
			top:=queue[0]
			queue=queue[1:]
			if top==end{
				return level
			}

			for i:=0;i<len(top);i++{
				for t:=0;t<len(cgStr);t++{
					if top[i]==cgStr[t]{
						continue
					}

					tmp:=[]byte(top)

					tmp[i] = cgStr[t]
					newStr := string(tmp)
					for j:=0;j<len(bank);j++{
						if mark[j]==1{
							continue
						}
						if bank[j]==newStr{
							queue = append(queue,newStr)
							mark[j] = 1
						}
					}
				}
			}
		}
		level++
	}
	return -1
}

var mark map[int]int



func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}

func main() {

}

/*
	总结
	1. 这个我使用了广度搜索写 0.0 AC了
*/
