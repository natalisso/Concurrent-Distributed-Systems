package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

var wg = &sync.WaitGroup{}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func sendMessage(conn *net.UDPConn, name string) {
	log.Println("Type 'join' to participate :)")
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", 1)
		text = name + ": " + text
		message := []byte(text)
		_, err := conn.Write(message)
		checkError(err)
	}
}

func receiveMessage(conn *net.UDPConn) {
	buffer := make([]byte, 1024)

	for {
		n, _, err := conn.ReadFromUDP(buffer)
		checkError(err)
		fmt.Println(string(buffer[:n]))
	}
}

func main() {
	/* 	hostName := "localhost"
	   	portNum := "6000" */

	// pego o endereço do comando do terminal
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	RemoteAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	//LocalAddr := nil; https://golang.org/pkg/net/#DialUDP

	conn, err := net.DialUDP("udp", nil, RemoteAddr)
	checkError(err)

	log.Printf("Established connection to %s \n", service)
	log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
	log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

	reader := bufio.NewReader(os.Stdin)
	log.Printf("Type your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", 1)

	log.Printf("Welcome to the Chat, " + name + "! :)")
	defer conn.Close()

	// write a message to server
	wg.Add(1) // impedir que a thread main acabe antes que as outras
	go sendMessage(conn, name)
	wg.Add(1)
	go receiveMessage(conn)
	wg.Wait()
	
}
