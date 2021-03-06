package udtconn

import (
	"net"
	"time"

	udt "github.com/vikulin/go-udt/udt"
)

type conn struct {
	udtConn udt.Connection
}

func (c *conn) Read(b []byte) (int, error) {
	return c.udtConn.Read(b)
}

func (c *conn) Write(b []byte) (int, error) {
	return c.udtConn.Write(b)
}

// LocalAddr returns the local network address.
// needed to fulfill the net.Conn interface
func (c *conn) LocalAddr() net.Addr {
	return c.udtConn.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (c *conn) RemoteAddr() net.Addr {
	return c.udtConn.RemoteAddr()
}

func (c *conn) Close() error {
	return c.udtConn.Close()
}

func (c *conn) SetDeadline(t time.Time) error {
	return c.udtConn.SetDeadline(t)
}

func (c *conn) SetReadDeadline(t time.Time) error {
	return c.udtConn.SetReadDeadline(t)
}

func (c *conn) SetWriteDeadline(t time.Time) error {
	return c.udtConn.SetWriteDeadline(t)
}

var _ net.Conn = &conn{}
