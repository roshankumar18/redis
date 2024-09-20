package main

import "net"

type Peer struct {
	conn  net.Conn
	msgCh chan []byte
}

func NewPeer(conn net.Conn, msgCh chan []byte) *Peer {
	return &Peer{
		conn:  conn,
		msgCh: msgCh,
	}
}

func (p *Peer) readLoop() {
	for {
		buf := make([]byte, 1024)
		n, err := p.conn.Read(buf)
		if err != nil {
			break
		}
		copyBuf := make([]byte, n)
		copy(copyBuf, buf)
		p.msgCh <- copyBuf
	}
}
