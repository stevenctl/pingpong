package server

import (
	"bufio"
	"log"
	"net"
)

type Server struct {
	listener net.Listener
}

// New starts listening on the given host and port.
func New(host, port string) (*Server, error) {
	l, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return nil, err
	}
	return &Server{listener: l}, nil
}

func (s *Server) conns() <-chan net.Conn {
	conns := make(chan net.Conn)
	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				log.Printf("failed accepting connection: %v\n", err)
			}
			conns <- conn
		}
	}()
	return conns
}

// Run starts accepting connections and handling them. Blocks until stop channel is closed.
func (s *Server) Run(stop <-chan struct{}) {
	log.Printf("listening on %s", s.listener.Addr())
	conns := s.conns()
	for {
		select {
		case conn := <-conns:
			go s.handleConn(conn)
		case <-stop:
			if err := s.listener.Close(); err != nil {
				log.Printf("failed closing listener\n")
			}
			return
		}
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer func() {
		log.Printf("[%s] closing\n", conn.RemoteAddr())
		_ = conn.Close()
	}()
	r := bufio.NewReader(conn)
	log.Printf("[%s] new connection\n", conn.RemoteAddr())
	for {
		if _, err := r.ReadBytes('\n'); err != nil {
			log.Printf("[%s] error reading ping: %v\n", conn.RemoteAddr(), err)
			return
		}
		if _, err := conn.Write([]byte("pong\n")); err != nil {
			log.Printf("[%s] error writing pong: %v\n", conn.RemoteAddr(), err)
			return
		}
	}
}
