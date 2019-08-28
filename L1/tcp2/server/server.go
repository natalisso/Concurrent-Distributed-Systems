package main

import (
	"net"
	"os"
	"bufio"
	"log"
	//"fmt"
	//"strings"
	"time"
)

type client struct {
    name string
	conn  net.Conn
	reader *bufio.Reader
}

func checkError(err error) {
	// Verificação de erros durante a troca de menssagens
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return 
	}
}

func receiveMessage(user *client) string {
	// Recebe o nome do client
	msg, err := (*user).reader.ReadString('\n')
	checkError(err)

	return msg
}

func sendMessage(user *client, msg string){
	(*user).conn.Write([]byte(string(msg)))	
}

func writeFile(user *client, data string, logMsg string, dataBase *os.File){
	t := time.Now().UTC()
	if _, err := dataBase.Write([]byte(t.Format("2006,01,02") + "," + data)); err != nil {
		// Retorna uma mensagem de status negativo em caso de erro
		logMsg = err.Error()
		sendMessage(user,logMsg)
		panic(err)
	}

	// Retorna uma mensagem de status possitivo
	sendMessage(user,logMsg)
}

func handleConn(conn net.Conn){
	// Cria um objeto do tipo client
	user := client{"",conn,bufio.NewReader(conn)}
	//log.Printf("Serving %s\n", conn.RemoteAddr().String())

	// Abre arquivo de saida 
	nameDataBase := "./data_bases/dataBase" + conn.RemoteAddr().String() + ".csv"
    dataBase, err := os.Create(nameDataBase)
    if err != nil {
		panic(err)
    }
    // Fecha o arquivo de saída no final da execução da função e verifica os erros retornados
    defer func() {
		if err := dataBase.Close(); err != nil {
			panic(err)
        }
	}()
		
	stp := true
	for stp{
		commandName, err := user.reader.ReadString(' ')
		checkError(err)
		
		switch commandName {
		case "MSG ":
			// Recebe os dados
			data := receiveMessage(&user)

			// Escreve no arquivo o dado recebido, verificando possíveis erros na escrita
			writeFile(&user,data,"Data Stored!\n",dataBase)
		case "STOP ":
			// Recebe os dados
			data := receiveMessage(&user)

			// Escreve no arquivo o dado recebido, verificando possíveis erros na escrita
			writeFile(&user,data,"STOP\n",dataBase)

			// Fecha a conexão e encerra o hadle desse client
			conn.Close()
			stp = false
		default:
			errMsg := "Unknown command!\n"
			sendMessage(&user,errMsg)
			stp = false
		}
	}
	return
}

func main() {
	// Inicializa o servidor na porta 8080 e protocolo TCP 
	// fmt.Println("Server waiting for connections...")
	port := ":8080"
	l, err1 := net.Listen("tcp", port)
	checkError(err1)

	// Fecha o socket no final da execução
	defer l.Close()

	for {
		// Aceita conexões e inicia a rotina de tratamento (goroutine)
		conn, err2 := l.Accept()
		checkError(err2)

		// Lida com a requisição do client em uma goroutine
		go handleConn(conn)
	}
}
