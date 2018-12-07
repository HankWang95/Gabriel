package server

import "testing"

var (
	l *serverGRPC
)

func TestNewListener(t *testing.T) {
	l = NewListener()
}

func TestServerGRPC_Listen(t *testing.T) {
	l.Listen()
}