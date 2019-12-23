package main

import (
	"fmt"
	"net/http"
)

func one(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"welcome to my life~one")
}
func two(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"welcome to my life~two")
}
func main(){
	http.HandleFunc("/one",one)
	http.HandleFunc("/two",two)
	http.ListenAndServe(":1234",nil)
}
