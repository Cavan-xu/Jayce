package main

import (
	"fmt"
	"net"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8090")
	_, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
	}
}
