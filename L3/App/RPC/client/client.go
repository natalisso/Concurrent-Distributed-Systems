package main

import (
	"fmt"
	"log"
	"net/rpc"
	"strconv"
	"time"
)

const SAMPLE_SIZE = 10002
const NUM_CLIENTS = "1"
const SERVER_PORT = 1313

type Book struct {
	Title string
	ISBN  string
	Year  int
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var reply bool

	// Conectando ao servidor
	client, err := rpc.Dial("tcp", ":"+strconv.Itoa(SERVER_PORT))
	checkError(err, "The server isn't ready")

	defer client.Close()

	// Invoca request
	for i := 0; i < SAMPLE_SIZE; i++ {

		// Prepara request
		msgRequest := Book{Title: "Name", ISBN: "83123-2", Year: 1923}

		// Invoca request
		err := client.Call("DataBase.Save", msgRequest, &reply)
		checkError(err, "Error communicating with server")

		fmt.Println(reply)
		time.Sleep(10 * time.Millisecond)
	}
}
