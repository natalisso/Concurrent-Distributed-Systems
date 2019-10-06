package srh

import (
	"net"
	"strconv"
	"log"
	"encoding/binary"
)

type SRH struct {
	ServerHost string
	ServerPort int
}

var ln net.Listener
var conn net.Conn
var err error

func (srh SRH) Receive() []byte {

	// create listener
	ln, err = net.Listen("tcp", srh.ServerHost+":"+strconv.Itoa(srh.ServerPort))
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	// accept connections
	conn, err = ln.Accept()
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	// receive message's size
	size := make([]byte, 4)
	_, err = conn.Read(size)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}
	sizeInt := binary.LittleEndian.Uint32(size)

	// receive message
	msg := make([]byte, sizeInt)
	_, err = conn.Read(msg)
	if err != nil {
		log.Fatalf("SRH:: %s", err)
	}

	return msg
}

func (SRH) Send(msgToClient []byte) {

	// send message's size
	size := make([]byte, 4)
	l := uint32(len(msgToClient))
	binary.LittleEndian.PutUint32(size, l)
	conn.Write(size)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	// send message
	_, err = conn.Write(msgToClient)
	if err != nil {
		log.Fatalf("CRH:: %s", err)
	}

	// close connection
	conn.Close()
	ln.Close()
}



