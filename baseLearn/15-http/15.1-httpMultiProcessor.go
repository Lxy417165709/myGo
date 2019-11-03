package main

import (
	"fmt"
	"net/http"
)

type myHandle1 struct {
	Text string
}

func (hd *myHandle1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, hd.Text)
}

type myHandle2 struct {
	Text string
}

func (hd *myHandle2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, hd.Text)
}

func main() {
	h1 := myHandle1{"welcome to my life~"}
	h2 := myHandle2{"You are lovely~"}
	http.Handle("/one", &h1)
	http.Handle("/two", &h2)
	http.ListenAndServe(":8090", nil)
}

/*
	总结
	1. 多处理器就是用类的思维去接收和回复请求
	2. 多函数就是用函数的思维去接收和回复请求
	3. 推荐用多函数,比较方便
*/
