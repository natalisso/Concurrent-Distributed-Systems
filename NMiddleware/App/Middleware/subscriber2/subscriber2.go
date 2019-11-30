package main

import (
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/brokerproxy"
	"Concurrent-Distributed-Systems/NMiddleware/shared"
	"fmt"
)

// CONSUMIDOR
func main() {
	bp := brokerproxy.NewBrokerProxy("", true, shared.N_HOST_PB, shared.N_PORT_SB)
	//bp.ConnectionBroker()

	bp.Queue_Declare("Direct-Q")
	bp.Queue_Bind("Direct-X", "Direct-Q", "Key1")

	for true {
		fmt.Printf("Received: %s\n", bp.Basic_Consume("Direct-Q"))
	}
}
