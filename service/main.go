package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/korkmazkadir/coordinator/registery"
)

func main() {

	cordinator := registery.NewNodeRegistery()
	rpc.Register(cordinator)

	hostName := getHostName()

	rpc.HandleHTTP()
	l, e := net.Listen("tcp", fmt.Sprintf("%s:1234", hostName))
	if e != nil {
		log.Fatal("listen error:", e)
	}

	log.Printf("Coordinator service started and listening %s:1234\n", hostName)

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}

}

func getHostName() string {

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	log.Printf("WARNING: uses empty host name rather than %s \n", hostname)

	//return hostname
	return ""
}
