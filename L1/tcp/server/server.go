package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"math/rand"
	"time"
	"os"
	"sync"
)
 	
type client struct {
    name string
	conn  net.Conn
	reader *bufio.Reader
}

var mutex = &sync.Mutex{}

func checkError(err error) bool{
	// Verificação de erros durante a troca de menssagens
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return false
	}else if err.Error() == "EOF cmd"{
		fmt.Println(err)
		return true
	}else{
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	return false
}

func wannaJoin(curClient *client, users *[] client){
	ret := "Welcome to the Chat, " + (*curClient).name + "! Type 'STOP' to exit \n"
	(*curClient).conn.Write([]byte(string(ret)))
	
	mutex.Lock()
	*users = append(*users,client{(*curClient).name,(*curClient).conn, (*curClient).reader})	
	defer mutex.Unlock()
}

func wannaLeave(curClient *client, users *[] client){
	var usersAux []client
	for _, user := range *users{
		if user.conn.RemoteAddr() != (*curClient).conn.RemoteAddr(){
			usersAux = append(usersAux,user)
		}
	} 
	*users = usersAux
}

func firstConnections(curClient *client) {
	// Recebe o nome do client
	name, err := (*curClient).reader.ReadString('\n')
	checkError(err)

	(*curClient).name = strings.Replace(name, "\n", "", 1)
	ret := "Hello " + (*curClient).name + "!\n"
	(*curClient).conn.Write([]byte(string(ret)))
	
}

func receiveMessage(curClient *client) string{
	// Recebe uma mensagem do cliente
	text, err := (*curClient).reader.ReadString('\n')
	checkError(err)

	text = (*curClient).name + ": " + text
	return text

}

func broadcastMessage(text string, users *[] client){
	fmt.Print(text)
	for _,c := range *users{
		c.conn.Write([]byte(string(text)))
	}
}

func handleConnection(conn net.Conn, users *[] client) {
	user := client{"",conn,bufio.NewReader(conn)}
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())

	stp := true
	for stp{
		user.reader = bufio.NewReader(conn)
		commandName, err := user.reader.ReadString(' ')
		ck := checkError(err)
		if ck == true && len(*users) == 1{
			stp = false
		}
		fmt.Println("cmd = ",commandName)
		switch commandName {
		case "NAME ":
			firstConnections(&user)		
		case "JOIN ":
			fmt.Println("join")
			wannaJoin(&user,users)
		case "MSG ":
			fmt.Println("msg")
			text := receiveMessage(&user)
			broadcastMessage(text, users)
		case "STOP ":
			fmt.Println("stop")
			conn.Write([]byte(string("STOP\n")))
			wannaLeave(&user,users)
			conn.Close()
			stp = false
		default:
			fmt.Println("errorr")
			errMsg := "Unknown command!\n"
			conn.Write([]byte(string(errMsg)))
		}
	}
	
}

func main() {
	// Inicializa o servidor na porta 8080 e protocolo TCP 
	fmt.Println("Server waiting for connections...")
	port := ":8081"
	l, err1 := net.Listen("tcp", port)
	checkError(err1)

	// Fecha o socket na saida
	defer l.Close()

		// Lida com a questão da concorrência
	rand.Seed(time.Now().Unix())
	var users [] client

	for {
		// Aceita conexões e inicia a rotina de tratamento (goroutine)
		conn, err2 := l.Accept()
		checkError(err2)
		fmt.Println("N USERS =",len(users))
		go handleConnection(conn, &users)

	}
}

