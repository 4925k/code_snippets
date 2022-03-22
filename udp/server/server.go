package main

import (
	"fmt"
	"net"
)

var (
	host = "0.0.0.0"
	port = 6969
)

func main() {
	addr := net.UDPAddr{IP: net.ParseIP(host), Port: port}
	handleServer(&addr)
}

func handleServer(addr *net.UDPAddr) {

	connection, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 10240)

	for {
		n, _, err := connection.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n-1]))

		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
