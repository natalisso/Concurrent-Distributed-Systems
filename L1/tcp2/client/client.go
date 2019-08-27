package main

import (
	"net"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"time"
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

		// getName(user)
		// (*user).name = strings.Replace(msg, "\n", "", 1)
		_,err := (*user).conn.Write([]byte(msg))
		checkError(err)
	
}

func receiveMessage(user *client){
	menssage, err := (*user).reader.ReadString('\n')
	checkError(err)

	// Escrevendo a resposta do servidor no terminal
	fmt.Printf(menssage)	
}

func main() {
	// Servidor na máquina local na porta 8080 (default)
	server := "172.22.69.224:8080"
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

	// Mandando o nome
	sendMessage(&user,"MSG " + user.name+"\n")
	receiveMessage(&user)

	x := 0.0000000
	for i:=0; i <= 6; i++{
		if i == 6{
			sendMessage(&user,"STOP " + strconv.FormatFloat(x,'f',6,64) +"\n")
			receiveMessage(&user)
		}else{
			time1 := time.Now()
			sendMessage(&user,"MSG " + strconv.FormatFloat(x,'f',6,64) +"\n")
			receiveMessage(&user)
			time2 := time.Now()
			x = float64(time2.Sub(time1).Nanoseconds()) / 1E6
		}
	}
}