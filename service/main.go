package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/korkmazkadir/coordinator/registery"
)

func main() {

	cordinator := registery.NewNodeRegistery()
	rpc.Register(cordinator)

	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	log.Println("Coordinator service started and listening the port 1234")

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}

}
