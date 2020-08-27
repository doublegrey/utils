package utils

import (
	"context"
	"net"
	"os"
	"syscall"
	"time"
)

// sock, err := uds.ConnectSocket("/env/example.sock", time.Second*1, time.Second*5)
//
// sock.Write(msg)
func ConnectSocket(address string, interval, deadline time.Duration) (*net.UnixConn, error) {
	type Result struct {
		conn *net.UnixConn
		err  error
	}
	ctx, cancel := context.WithTimeout(context.Background(), deadline)
	defer cancel()
	done := make(chan Result, 1)
	go func(done chan<- Result) {
		wait := func() bool {
			select {
			case <-time.After(interval):
				return true
			case <-ctx.Done():
				// Canceled, don't continue
				done <- Result{err: ctx.Err()}
				return false
			}
		}
		// Try connecting until either connected or canceled
		for {
			addr, err := net.ResolveUnixAddr("unix", address)
			if err != nil {
				if !wait() {
					return
				}
				continue
			}
			conn, err := net.DialUnix("unix", nil, addr)
			if err != nil {
				if !wait() {
					return
				}
				continue
			}
			// Connection established successfully
			done <- Result{conn: conn}
			break
		}
	}(done)
	r := <-done
	return r.conn, r.err
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
