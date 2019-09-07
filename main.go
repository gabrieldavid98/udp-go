package main

import (
	"log"
	"net"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":8000")
	log.Println(err)

	udpConnection, err := net.ListenUDP("udp", udpAddr)
	log.Println("Udp server listen at ", udpAddr)

	for {
		buffer := make([]byte, 1024)
		n, _, _ := udpConnection.ReadFromUDP(buffer[0:])
		log.Println("Remote ip:", udpConnection.RemoteAddr())
		log.Println("Datos:", string(buffer[0:n]))
	}
}
