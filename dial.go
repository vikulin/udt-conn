package udtconn

import (
	"context"
	"net"

	udt "github.com/vikulin/go-udt/udt"
)


// Listen creates a UDT listener on the given network interface
func Listen(network, laddr string) (net.Listener, error) {
	ln, err := udt.ListenUDT(network, laddr)
	if err != nil {
		return nil, err
	}
	return &server{
		udtServer: ln,
	}, nil
}

// Listen creates a UDT listener on the given network interface
func ListenContext(ctx context.Context, network, laddr string) (net.Listener, error) {
	ln, err := udt.ListenUDTContext(ctx, network, laddr)
	if err != nil {
		return nil, err
	}
	return &server{
		udtServer: ln,
	}, nil
}

// Dial creates a new UDT connection
// it returns once the connection is established and secured with forward-secure keys
func Dial(raddr string) (net.Conn, error) {
	// DialAddr returns once a forward-secure connection is established
	addr, err := net.ResolveUDPAddr("udp", raddr)
	if err != nil {
		return nil, err
	}
	uConn, err := udt.DialUDT("udp", "0.0.0.0:0", addr, true)
	if err != nil {
		return nil, err
	}

	return &conn{
		udtConn:   uConn,
	}, nil
}

// DialAddrContext establishes a new UDT connection to a server using the provided context.
// See DialAddr for details.
func DialContext(ctx context.Context, raddr string,) (net.Conn, error) {
	addr, err := net.ResolveUDPAddr("udp", raddr)
	if err != nil {
		return nil, err
	}
	uConn, err := udt.DialUDTContext(ctx, "udp", "0.0.0.0:0", addr, true)
	if err != nil {
		return nil, err
	}

	return &conn{
		udtConn:   uConn,
	}, nil
}
