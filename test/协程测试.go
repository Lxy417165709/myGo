package main

func Inc() (v int) {
	defer func(){ v++ } ()
	return 42
}

func main() {
	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		println(&i)
		defer func(i int){ println(i) ;println(&i)} (i)
	}
}

