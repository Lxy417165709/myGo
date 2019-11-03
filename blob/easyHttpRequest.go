package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/html; charset=UTF-8")
	fmt.Fprint(w,`<h>hello world!</h>`)
}
func other(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/html; charset=UTF-8")
	fmt.Fprint(w,`<h>I am other!</h>`)
}
func main() {
	http.HandleFunc("/",home)
	http.HandleFunc("/other",other)
	http.ListenAndServe("127.0.0.1:8090",nil)
}
