package udtconn

import (
	"net"
	"time"

	udt "github.com/vikulin/go-udt"
)

type conn struct {
	udtConn udt.udtSocket
}

func newConn(udtConn udt.udtSocket) (*conn, error) {
	
	return &conn{
		udtConn:    udtConn,
	}, nil
}

func (c *conn) Read(b []byte) (int, error) {
	return c.Read(b)
}

func (c *conn) Write(b []byte) (int, error) {
	return c.Write(b)
}

// LocalAddr returns the local network address.
// needed to fulfill the net.Conn interface
func (c *conn) LocalAddr() net.Addr {
	return c.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (c *conn) RemoteAddr() net.Addr {
	return c.RemoteAddr()
}

func (c *conn) Close() error {
	return c.Close()
}

func (c *conn) SetDeadline(t time.Time) error {
	return c.SetDeadline(t)
}

func (c *conn) SetReadDeadline(t time.Time) error {
	return c.SetReadDeadline(t)
}

func (c *conn) SetWriteDeadline(t time.Time) error {
	return c.SetWriteDeadline(t)
}

var _ net.Conn = &conn{}
