package main 

import (
	"net"
	"net/rpc"
	"strconv"
	"log"
)

const SERVER_PORT = 1313

type RequestBank struct {}

type Request struct {
	Header string
	RequestNumber int
}

func (t *RequestBank) ReceiveMessage(req *Request, reply *string) (error) {
	*reply = "Here is your answer for the request " + strconv.Itoa(req.RequestNumber) + "!"

	return nil
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalf("FATAL ERROR -> %s: %s", msg, err)
	}
}


func main() {
	// cria instância do banco
	bank := new(RequestBank)
	
	// cria um novo servidor rpc e registra o banco
	server := rpc.NewServer()
	server.RegisterName("Bank", bank)

	// Cria um listen rpc-sender
	l, err := net.Listen("tcp", ":"+strconv.Itoa(SERVER_PORT))
	checkError(err, "Couldn't create the server")

	// Aguarda por chamadas
	log.Println("Server is ready (RPC TCP) ...")
	server.Accept(l)
}