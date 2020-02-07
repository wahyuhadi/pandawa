package services

import (
	"net"
	"time"
)

// check open port
// host and port with string
func CheckPort(host, port string) bool {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return false
	}
	if conn != nil {
		defer conn.Close()
		return true
	}

	return true
}
