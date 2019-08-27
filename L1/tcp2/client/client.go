package main

import (
	"net"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "time"
)

type client struct {
    name string
	conn  net.Conn
	reader *bufio.Reader
}

func checkError(err error) {
	if err != nil {
		log.Printf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func getName(user *client){
	fmt.Printf("Type your name: ")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	checkError(err)

	(*user).name = strings.Replace(name, "\n", "", 1)

}

func sendMessage(user *client, msg string){
	if msg == "STOP"{
		_,err := (*user).conn.Write([]byte(string("STOP \n")))
		checkError(err)
	}else{
		// getName(user)
		// (*user).name = strings.Replace(msg, "\n", "", 1)
		_,err := (*user).conn.Write([]byte(string("MSG " + msg)))
		checkError(err)
	}
}

func receiveMessage(user *client){
	menssage, err := (*user).reader.ReadString('\n')
	checkError(err)

	// Escrevendo a resposta do servidor no terminal
	fmt.Printf(menssage)	
}

func main() {
	// Servidor na máquina local na porta 8080 (default)
	server := "127.0.0.1:8080"
	// Pego o endereço ip e a porta do servidor caso tenham sido passados como argumento
	if len(os.Args) == 2 {
		server = os.Args[1]	
	}

	// Conectando ao servidor
	conn, err := net.Dial("tcp", server)
	checkError(err)
	
	fmt.Printf("Connected to server at %s!\n",server)

	user := client{"nat",conn,bufio.NewReader(conn)}
	// getName(&user)

	for i:=0; i < 5; i++{
		if i == 4{
			// time1 := time.Now()
			sendMessage(&user,"STOP")
			receiveMessage(&user)
			// time2 := time.Now()
			// x := (float64) time2.Sub(time1).Nanoseconds / 1E6
		}else{
			sendMessage(&user,user.name+"\n")
			receiveMessage(&user)
		}

	}

}