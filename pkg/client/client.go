package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type Client struct {
	conn     net.Conn
	interval time.Duration
}

// New connects to the remote server
func New(remoteAddr string, interval time.Duration) (*Client, error) {
	conn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn, interval: interval}, nil
}

// Run starts the ping/pong/sleep cycle
func (c *Client) Run() {
	r := bufio.NewReader(c.conn)
	for {
		if _, err := c.conn.Write([]byte("ping\n")); err != nil {
			fmt.Println(err)
			return
		}
		log.Println("> ping")
		if got, err := r.ReadString('\n'); err == nil {
			log.Print("< " + got)
		} else {
			fmt.Println(err)
			return
		}
		time.Sleep(c.interval)
	}
}
