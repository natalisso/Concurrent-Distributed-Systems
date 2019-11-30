package main

import (
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/brokerproxy"
	"Concurrent-Distributed-Systems/NMiddleware/shared"
	"fmt"
)

// PRODUTOR
func main() {
	bp := brokerproxy.NewBrokerProxy("", true, shared.N_HOST_PB, shared.N_PORT_SB)
	//bp.ConnectionBroker()

	bp.Exchange_Declare("Direct-X", "direct")
	msg := "Olá, consumidor!"

	fmt.Println("Ready to send a message")
	fmt.Scanln()
	bp.Basic_Publish("Direct-X", "Key1", msg)

	fmt.Scanln()
}
