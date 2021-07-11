package 杂物

import (
	"fmt"
	"unsafe"
)

func test (sli []int){
	fmt.Println("函数内",unsafe.Pointer(&sli))	// 切片a的地址
	fmt.Printf("函数内 %p\n",sli)			// 切片a指向的地址
}


func main() {
	a :=make([]int,0,0)
	fmt.Println(unsafe.Pointer(&a))	// 切片a的地址
	fmt.Printf("%p\n",a)		// 切片a指向的地址
	test(a)
	// 0xc042002440
	// 0x551e08
	// 函数内 0xc0420024a0
	// 函数内 0x551e08



	b := a
	fmt.Println(unsafe.Pointer(&b))	// 切片b的地址
	fmt.Printf("%p\n",b)		// 切片b指向的地址


	// 0xc04204e3e0
	// 0x551e08
	// 0xc04204e440
	// 0x551e08

	a = append(a,1)

	fmt.Println(unsafe.Pointer(&a))	// 切片a的地址
	fmt.Printf("%p\n",a)		// 切片a指向的地址

	fmt.Println(unsafe.Pointer(&b))	// 切片b的地址
	fmt.Printf("%p\n",b)		// 切片b指向的地址


	// 0xc04204e3e0
	// 0xc0420540a0
	// 0xc04204e440
	// 0x551e08

	fmt.Println(unsafe.Pointer(&a))	// 切片a的地址
	fmt.Printf("%p\n",a)		// 切片a指向的地址
	a = a[:len(a)-1]
	fmt.Println(unsafe.Pointer(&a))	// 切片a的地址
	fmt.Printf("%p\n",a)		// 切片a指向的地址

	// 0xc042002440
	// 0xc042008098
	// 0xc042002440
	// 0xc042008098


	//fmt.Println(unsafe.Sizeof(a))
}

/*
	总结
	1. 每个变量在内存中的地址都是不一样的!(C语言中的引用则是地址都一样)
	2. 切片赋值时，赋值的是切片指向的地址，切片长度，切片容量，即切片结构的内容
	3. 切片使用append函数时，当超出切片容量，append会返回另外一个切片，该切片与原切片指向的地址是不一样的。
       如果之前把未append的切片a赋值给别的切片b，那此时b!=a
       如果没超出，a指向的地址不变，长度增加1，
       如果之前把未append的切片a赋值给别的切片b，那此时b==a
	4. a = a[:len(a)-1] 不会改变切片的地址以及指向的地址
	5. 切片作为参数时，传递的也是切片的内容，并非C语言的引用(即形参和实参所在内存地址一样)
*/