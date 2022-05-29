package udtconn

import (
	"net"

	udt "github.com/vikulin/go-udt/udt"
)

type server struct {
	udtServer udt.Listener
}

var _ net.Listener = &server{}

// Accept waits for and returns the next connection to the listener.
func (s *server) Accept() (net.Conn, error) {
	return s.udtServer.Accept()
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (s *server) Close() error {
	return s.udtServer.Close()
}

// Addr returns the listener's network address.
func (s *server) Addr() net.Addr {
	return s.udtServer.Addr()
}
