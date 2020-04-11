package uds

import (
	"net"
	"os"
	"syscall"
	"time"
)

// sock, err := uds.Listen("/env/example.sock", time.Second*2)
//
// sock.Write(msg)
func Listen(address string, delay time.Duration) (*net.UnixConn, error) {
	for {
		addr, err := net.ResolveUnixAddr("unix", address)
		if err != nil {
			return nil, err
		}
		conn, err := net.DialUnix("unix", nil, addr)
		if err != nil { //FIXME
			time.Sleep(delay)
			continue
		}
		defer conn.Close()
		return conn, nil
	}
}

// listener, err := uds.OpenSocket("/env/example.sock")
func OpenSocket(a string) (*net.UnixListener, error) {
	err := syscall.Unlink(a)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}
	addr, err := net.ResolveUnixAddr("unix", a)
	if err != nil {
		return nil, err
	}
	listener, err := net.ListenUnix("unix", addr)
	if err != nil {
		return nil, err
	}
	return listener, nil
}
