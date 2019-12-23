package main
func compress(chars []byte) int {
	writer,reader:=0,0
	for reader!=len(chars){
		times:=0
		begin:=reader
		char:=chars[reader]
		for begin<len(chars) && chars[begin]==char{
			begin++
			times++
		}

		chars[writer] = char
		writer++
		if times!=1{
			num:=1000
			flag:=0
			for num!=0{
				tmp := times/num
				if tmp!=0 || flag==1{
					chars[writer] = uint8(tmp)+'0'
					writer++
					times = times%num
					flag=1
				}
				num/=10
			}
		}
		reader = begin
	}
	return writer
}
func main() {

}

/*
	总结
	1. 这题我是使用双指针法AC的
	2. 第一次写的时候也想到了双指针，然后实际操作的话是暴力移动0.0
	   然后看了官方题解后，就借助双指针AC了
	3. 其实可以当做一个模板喔这个 0.0
*/