package main

import (
	"net/http"
)


func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)

	// 浏览器访问: http://localhost:8080/static/img/3.gif 就能获得资源了
}

