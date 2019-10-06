package main

import (
	"net/rpc"
	"time"
	"log"
	"fmt"
	"strconv"
)

const SAMPLE_SIZE = 10002
const NUM_CLIENTS = "1"
const SERVER_PORT = 1313


type Request struct {
	Header string
	RequestNumber int
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}


func main() {
	var reply string

	// Conectando ao servidor
	client, err := rpc.Dial("tcp", ":"+ strconv.Itoa(SERVER_PORT))
	checkError(err, "The server isn't ready")

	defer client.Close()

	// Invoca request
	for i := 0; i < SAMPLE_SIZE; i++ {

		// Prepara request
		msgRequest := Request{Header:"Request", RequestNumber: i}
		
		// Invoca request
		err := client.Call("Bank.ReceiveMessage", msgRequest, &reply)
		checkError(err, "Error communicating with server")

		fmt.Println(reply)
		time.Sleep(10 * time.Millisecond)
	}
}