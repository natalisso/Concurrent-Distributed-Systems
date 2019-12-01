package main

import (
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/brokerproxy"
	"Concurrent-Distributed-Systems/NMiddleware/shared"
	"log"
	"os"
)

// CONSUMIDOR
func main() {
	bp := brokerproxy.NewBrokerProxy("", true, shared.N_HOST_PB, shared.N_PORT_SB)
	//bp.ConnectionBroker()

	bp.Queue_Declare("Direct-Q")
	bp.Queue_Bind("Direct-X", "Direct-Q", "Key1")

	f, err := os.Create("./s")
	if err != nil {
		log.Fatalf("Error")
	}

	for i := 0; i < 10000; i++ {
		f.WriteString(bp.Basic_Consume("Direct-Q"))
		//fmt.Printf("Received: %s\n", bp.Basic_Consume("Direct-Q"))
	}
	f.Close()
}
