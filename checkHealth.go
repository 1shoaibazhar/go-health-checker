package main

import (
	"fmt"
	"net"
	"time"
)

func CheckHealth(destination, port string) string {
	var status string
	address := fmt.Sprintf("%s:%s", destination, port)
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		status = fmt.Sprintf("[LINK IS DOWN] %s is unreachable :(\n Error %v\n", address, err)
	} else {
		status = fmt.Sprintf("[LINK IS UP] %s is reachable :)\n From: %s\n To: %s\n  ", address, conn.LocalAddr(), conn.RemoteAddr())
	}

	return status
}
