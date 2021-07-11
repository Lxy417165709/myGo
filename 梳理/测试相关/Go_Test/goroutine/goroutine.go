package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("part 1:")
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("%p %d\t", &i, i)
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println()

	fmt.Println("part 2:")
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%p %d\t", &i, i)
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println()

	fmt.Println("part 3:")
	for i := 0; i < 10; i++ {
		go func(i *int) {
			fmt.Printf("%p %p %d\t", &i, i, *i)
		}(&i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println()

	// part 1:
	// 0xc00000a0e0 9	0xc00000a100 0	0xc00000a108 1	0xc00000a118 2	0xc00000a128 3	0xc00000a138 4	0xc00000a148 5	0xc00000a158 6	0xc00000a168 7	0xc00000a178 8
	// part 2:
	// 0xc000112008 5	0xc000112008 8	0xc000112008 8	0xc000112008 8	0xc000112008 9	0xc000112008 10	0xc000112008 10	0xc000112008 10	0xc000112008 10	0xc000112008 10
	// part 3:
	// 0xc000130008 0xc000112008 10	0xc000130010 0xc000112008 10	0xc000130018 0xc000112008 10	0xc000130020 0xc000112008 10	0xc000130028 0xc000112008 10	0xc000130030 0xc000112008 10	0xc000130038 0xc000112008 10	0xc000130040 0xc000112008 10	0xc000130048 0xc000112008 10	0xc000130050 0xc000112008 10

	// 可见：
	// 		part 1: 每个协程 都持有     i变量  		---> (外i唯一)
	// 		part 2: 每个协程 持有同一个 i变量  		---> (内i多个) (外i唯一)
	//		part 3: 每个协程 都持有     i指针变量  	---> (内i多个) (外i唯一)
}
