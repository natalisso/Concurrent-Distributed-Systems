package shared

import (
	"log"
)

const SAMPLE_SIZE = 1000

type Request struct {
	Header string
	RequestNumber int
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}