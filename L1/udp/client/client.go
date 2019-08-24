package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

var wg = &sync.WaitGroup{}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func sendMenssage(conn *net.UDPConn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		log.Printf("Write a message: ")
		text, _ := reader.ReadString('\n')
		message := []byte(text)
		_, err := conn.Write(message)
		checkError(err)
	}
}

func receiveMenssage(conn *net.UDPConn) {
	buffer := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		checkError(err)
		fmt.Println("UDP Server : ", addr)
		fmt.Println("Received from UDP server : ", string(buffer[:n]))
	}
}

func main() {
	/* 	hostName := "localhost"
	   	portNum := "6000" */

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	RemoteAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	//LocalAddr := nil
	// see https://golang.org/pkg/net/#DialUDP

	conn, err := net.DialUDP("udp", nil, RemoteAddr)
	checkError(err)
	// note : you can use net.ResolveUDPAddr for LocalAddr as well
	//        for this tutorial simplicity sake, we will just use nil

	log.Printf("Established connection to %s \n", service)
	log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
	log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())
	log.Printf("------------------ Welcome to the Chat! :) -----------------------")

	defer conn.Close()

	/* // write a message to server
	message := []byte("Hello UDP server!")
	_, err = conn.Write(message) */
	wg.Add(1)
	go sendMenssage(conn)
	wg.Add(1)
	go receiveMenssage(conn)
	wg.Wait()
	/* if err != nil {
		log.Println(err)
	} */

	// receive message from server
	/* 	buffer := make([]byte, 1024)
	   	n, addr, err := conn.ReadFromUDP(buffer)

	   	fmt.Println("UDP Server : ", addr)
	   	fmt.Println("Received from UDP server : ", string(buffer[:n])) */

}
