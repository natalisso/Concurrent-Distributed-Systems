package main

import (
	"log"
	"net"
	"net/rpc"
	"strconv"
)

const SERVER_PORT = 1313

type Book struct {
	Title string
	ISBN  string
	Year  int
}

type dataBase struct {
	books []Book
}

func (t *dataBase) Save(req *Book, reply *bool) error {
	t.books = append(t.books, *req)
	*reply = true
	return nil
}

func (t *dataBase) Search(req *Book, reply *bool) error {
	found := false
	for i := 0; i < len(t.books); i++ {
		if t.books[i].Title == req.Title {
			found = true
		}
	}
	*reply = found

	return nil
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalf("FATAL ERROR -> %s: %s", msg, err)
	}
}

func main() {
	// cria instância do banco
	dtBase := new(dataBase)

	// cria um novo servidor rpc e registra o banco
	server := rpc.NewServer()
	server.RegisterName("DataBase", dtBase)

	// Cria um listen rpc-sender
	l, err := net.Listen("tcp", ":"+strconv.Itoa(SERVER_PORT))
	checkError(err, "Couldn't create the server")

	// Aguarda por chamadas
	log.Println("Server is ready (RPC TCP) ...")
	server.Accept(l)
}
