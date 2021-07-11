package main

import (
	"net"
	"net/rpc"
)

type RPCService struct {
}

func (r *RPCService) Hello(request string, reply *string) error {
	*reply = "Hello:" + request
	return nil
}
func (r *RPCService) CountSum(request []int, reply *int) error {
	var sum int
	for i := 0; i < len(request); i++ {
		sum += request[i]
	}
	*reply = sum
	return nil
}
func main() {
	var err error
	if err = rpc.RegisterName("RPCService", new(RPCService)); err != nil {
		panic(err)
	}
	var listener net.Listener
	if listener, err = net.Listen("tcp", ":1234"); err != nil {
		panic(err)
	}
	var conn net.Conn
	if conn, err = listener.Accept(); err != nil {
		panic(err)
	}
	rpc.ServeConn(conn)
}
