package server

import (
	"crypto/tls"
	"log"
	"net"
	"sync"
	"time"
)

type State uint8

const (
	Stop State = iota
	Start
)

type ServerConfig struct {
	Current  State
	Address  string
	Security *tls.Config
}

type Server struct {
	ServerConfig
	listener net.Listener
	quit     chan any
	wg       sync.WaitGroup
}

func NewServer(cfg ServerConfig) (*Server, error) {
	server := &Server{
		ServerConfig: cfg,
		quit:         make(chan any),
	}

	log.Println("trying to listing in ", cfg.Address)
	var l net.Listener
	var err error
	if cfg.Security != nil {
		l, err = tls.Listen("tcp", cfg.Address, cfg.Security)
	} else {
		l, err = net.Listen("tcp", cfg.Address)
	}

	if err != nil {
		return nil, err
	}

	log.Println("setting listener")

	server.listener = l
	log.Println("wait group")
	server.wg.Add(1)
	go server.Serve()
	return server, nil
}

func (server *Server) Serve() {
	defer server.wg.Done()

	for {
		conn, err := server.listener.Accept()
		if err != nil {
			select {
			case <-server.quit:
				return
			default:
				log.Panicln("ERROR acepting connection", err)
			}
		} else {
			server.wg.Add(1)
			go func() {
				server.HandleWR(conn)
				server.wg.Done()
			}()
		}
	}
}

func (server *Server) Stop() {
	close(server.quit)
	server.listener.Close()
	server.wg.Wait()
}

func (server *Server) HandleWR(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)

	readN, err := conn.Read(buf)

	if err != nil {
		log.Println("falta error :", err)
		return
	}

	if readN == 0 {
		return
	}

	conn.Write([]byte("accepted"))
	log.Printf("received from %v: %s", conn.RemoteAddr(), string(buf[:readN]))
}

func InitBackendListener(state chan ServerConfig) {
	var server *Server
	var err error

	for {
		select {
		case newState := <-state:
			log.Println("---------- new target state", newState.Current)

			if server == nil && newState.Current == Start {
				log.Println("trying to do first start pal")
				server, err = NewServer(newState)
				if err != nil {
					log.Println("FATAL ERROR ", err)
				}
				continue
			}
			if server != nil && server.ServerConfig.Current != newState.Current {
				log.Println("want to do a transition for ", newState.Current)
				switch newState.Current {
				case Start:
					log.Println("STARTING AGAIN")
					server, err = NewServer(newState)
					if err != nil {
						log.Panicln("FATAL - Unable to start server pal")
						continue
					}
				case Stop:
					server.Stop()
					server.Current = Stop
				default:
					log.Println("stuck")
				}
			} else {
				log.Println("--- IGNORING ---- ")
			}
		case <-time.After(time.Second * 2):
			log.Println("listener nothing to do")
		}
	}
}
