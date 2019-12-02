package main

import (
	"NMiddleware/Middleware/Distribution/brokerproxy"
	"NMiddleware/shared"
	"log"
	"os"
	"fmt"
)

// CONSUMIDOR
func main() {
	bp := brokerproxy.NewBrokerProxy("", true, shared.N_HOST_PB, shared.N_PORT_SB)
	//bp.ConnectionBroker()

	bp.Queue_Declare("Q1")
	bp.Queue_Bind("Fanout-X", "Q1", "Key1.*.com")

	f, err := os.Create("./s")
	if err != nil {
		log.Fatalf("Error")
	}

	for i := 0; i < 10000; i++ {
		// f.WriteString(bp.Basic_Consume("Topic-Q"))
		fmt.Printf("Received: %s\n", bp.Basic_Consume("Q1"))
	}
	f.Close()
}
