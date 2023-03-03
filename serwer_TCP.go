package main

import (
	"fmt"
	"net"
)

type Srver struct {
	ListenAddres string
	ln           net.Listener
	qiitch       chan struct{}
}

func NewServer(ListenAddres string) *Srver {
	return &Srver{
		ListenAddres: ListenAddres,
		qiitch:       make(chan struct{}),
	}
}
func (s *Srver) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddres)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln
	<-s.qiitch
	return nil
}
func (s *Srver) acceptloop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}
		go s.readloop(conn)
	}
}
func (s *Srver) readloop(conn net.Conn) {
	defer conn.Close()
	bufor := make([]byte, 2048)
	for {
		n, err := conn.Read(bufor)
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}
		msge := bufor[:n]
		fmt.Println(string(msge))
	}
}

func main() {
	server := NewServer(":3000")
	server.Start()
}
