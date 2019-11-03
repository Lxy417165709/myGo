package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"您看到我了")
}
func sayHello1(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"您看到我了")
}
func main() {
	http.HandleFunc("/",sayHello)
	http.HandleFunc("/1",sayHello1)
	log.Println("启动了")
	err := http.ListenAndServe(":9091",nil)
	if err != nil{
		log.Fatal("List 9091")
	}
}