package main

import (
	"fmt"
	"html/template"
	"net/http"
)
func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("baseLearn/15-http/view/index.html") //注意路径
	t.Execute(w, nil)

}
func login(w http.ResponseWriter,r *http.Request){
	// 1.获取表单数据方法一
	fmt.Println(r.FormValue("username"))
	fmt.Println(r.FormValue("password"))
	// 123456
	// 4567897

	// 2.获取表单数据方法二
	r.ParseForm()			// 一定要先解析,否则r.form == nil
	fmt.Println(r.Form)		// r.Form的类型是map[string] []string
	// map[username:[123456] password:[4567897]]

}
func main() {
	http.HandleFunc("/",home)
	http.HandleFunc("/login",login)
	http.ListenAndServe(":8989",nil)
}
