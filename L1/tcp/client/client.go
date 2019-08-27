package main

import (
	"net"
	"os"
	"fmt"
	"sync"
	"bufio"
)

var wg = &sync.WaitGroup{}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func sendMessage(conn net.Conn){
	for{
		// Lendo entrada do teclado
		reader := bufio.NewReader(os.Stdin)
		menssage, err := reader.ReadString('\n')
		checkError(err)

		if menssage == "STOP\n"{
			conn.Write([]byte(string("STOP \n")))
			bresync.WaitGroup{}ak
		}
		// Escrevendo a mensagem na conexão
		conn.Write([]byte(string("MSG "+menssage+"\n")))
	}
	wg.Done()
}

func receiveMessage(conn net.Conn) {
	for{
		// Ouvindo a resposta do servidor
		menssage, err := bufio.NewReader(conn).ReadString('\n')
		checkError(err)

		// Escrevendo a resposta do servidor no terminal
		if menssage == "STOP\n"{
			fmt.Print("stop client!")
			break
		}
		fmt.Printf(menssage)	
	}
	wg.Done()
}

func getName(conn net.Conn){
	fmt.Printf("Type your name: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	checkError(err)

	conn.Write([]byte(string("NAME "+text+"\n")))

	menssage, err := bufio.NewReader(conn).ReadString('\n')
	checkError(err)
	fmt.Printf(menssage)
}

func startChat(conn net.Conn){
	fmt.Println("Type 'JOIN' to participate :)")

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	checkError(err)
	fmt.Print(text)
	if text == "JOIN\n" {
		conn.Write([]byte(string("JOIN \n")))
	}else{
		fmt.Println("errouuu")
		conn.Write([]byte(string("ERR \n")))
	}
}

func main() {

	// Servidor na máquina local na porta 8080 (default)
	server := "127.0.0.1:8081"
	// Pego o endereço ip e a porta do servidor caso tenham sido passados como argumento
	if len(os.Args) == 2 {
		server = os.Args[1]	
	}

	// Conectando ao servidor
	conn, err := net.Dial("tcp", server)
	checkError(err)

	getName(conn)
	startChat(conn)

	received, err := bufio.NewReader(conn).ReadString('\n')
	checkError(err)
	for received == "Unknown command!\n"{
		startChat(conn)
		received, err = bufio.NewReader(conn).ReadString('\n')
		checkError(err)
	}

	wg.Add(1)
	go sendMessage(conn)
	wg.Add(1)
	go receiveMessage(conn)
	wg.Wait()
		
}
