package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/korkmazkadir/coordinator/registery"
)

func main() {

	cordinator := new(registery.NodeRegistery)
	rpc.Register(cordinator)

	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	log.Println("Coordinator service started and listening the port 1234")

	// this is a single thread by design
	for {
		conn, _ := l.Accept()
		rpc.ServeConn(conn)
	}

}
