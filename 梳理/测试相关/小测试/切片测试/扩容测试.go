package main

import "fmt"

func foo(length,capity int,elementCount int){
	slice := make([]int,length,capity)
	fmt.Printf("未扩容:   len: %d    cap: %d  -> ",len(slice),cap(slice))
	for i:=0;i<elementCount;i++{
		slice = append(slice,0)
	}
	//apSlice := make([]int,elementCount)
	//slice = append(slice,apSlice...)	 // 添加elementCount个元素


	fmt.Printf("扩容后:   len: %d    cap: %d \n",len(slice),cap(slice))
}


func main() {
	for i:=uint8(0);i<12;i++{
		foo(1<<i,1<<i,1024)
	}
}

/*
elementCount := 1
未扩容:   len: 1    cap: 1  -> 扩容后:   len: 2    cap: 2
未扩容:   len: 2    cap: 2  -> 扩容后:   len: 3    cap: 4
未扩容:   len: 4    cap: 4  -> 扩容后:   len: 5    cap: 8
未扩容:   len: 8    cap: 8  -> 扩容后:   len: 9    cap: 16
未扩容:   len: 16    cap: 16  -> 扩容后:   len: 17    cap: 32
未扩容:   len: 32    cap: 32  -> 扩容后:   len: 33    cap: 64
未扩容:   len: 64    cap: 64  -> 扩容后:   len: 65    cap: 128
未扩容:   len: 128    cap: 128  -> 扩容后:   len: 129    cap: 256
未扩容:   len: 256    cap: 256  -> 扩容后:   len: 257    cap: 512
未扩容:   len: 512    cap: 512  -> 扩容后:   len: 513    cap: 1024
未扩容:   len: 1024    cap: 1024  -> 扩容后:   len: 1025    cap: 1280
未扩容:   len: 2048    cap: 2048  -> 扩容后:   len: 2049    cap: 2560



一次性加1024个
未扩容:   len: 1    cap: 1  -> 扩容后:   len: 1025    cap: 1184
未扩容:   len: 2    cap: 2  -> 扩容后:   len: 1026    cap: 1184
未扩容:   len: 4    cap: 4  -> 扩容后:   len: 1028    cap: 1184
未扩容:   len: 8    cap: 8  -> 扩容后:   len: 1032    cap: 1184
未扩容:   len: 16    cap: 16  -> 扩容后:   len: 1040    cap: 1184
未扩容:   len: 32    cap: 32  -> 扩容后:   len: 1056    cap: 1184
未扩容:   len: 64    cap: 64  -> 扩容后:   len: 1088    cap: 1184
未扩容:   len: 128    cap: 128  -> 扩容后:   len: 1152    cap: 1184
未扩容:   len: 256    cap: 256  -> 扩容后:   len: 1280    cap: 1280
未扩容:   len: 512    cap: 512  -> 扩容后:   len: 1536    cap: 1536
未扩容:   len: 1024    cap: 1024  -> 扩容后:   len: 2048    cap: 2560
未扩容:   len: 2048    cap: 2048  -> 扩容后:   len: 3072    cap: 3408


分开加 1024 个
未扩容:   len: 1    cap: 1  -> 扩容后:   len: 1025    cap: 1280
未扩容:   len: 2    cap: 2  -> 扩容后:   len: 1026    cap: 1280
未扩容:   len: 4    cap: 4  -> 扩容后:   len: 1028    cap: 1280
未扩容:   len: 8    cap: 8  -> 扩容后:   len: 1032    cap: 1280
未扩容:   len: 16    cap: 16  -> 扩容后:   len: 1040    cap: 1280
未扩容:   len: 32    cap: 32  -> 扩容后:   len: 1056    cap: 1280
未扩容:   len: 64    cap: 64  -> 扩容后:   len: 1088    cap: 1280
未扩容:   len: 128    cap: 128  -> 扩容后:   len: 1152    cap: 1280
未扩容:   len: 256    cap: 256  -> 扩容后:   len: 1280    cap: 1280
未扩容:   len: 512    cap: 512  -> 扩容后:   len: 1536    cap: 1696
未扩容:   len: 1024    cap: 1024  -> 扩容后:   len: 2048    cap: 2304
未扩容:   len: 2048    cap: 2048  -> 扩容后:   len: 3072    cap: 3408

*/