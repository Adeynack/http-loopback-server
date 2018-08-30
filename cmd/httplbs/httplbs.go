package main

import (
	"flag"
	"fmt"
	"github.com/adeynack/http-loopback-server"
	"log"
)

const (
	defaultPort = 10001
)

func main() {
	port := *flag.Int("port", defaultPort, fmt.Sprintf("port on which to listen for incoming request (default: %d)", defaultPort))
	log.Printf("Port: %v", port)
	addr := fmt.Sprintf(":%v", port)
	http_loopback_server.Serve(addr)
}

