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

func isNewUser(id string, users []*net.UDPAddr) bool {

	for _, idCur := range users {
		if idCur.String() == id {
			return false
		}
	}
	
	return true
}

func sendMessageToAll(users *[]*net.UDPAddr, msg []byte, conn *net.UDPConn) {
	for _, idCur := range *users {
		_, err := conn.WriteToUDP(msg, idCur)
		checkError(err)
	}
}

func handleUDPConnection(conn *net.UDPConn, users *[]*net.UDPAddr) {

	// here is where you want to do stuff like read or write to client

	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)
	checkError(err)
	fmt.Println("UDP client : ", addr)
	fmt.Println("Received from UDP client :  ", string(buffer[:n]))

	if isNewUser(addr.String(), *users) {
		*users = append(*users, addr)
	}
	fmt.Println(users)

	// NOTE : Need to specify client address in WriteToUDP() function
	//        otherwise, you will get this error message
	//        write udp : write: destination address required if you use Write() function instead of WriteToUDP()

	// write message back to all clients
	sendMessageToAll(users, buffer, conn)
}

func main() {
	hostName := "localhost"
	portNum := "6000"
	service := hostName + ":" + portNum
	var users []*net.UDPAddr

	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	// setup listener for incoming UDP connection
	ln, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	fmt.Println("UDP server up and listening on port 6000")

	defer ln.Close()

	for {
		// wait for UDP client to connect
		handleUDPConnection(ln, &users)
	}

}
