package server

import (
	"log"
	"net"
	"time"
)

type State uint8

const (
	Stop State = iota
	Start
)

type Server struct {
	Current State
	Port    int
}

func (server *Server) HandleWR(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	readN, err := conn.Read(buf)

	if err != nil {
		log.Println("falta error :", err)
		return
	}

	conn.Write([]byte("accepted"))
	log.Println(string(buf[0:readN]))
}

func InitBackendListener(state chan State, addr chan string) {
	for {
		time.Sleep(time.Second * 2)
		log.Println("listener loop")
	}
}
