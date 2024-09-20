package main

import (
	"fmt"
	"log/slog"
	"net"
)

const defaultAddr = ":5666"

type Config struct {
	ListenerAddreess string
}
type Server struct {
	Config
	peers     map[*Peer]bool
	ln        net.Listener
	addPeerch chan *Peer
	msgCh     chan []byte
}

func NewServer(config Config) *Server {
	if len(config.ListenerAddreess) == 0 {
		config.ListenerAddreess = defaultAddr
	}
	return &Server{
		Config:    config,
		peers:     make(map[*Peer]bool),
		addPeerch: make(chan *Peer),
		msgCh:     make(chan []byte),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenerAddreess)
	if err != nil {
		return err
	}
	s.ln = ln
	fmt.Println("Listening on", s.ListenerAddreess)
	go s.loop()
	return s.AcceptLoop()

}

func (s *Server) loop() {
	for {
		select {
		case rawMsg := <-s.msgCh:
			fmt.Println("Received message", string(rawMsg))
		case peer := <-s.addPeerch:
			s.peers[peer] = true
		}
	}
}

func (s *Server) AcceptLoop() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			slog.Error("err", err)
			continue
		}
		go s.handleConn(conn)
	}
}
func (s *Server) handleConn(conn net.Conn) {

	peer := NewPeer(conn, s.msgCh)
	s.addPeerch <- peer
	fmt.Println("New peer connected", conn.RemoteAddr())
	peer.readLoop()
}
func main() {
	server := NewServer(Config{
		ListenerAddreess: ":8080",
	})
	server.Start()
}
