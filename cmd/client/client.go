package main

import (
	"github.com/stevenctl/pingpong/cmd/common"
	"github.com/stevenctl/pingpong/pkg/client"
	"log"
	"net"
	"time"
)

var defaultArgs = common.Args{Host: "localhost", Port: "5000", ClientArgs: common.ClientArgs{Interval: 5 * time.Second}}

func main() {
	args := common.ParseArgs(defaultArgs)
	addr := net.JoinHostPort(args.Host, args.Port)
	log.Printf("Connecting to %s\n", net.JoinHostPort(args.Host, args.Port))
	c, err := client.New(addr, args.Interval)
	if err != nil {
		log.Print("failed connecting: %v\n", err)
	}
	c.Run()
}
