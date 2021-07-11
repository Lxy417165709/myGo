package main

import (
	"fmt"
	"sort"
)

// 测试增长后的容量
func TestExpansion(oldCap int) int {
	s := make([]int, oldCap, oldCap)
	s = append(s, 1)
	return cap(s)
}

// 形成测试集
func formSlice(l, diff int) []int {
	testSet := make([]int, 0)
	for i := 0; i < l; i++ {
		c := (1 << uint(i)) - diff
		if c >= 0 {
			testSet = append(testSet, c)
		}
	}
	return testSet
}

// 展示结果
func showResult(testSet []int){
	for _, v := range testSet {
		old, ne := v, TestExpansion(v)
		fmt.Printf("old: %d  new: %d    rate: %f\n", old, ne, float64(ne)/float64(old))
	}
	fmt.Println()
}


func main() {
	testSet := formSlice(28,0)
	testSet1 := formSlice(28,1)
	testSet2 := formSlice(28,2)
	testSet3 := formSlice(28,10)
	testSet4 := formSlice(28,15)
	testSet5 := formSlice(28,30)

	allTest := append(testSet,testSet1...)
	allTest = append(allTest,testSet2...)
	allTest = append(allTest,testSet3...)
	allTest = append(allTest,testSet4...)
	allTest = append(allTest,testSet5...)
	sort.Ints(allTest)
	showResult(allTest)
	// old: 0  new: 1    rate: +Inf
	//old: 0  new: 1    rate: +Inf
	//old: 1  new: 2    rate: 2.000000
	//old: 1  new: 2    rate: 2.000000
	//old: 1  new: 2    rate: 2.000000
	//old: 2  new: 4    rate: 2.000000
	//old: 2  new: 4    rate: 2.000000
	//old: 2  new: 4    rate: 2.000000
	//old: 3  new: 6    rate: 2.000000
	//old: 4  new: 8    rate: 2.000000
	//old: 6  new: 12    rate: 2.000000
	//old: 6  new: 12    rate: 2.000000
	//old: 7  new: 14    rate: 2.000000
	//old: 8  new: 16    rate: 2.000000
	//old: 14  new: 28    rate: 2.000000
	//old: 15  new: 30    rate: 2.000000
	//old: 16  new: 32    rate: 2.000000
	//old: 17  new: 36    rate: 2.117647
	//old: 22  new: 44    rate: 2.000000
	//old: 30  new: 60    rate: 2.000000
	//old: 31  new: 64    rate: 2.064516
	//old: 32  new: 64    rate: 2.000000
	//old: 34  new: 72    rate: 2.117647
	//old: 49  new: 112    rate: 2.285714
	//old: 54  new: 112    rate: 2.074074
	//old: 62  new: 128    rate: 2.064516
	//old: 63  new: 128    rate: 2.031746
	//old: 64  new: 128    rate: 2.000000
	//old: 98  new: 224    rate: 2.285714
	//old: 113  new: 256    rate: 2.265487
	//old: 118  new: 256    rate: 2.169492
	//old: 126  new: 256    rate: 2.031746
	//old: 127  new: 256    rate: 2.015748
	//old: 128  new: 256    rate: 2.000000
	//old: 226  new: 512    rate: 2.265487
	//old: 241  new: 512    rate: 2.124481
	//old: 246  new: 512    rate: 2.081301
	//old: 254  new: 512    rate: 2.015748
	//old: 255  new: 512    rate: 2.007843
	//old: 256  new: 512    rate: 2.000000
	//old: 482  new: 1024    rate: 2.124481
	//old: 497  new: 1024    rate: 2.060362
	//old: 502  new: 1024    rate: 2.039841
	//old: 510  new: 1024    rate: 2.007843
	//old: 511  new: 1024    rate: 2.003914
	//old: 512  new: 1024    rate: 2.000000
	//old: 994  new: 2048    rate: 2.060362
	//old: 1009  new: 2048    rate: 2.029732
	//old: 1014  new: 2048    rate: 2.019724
	//old: 1022  new: 2048    rate: 2.003914
	//old: 1023  new: 2048    rate: 2.001955
	//old: 1024  new: 1280    rate: 1.250000
	//old: 2018  new: 2560    rate: 1.268583
	//old: 2033  new: 2560    rate: 1.259223
	//old: 2038  new: 2560    rate: 1.256133
	//old: 2046  new: 2560    rate: 1.251222
	//old: 2047  new: 2560    rate: 1.250611
	//old: 2048  new: 2560    rate: 1.250000
	//old: 4066  new: 5120    rate: 1.259223
	//old: 4081  new: 5120    rate: 1.254594
	//old: 4086  new: 5120    rate: 1.253059
	//old: 4094  new: 5120    rate: 1.250611
	//old: 4095  new: 5120    rate: 1.250305
	//old: 4096  new: 5120    rate: 1.250000
	//old: 8162  new: 10240    rate: 1.254594
	//old: 8177  new: 10240    rate: 1.252293
	//old: 8182  new: 10240    rate: 1.251528
	//old: 8190  new: 10240    rate: 1.250305
	//old: 8191  new: 10240    rate: 1.250153
	//old: 8192  new: 10240    rate: 1.250000
	//old: 16354  new: 20480    rate: 1.252293
	//old: 16369  new: 20480    rate: 1.251145
	//old: 16374  new: 20480    rate: 1.250763
	//old: 16382  new: 20480    rate: 1.250153
	//old: 16383  new: 20480    rate: 1.250076
	//old: 16384  new: 20480    rate: 1.250000
	//old: 32738  new: 40960    rate: 1.251145
	//old: 32753  new: 40960    rate: 1.250572
	//old: 32758  new: 40960    rate: 1.250382
	//old: 32766  new: 40960    rate: 1.250076
	//old: 32767  new: 40960    rate: 1.250038
	//old: 32768  new: 40960    rate: 1.250000
	//old: 65506  new: 81920    rate: 1.250572
	//old: 65521  new: 81920    rate: 1.250286
	//old: 65526  new: 81920    rate: 1.250191
	//old: 65534  new: 81920    rate: 1.250038
	//old: 65535  new: 81920    rate: 1.250019
	//old: 65536  new: 81920    rate: 1.250000
	//old: 131042  new: 163840    rate: 1.250286
	//old: 131057  new: 163840    rate: 1.250143
	//old: 131062  new: 163840    rate: 1.250095
	//old: 131070  new: 163840    rate: 1.250019
	//old: 131071  new: 163840    rate: 1.250010
	//old: 131072  new: 163840    rate: 1.250000
	//old: 262114  new: 327680    rate: 1.250143
	//old: 262129  new: 327680    rate: 1.250072
	//old: 262134  new: 327680    rate: 1.250048
	//old: 262142  new: 327680    rate: 1.250010
	//old: 262143  new: 327680    rate: 1.250005
	//old: 262144  new: 327680    rate: 1.250000
	//old: 524258  new: 655360    rate: 1.250072
	//old: 524273  new: 655360    rate: 1.250036
	//old: 524278  new: 655360    rate: 1.250024
	//old: 524286  new: 655360    rate: 1.250005
	//old: 524287  new: 655360    rate: 1.250002
	//old: 524288  new: 655360    rate: 1.250000
	//old: 1048546  new: 1310720    rate: 1.250036
	//old: 1048561  new: 1310720    rate: 1.250018
	//old: 1048566  new: 1310720    rate: 1.250012
	//old: 1048574  new: 1310720    rate: 1.250002
	//old: 1048575  new: 1310720    rate: 1.250001
	//old: 1048576  new: 1310720    rate: 1.250000
	//old: 2097122  new: 2621440    rate: 1.250018
	//old: 2097137  new: 2621440    rate: 1.250009
	//old: 2097142  new: 2621440    rate: 1.250006
	//old: 2097150  new: 2621440    rate: 1.250001
	//old: 2097151  new: 2621440    rate: 1.250001
	//old: 2097152  new: 2621440    rate: 1.250000
	//old: 4194274  new: 5242880    rate: 1.250009
	//old: 4194289  new: 5242880    rate: 1.250004
	//old: 4194294  new: 5242880    rate: 1.250003
	//old: 4194302  new: 5242880    rate: 1.250001
	//old: 4194303  new: 5242880    rate: 1.250000
	//old: 4194304  new: 5242880    rate: 1.250000
	//old: 8388578  new: 10485760    rate: 1.250004
	//old: 8388593  new: 10485760    rate: 1.250002
	//old: 8388598  new: 10485760    rate: 1.250001
	//old: 8388606  new: 10485760    rate: 1.250000
	//old: 8388607  new: 10485760    rate: 1.250000
	//old: 8388608  new: 10485760    rate: 1.250000
	//old: 16777186  new: 20971520    rate: 1.250002
	//old: 16777201  new: 20971520    rate: 1.250001
	//old: 16777206  new: 20971520    rate: 1.250001
	//old: 16777214  new: 20971520    rate: 1.250000
	//old: 16777215  new: 20971520    rate: 1.250000
	//old: 16777216  new: 20971520    rate: 1.250000
	//old: 33554402  new: 41943040    rate: 1.250001
	//old: 33554417  new: 41943040    rate: 1.250001
	//old: 33554422  new: 41943040    rate: 1.250000
	//old: 33554430  new: 41943040    rate: 1.250000
	//old: 33554431  new: 41943040    rate: 1.250000
	//old: 33554432  new: 41943040    rate: 1.250000
	//old: 67108834  new: 83886080    rate: 1.250001
	//old: 67108849  new: 83886080    rate: 1.250000
	//old: 67108854  new: 83886080    rate: 1.250000
	//old: 67108862  new: 83886080    rate: 1.250000
	//old: 67108863  new: 83886080    rate: 1.250000
	//old: 67108864  new: 83886080    rate: 1.250000
	//old: 134217698  new: 167772160    rate: 1.250000
	//old: 134217713  new: 167772160    rate: 1.250000
	//old: 134217718  new: 167772160    rate: 1.250000
	//old: 134217726  new: 167772160    rate: 1.250000
	//old: 134217727  new: 167772160    rate: 1.250000
	//old: 134217728  new: 167772160    rate: 1.250000

	//showResult(testSet)
	//showResult(testSet1)
	//showResult(testSet2)
	//old: 1  new: 2    rate: 2.000000
	//old: 2  new: 4    rate: 2.000000
	//old: 4  new: 8    rate: 2.000000
	//old: 8  new: 16    rate: 2.000000
	//old: 16  new: 32    rate: 2.000000
	//old: 32  new: 64    rate: 2.000000
	//old: 64  new: 128    rate: 2.000000
	//old: 128  new: 256    rate: 2.000000
	//old: 256  new: 512    rate: 2.000000
	//old: 512  new: 1024    rate: 2.000000
	//old: 1024  new: 1280    rate: 1.250000
	//old: 2048  new: 2560    rate: 1.250000
	//old: 4096  new: 5120    rate: 1.250000
	//old: 8192  new: 10240    rate: 1.250000
	//old: 16384  new: 20480    rate: 1.250000
	//old: 32768  new: 40960    rate: 1.250000
	//old: 65536  new: 81920    rate: 1.250000
	//old: 131072  new: 163840    rate: 1.250000
	//old: 262144  new: 327680    rate: 1.250000
	//old: 524288  new: 655360    rate: 1.250000
	//old: 1048576  new: 1310720    rate: 1.250000
	//old: 2097152  new: 2621440    rate: 1.250000
	//old: 4194304  new: 5242880    rate: 1.250000
	//old: 8388608  new: 10485760    rate: 1.250000
	//old: 16777216  new: 20971520    rate: 1.250000
	//old: 33554432  new: 41943040    rate: 1.250000
	//old: 67108864  new: 83886080    rate: 1.250000
	//old: 134217728  new: 167772160    rate: 1.250000
	//
	//old: 0  new: 1    rate: +Inf
	//old: 1  new: 2    rate: 2.000000
	//old: 3  new: 6    rate: 2.000000
	//old: 7  new: 14    rate: 2.000000
	//old: 15  new: 30    rate: 2.000000
	//old: 31  new: 64    rate: 2.064516
	//old: 63  new: 128    rate: 2.031746
	//old: 127  new: 256    rate: 2.015748
	//old: 255  new: 512    rate: 2.007843
	//old: 511  new: 1024    rate: 2.003914
	//old: 1023  new: 2048    rate: 2.001955
	//old: 2047  new: 2560    rate: 1.250611
	//old: 4095  new: 5120    rate: 1.250305
	//old: 8191  new: 10240    rate: 1.250153
	//old: 16383  new: 20480    rate: 1.250076
	//old: 32767  new: 40960    rate: 1.250038
	//old: 65535  new: 81920    rate: 1.250019
	//old: 131071  new: 163840    rate: 1.250010
	//old: 262143  new: 327680    rate: 1.250005
	//old: 524287  new: 655360    rate: 1.250002
	//old: 1048575  new: 1310720    rate: 1.250001
	//old: 2097151  new: 2621440    rate: 1.250001
	//old: 4194303  new: 5242880    rate: 1.250000
	//old: 8388607  new: 10485760    rate: 1.250000
	//old: 16777215  new: 20971520    rate: 1.250000
	//old: 33554431  new: 41943040    rate: 1.250000
	//old: 67108863  new: 83886080    rate: 1.250000
	//old: 134217727  new: 167772160    rate: 1.250000
	//
	//old: 0  new: 1    rate: +Inf
	//old: 2  new: 4    rate: 2.000000
	//old: 6  new: 12    rate: 2.000000
	//old: 14  new: 28    rate: 2.000000
	//old: 30  new: 60    rate: 2.000000
	//old: 62  new: 128    rate: 2.064516
	//old: 126  new: 256    rate: 2.031746
	//old: 254  new: 512    rate: 2.015748
	//old: 510  new: 1024    rate: 2.007843
	//old: 1022  new: 2048    rate: 2.003914
	//old: 2046  new: 2560    rate: 1.251222
	//old: 4094  new: 5120    rate: 1.250611
	//old: 8190  new: 10240    rate: 1.250305
	//old: 16382  new: 20480    rate: 1.250153
	//old: 32766  new: 40960    rate: 1.250076
	//old: 65534  new: 81920    rate: 1.250038
	//old: 131070  new: 163840    rate: 1.250019
	//old: 262142  new: 327680    rate: 1.250010
	//old: 524286  new: 655360    rate: 1.250005
	//old: 1048574  new: 1310720    rate: 1.250002
	//old: 2097150  new: 2621440    rate: 1.250001
	//old: 4194302  new: 5242880    rate: 1.250001
	//old: 8388606  new: 10485760    rate: 1.250000
	//old: 16777214  new: 20971520    rate: 1.250000
	//old: 33554430  new: 41943040    rate: 1.250000
	//old: 67108862  new: 83886080    rate: 1.250000
	//old: 134217726  new: 167772160    rate: 1.250000


	// 说明：
	//       老容量 <= 1024时，  扩容倍率约为2   (大于2)
	//       老容量 > 1024时，   扩容倍率约为1.25(大于1.25)
	// 猜测：可能涉及到了内存优化
}

//func main(){
//	s := make([]int,1,1)
//	//s = append(s,1) 		cap: 2
//	// s = append(s,1,2)  cap: 4
//	// s = append(s,1,2,3) cap: 4
//	//s = append(s,1,2,3,4) cap: 6
//	//s = append(s,1,2,3,4,5,6) cap: 8
//
//	p := make([]int,1,2)
//	//p = append(p,1)		cap: 2
//	//p = append(p,1,2)		cap: 4
//	//p = append(p,1,2,3)   cap: 4
//
//	k := make([]int,3,3)
//	//k = append(k,1,2,3,4,5,6) cap: 10
//
//	t := make([]int,3,3)
//	//t = append(t,1)
//	//t = append(t,2)
//	//t = append(t,3)
//	//t = append(t,4)
//	//t = append(t,5)
//	//t = append(t,6)   cap: 12
//
//
//	fmt.Println(cap(s))
//	fmt.Println(cap(p))
//	fmt.Println(cap(k))
//	fmt.Println(cap(t))
//
//	// 总结：向a切片插入b切片，和分别将b的元素插入a切片，a切片的扩容规则是不一样的
//}
