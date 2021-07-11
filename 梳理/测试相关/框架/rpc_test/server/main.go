package main

import (
	"log"
	"net"
	"net/rpc"
)

type Calculator struct{}

func (c *Calculator) FormResultOfMultiply(parameter map[string]int, result *int) error {
	*result = parameter["leftNum"] * parameter["rightNum"]
	return nil
}

type Checker struct{}

func (c *Checker) IsValidUser(parameters map[string]string, reply *bool) error {
	*reply = parameters["email"] == "123" && parameters["password"] == "456"
	return nil
}

func main() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatal(err)
		return
	}
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
		return
	}
	calculator := new(Calculator)
	if err := rpc.Register(calculator); err != nil {
		log.Fatal(err)
		return
	}
	checker := new(Checker)
	if err := rpc.Register(checker); err != nil {
		log.Fatal(err)
		return
	}
	rpc.Accept(inbound)
}
