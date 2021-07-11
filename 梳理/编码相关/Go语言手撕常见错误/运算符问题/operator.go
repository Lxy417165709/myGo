package main

import "fmt"

// 目的: 输出int32的最大值 (go语言没有错，'<<'的优先级居然高于'-'号...)
// 这个没有错
func problemOne() {
	ans := 1<<31 - 1
	fmt.Println(ans)
	// 2147483647

	/*
		long long x = 1<<31 - 1;
		cout<<x<<endl;

		// C++采用这样的写法时，输出的是 1073741824
	*/
}

func main() {
	problemOne()

}
