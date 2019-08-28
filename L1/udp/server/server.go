package main

import (
	//"fmt"
	"log"
	"net"

	// "strings"
	//"math/rand"
	"os"
	"strings"
	//"sync"
	"time"
)

var users []client

type client struct {
	name string
	conn net.Conn
	addr *net.UDPAddr
}

func isNewUser(id string, users []client) bool {

	// varro o slice de client para ver se o usuário já foi conectado
	for _, idCur := range users {
		if idCur.addr.String() == id {
			return false
		}
	}
	return true
}

func checkError(err error) {
	// Verificação de erros durante a troca de menssagens
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return
	}
}

func receiveMessage(conn *net.UDPConn) (string, *net.UDPAddr) {
	// Recebe a mensagem do cliente
	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)
	checkError(err)
	//fmt.Println("receiveMessage: ", string(buffer[:n]))

	return string(buffer[:n]), addr
}

func sendMessage(addr *net.UDPAddr, msg string, conn *net.UDPConn) {
	msgToClient := []byte(msg)
	_, err := conn.WriteToUDP(msgToClient, addr)
	checkError(err)
}

func handleConn(conn *net.UDPConn, msg string, addr *net.UDPAddr) {

	var user client
	var dataBase *os.File

	i := strings.IndexByte(msg, ' ')
	commandName := msg[:i]
	//fmt.Println("Comando separado", commandName)

	switch commandName {
	case "MSG":

		// Verifica se ta salvo
		if isNewUser(addr.String(), users) {
			user = client{msg[i+1:], conn, addr}
			users = append(users, user)

			// open output file
			nameDataBase := "./data_bases/dataBase" + addr.String() + ".csv"
			var err error
			dataBase, err = os.Create(nameDataBase)
			checkError(err)

		} else {
			var err error
			dataBase, err = os.OpenFile("./data_bases/dataBase"+addr.String()+".csv", os.O_APPEND|os.O_WRONLY, 0600)
			checkError(err)
		}

		data := msg[i+1:]
		//fmt.Println("dado da mensagem:", data)
		// Escrevo no arquivo o que foi recebido junto com um formato de tempo
		t := time.Now().UTC()
		if _, err := dataBase.Write([]byte(t.Format("2006,01,02") + "," + data)); err != nil {
			panic(err)
		}

		// Fecho o arquivo depois de utilizá-lo
		if err := dataBase.Close(); err != nil {
			panic(err)
		}

		logMsg := "Data Stored!\n"
		sendMessage(addr, logMsg, conn)

	case "STOP":

		// Verficia se ta salvo
		if isNewUser(addr.String(), users) {
			user = client{msg[i+1:], conn, addr}
			users = append(users, user)

			// open output file
			nameDataBase := "./data_bases/dataBase" + addr.String() + ".csv"
			var err error
			dataBase, err = os.Create(nameDataBase)
			checkError(err)
		} else {
			var err error
			dataBase, err = os.OpenFile("./data_bases/dataBase"+addr.String()+".csv", os.O_APPEND|os.O_WRONLY, 0600)
			checkError(err)
		}
		data := msg[i+1:]
		//fmt.Println("dado da mensagem:", data)

		// Escrevendo no arquivo
		t := time.Now().UTC()
		if _, err := dataBase.Write([]byte(t.Format("2006,01,02") + "," + data)); err != nil {
			panic(err)
		}

		if err := dataBase.Close(); err != nil {
			panic(err)
		}
		stopSign := "STOP\n"
		sendMessage(addr, stopSign, conn)

	default:
		errMsg := "Unknown command!\n"
		sendMessage(addr, errMsg, conn)
	}
	return
}

func main() {
	// Inicializa o servidor na porta 8080 e protocolo UDP
	//fmt.Println("Server waiting for connections...")
	port := ":8080"

	service := "localhost" + port
	// Retorna um endereço de udp
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	// Se conecta ao endereço na rede nomeada. (udpAddr)
	// Conecta o servidor
	l, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	// Fecha o socket na saida
	defer l.Close()

	for {
		// ao receber uma nova mensagem cria uma thread para ministrar o que foi recebido
		msgFromClient, addr := receiveMessage(l)
		go handleConn(l, msgFromClient, addr)
	}
}
