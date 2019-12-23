package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func home2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("baseLearn/15-http/view/index_cookies.html")
	t.Execute(w, nil)
}
func setCookies(w http.ResponseWriter, r *http.Request) {
	// Name为键,Value为值,HttpOnly为true防止XSS攻击,当其为false时,Cookies可以通过js获取
	// Path默认为"/",表示项目的所有目录都可以访问Cookies
	// Expires设置Cookies过期的时间,默认是一次浏览器关闭时清除Cookies
	c := http.Cookie{
		Name: "myKey",
		Value: "myValue",
		HttpOnly:true,
		Path:"/abc/bcd",
		Expires:time.Now().Add(10e9),	// 10s后过期
	}
	fmt.Println(c)
	http.SetCookie(w, &c)

}
func getCookies(w http.ResponseWriter, r *http.Request) {
	// 1.根据key取出Cookie
	c,_:=r.Cookie("myKey")
	fmt.Fprint(w,c)

	// 2.取出所有Cookies
	fmt.Fprint(w,r.Cookies())
}

func main() {
	http.HandleFunc("/", home2)
	http.HandleFunc("/setCookies", setCookies)
	http.HandleFunc("/getCookies", getCookies)
	http.HandleFunc("/abc/bcd", getCookies)	// 访问这个路由才可以查看自己设置的Cookies
	http.ListenAndServe(":8989", nil)

}

/*
	总结
	1. Name为键,Value为值,
	2. HttpOnly为true防止XSS攻击,当其为true时,Cookies不可以通过js获取,反之则可以获取
	3. Path默认为"/",表示项目的所有目录都可以访问Cookies
	4. Expires设置Cookies过期的时间,默认是一次浏览器关闭时清除Cookies
 */