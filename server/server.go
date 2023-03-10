package server

type State uint8

const (
	Stop State = iota
	Start
)

type Server struct {
	Current State
	Port    int
}
