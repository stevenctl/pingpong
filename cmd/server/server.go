package main

import (
	"github.com/stevenctl/pingpong/cmd/common"
	"github.com/stevenctl/pingpong/pkg/server"
	"log"
	"net"
)

var defaultArgs = common.Args{Host: "0.0.0.0", Port: "5000"}

func main() {
	args := common.ParseArgs(defaultArgs)
	log.Println(net.JoinHostPort(args.Host, args.Port))
	s, err := server.New(args.Host, args.Port)
	if err != nil {
		log.Printf("failed starting server: %v\n", err)
	}

	// intentionally never stop
	stop := make(chan struct{})
	defer close(stop)
	s.Run(stop)
}
