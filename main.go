package main

import "net"

type Config struct {
	ListenerAddreess string
}
type Server struct {
	Config
	ln net.Listener
}

func NewSever(config Config) *Server {
	return &Server{
		Config: config,
	}
}
func main() {}
