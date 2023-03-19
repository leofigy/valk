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

type ServerConfig struct {
	Current State
	Address string
}

type Server struct {
	ServerConfig
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

func InitBackendListener(state chan ServerConfig) {
	var server Server

	for {
		select {
		case newState := <-state:
			log.Println(newState)
			if server.ServerConfig.Current != newState.Current {
				log.Println("want to do a transition for ", newState.Current)
				switch newState.Current {
				case Start:
					server.ServerConfig = newState
				case Stop:
					server.ServerConfig = newState
				default:
					log.Println("stuck")
				}

			}
		case <-time.After(time.Second * 2):
			log.Println("listener nothing to do")
		}
	}
}
