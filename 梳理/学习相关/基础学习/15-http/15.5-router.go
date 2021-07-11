package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func demo(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	fmt.Fprint(w,"这是不普通的~","url=",vars["url"])
	// 访问 http://localhost:8090/sxt/lxy
	// 输出 这是不普通的~url=lxy
}
func test(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"这是普通的~")
	// 访问 http://localhost:8090/test
	// 输出 这是普通的~
}
func main() {
	r:=mux.NewRouter()
	r.HandleFunc("/sxt/{url}",demo)
	r.HandleFunc("/test",test)
	http.ListenAndServe(":8090",r)

}

/*
	总结
	1. r.HandleFunc("/sxt/{url}",demo) 访问/sxt目录是404的,必须是/sxt/xxx才会进入demo
	2. 其他的看上面就懂啦,宝贝儿~
*/
