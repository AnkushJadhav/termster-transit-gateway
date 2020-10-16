package websocket

import (
	"fmt"
	"net"
	"github.com/gorilla/websocket"
)

// WSServer represent a websocket server
type WSServer struct {
	bindport uint
	bindip net.IP
}

// New creates a websocket server that will listen on port p and ip
func New(ip string, p uint) (*WSServer, error) {
	frmIP := net.ParseIP(ip)
	if frmIP == nil {
		return nil ,fmt.Errorf("%s is not a valid IP address", ip)
	}

	return &WSServer{
		bindport: p,
		bindip: frmIP,
	}, nil
}