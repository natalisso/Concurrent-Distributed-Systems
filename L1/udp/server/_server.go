package main

import (
    "fmt"
    "net"
    "os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleClientUDP(conn net.PacketConn){
	defer conn.Close()
	buf := make([]byte, 50)
	_, addr, err := conn.ReadFrom(buf)
	checkError(err)
	fmt.Println(addr, ">", string(buf))
}	

func main() {
	service := ":8080"

	for {
		pc, err := net.ListenPacket("udp", service)
		checkError(err)
		handleClientUDP(pc)
	}

}