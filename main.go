package main

import (
	"log/slog"
	"net"
)

const defaultAddr = ":5666"

type Config struct {
	ListenerAddreess string
}
type Server struct {
	Config
	ln net.Listener
}

func NewSever(config Config) *Server {
	if len(config.ListenerAddreess) == 0 {
		config.ListenerAddreess = defaultAddr
	}
	return &Server{
		Config: config,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenerAddreess)
	if err != nil {
		return err
	}
	s.ln = ln
	s.AcceptLoop()
	return nil
}

func (s *Server) AcceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			slog.Error("err", err)
			continue
		}
	}
}
func main() {
	server := NewSever(Config{
		ListenerAddreess: ":8080",
	})
	server.Start()
}
