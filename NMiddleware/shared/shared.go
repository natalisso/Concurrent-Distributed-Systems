package shared

import (
	"log"
	"net"
	"strconv"
)

const SAMPLE_SIZE = 1000
const DATABASE_PORT = 1313
const NAMING_PORT = 1414
const MIOP_REQUEST = 1
const MIOP_REPLY = 2
const N_INSTANCES = 5

const N_HOST_MD = "localhost"
const N_HOST_PB = "localhost"
const N_HOST_SB = "localhost"
const N_PORT_MD = 1313
const N_PORT_PB = 1414
const N_PORT_SB = 1515

type Request struct {
	Op string
	P1 string
	P2 string
	P3 int
}

type Reply struct {
	Result []interface{}
}

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func ChecaErro(err error, msg string) {
	if err != nil {
		log.Fatalf("%s!!: %s", msg, err)
	}
	//fmt.Println(msg)
}

func FindNextAvailablePort() int { // TCP only
	i := 3000

	for i = 3000; i < 4000; i++ {
		port := strconv.Itoa(i)
		ln, err := net.Listen("tcp", ":"+port)

		if err == nil {
			ln.Close()
			break
		}
	}
	return i
}
