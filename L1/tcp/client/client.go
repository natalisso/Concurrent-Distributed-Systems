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

func receiveMessage(user *client) string{
	message, err := (*user).reader.ReadString('\n')
	checkError(err)

	// Escrevendo a resposta do servidor no terminal
	//fmt.Printf(message)	
	return message
}

func runClient(clientName string, server string) {
	// Conectando ao servidor
	conn, err := net.Dial("tcp", server)
	checkError(err)
	
	//fmt.Printf("Connected to server at %s!\n",server)
	//getName(&user)

	// Abre arquivo de saida 
	nameDataBase := "./data_bases/dataBase" + clientName + "(" + strconv.Itoa(NUMCLIENTS) + ")" + ".csv"
    dataBase, err := os.Create(nameDataBase)
    if err != nil {
		panic(err)
    }
	user := client{clientName,conn,bufio.NewReader(conn)}

	
	x := float64(NUMCLIENTS)
	for i:=0; i < 1E4; i++{
		if i == 1E4 - 1{
			time1 := time.Now()
			sendMessage(&user,"STOP " + strconv.FormatFloat(x,'f',6,64) +"\n")
			message := receiveMessage(&user)
			time2 := time.Now()
			x = float64(time2.Sub(time1).Nanoseconds()) / 1E6
			if _, err := dataBase.Write([]byte(message)); err != nil {		
				panic(err)
			}
		}else if i == 0{
			// Mandando o dado inicial para o server e aguarda o retorno
			time1 := time.Now()
			sendMessage(&user,"MSG " + "data\n")
			message := receiveMessage(&user)
			time2 := time.Now()
			x = float64(time2.Sub(time1).Nanoseconds()) / 1E6
			if _, err := dataBase.Write([]byte(message)); err != nil {		
				panic(err)
		}
		}else{
			time1 := time.Now()
			sendMessage(&user,"MSG " + strconv.FormatFloat(x,'f',6,64) +"\n")
			message := receiveMessage(&user)
			time2 := time.Now()
			x = float64(time2.Sub(time1).Nanoseconds()) / 1E6
			if _, err := dataBase.Write([]byte(message)); err != nil {		
				panic(err)
			}
		}
	}
	wg.Done()
}

func main(){
	// Servidor na máquina local na porta 4093 (default)
	server := "127.0.0.1:4093" 

	// Pego o numero de clients e o endereço ip e a porta do servidor caso tenham sido passados como argumento 
	if len(os.Args) == 2 {
		NUMCLIENTS,_ = strconv.Atoi(os.Args[1])
	}else if len(os.Args) == 3{
		NUMCLIENTS,_ = strconv.Atoi(os.Args[1])
		server = os.Args[2]	
	}
	
	// Inicializando as threads dos clients 
	wg.Add(NUMCLIENTS)
	for i:=0; i < NUMCLIENTS; i++{
		nome := strconv.Itoa(i+1)
		go runClient(nome,server)
	}
	wg.Wait()
}

