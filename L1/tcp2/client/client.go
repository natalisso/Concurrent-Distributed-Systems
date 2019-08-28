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
	"sync"
)

var wg = &sync.WaitGroup{}
var NUMCLIENTS = 1

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

func getData(user *client) string{
	fmt.Printf("Entry the data: ")

	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	checkError(err)

	return data
}

func sendMessage(user *client, msg string){
		//msg := getData(user)
		_,err := (*user).conn.Write([]byte(msg))
		checkError(err)	
}

func receiveMessage(user *client){
	_, err := (*user).reader.ReadString('\n')
	checkError(err)

	// Escrevendo a resposta do servidor no terminal
	//fmt.Printf(menssage)	
}

func runClient(clientName string, server string) {
	// Conectando ao servidor
	conn, err := net.Dial("tcp", server)
	checkError(err)
	
	//fmt.Printf("Connected to server at %s!\n",server)

	user := client{clientName,conn,bufio.NewReader(conn)}
	//getName(&user)

	// Mandando o nome do client para o server e aguarda o retorno
	sendMessage(&user,"MSG " + user.name + "\n")
	receiveMessage(&user)

	x := 0.000000
	for i:=0; i <= 1E4; i++{
		if i == 1E4{
			sendMessage(&user,"STOP " + strconv.FormatFloat(x,'f',6,64) +"\n")
			receiveMessage(&user)
		}else{
			time1 := time.Now()
			sendMessage(&user,"MSG " + strconv.FormatFloat(x,'f',6,64) +"\n")
			receiveMessage(&user)
			time2 := time.Now()
			x = float64(time2.Sub(time1).Nanoseconds()) / 1E3
		}
	}
	wg.Done()
}

func main(){
	// Servidor na máquina local na porta 8080 (default)
	server := "127.0.0.1:8080" 

	// Pego o numero de clients e o endereço ip e a porta do servidor caso tenham sido passados como argumento 
	if len(os.Args) == 2 {
		NUMCLIENTS,_ = strconv.Atoi(os.Args[1])
	}else if len(os.Args) == 3{
		NUMCLIENTS,_ = strconv.Atoi(os.Args[1])
		server = os.Args[2]	
	}
	
	// Inicializando as threads dos clients 
	nome := "nomeClient"
	wg.Add(NUMCLIENTS)
	for i:=0; i < NUMCLIENTS; i++{
		go runClient(nome,server)
	}
	wg.Wait()
}

