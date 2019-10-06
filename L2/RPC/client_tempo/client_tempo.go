package main

import (
	"net/rpc"
	"time"
	"log"
	// "fmt"
	"strconv"
	"os"
)

const SAMPLE_SIZE = 10000
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

	// Abre arquivo de saida 
	nameDataBase :="../../Analise_comparativa/RPC/dataBase"+NUM_CLIENTS+".csv"
	//fmt.Println("ARQ = ",nameDataBase)
	
	dataBase, err := os.Create(nameDataBase)
    if err != nil {
		panic(err)
	}

	defer dataBase.Close()
	
	if _, err := dataBase.Write([]byte("data\n")); err != nil {		
		panic(err)
	}

	// Invoca request
	for i := 0; i < SAMPLE_SIZE; i++ {
		t1 := time.Now()

		// Prepara request
		msgRequest := Request{Header:"Request", RequestNumber: i}
		
		// Invoca request
		err := client.Call("Bank.ReceiveMessage", msgRequest, &reply)
		checkError(err, "Error communicating with server")

		// fmt.Println(reply)
		t2 := time.Now()
		deltaTime := float64(t2.Sub(t1).Nanoseconds()) / 1E6
		if _, err := dataBase.Write([]byte(strconv.FormatFloat(deltaTime,'f',6,64)+"\n")); err != nil {		
			panic(err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}