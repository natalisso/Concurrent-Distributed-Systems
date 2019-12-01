package main

import (
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/brokerproxy"
	"Concurrent-Distributed-Systems/NMiddleware/shared"
	"fmt"
)

// PRODUTOR
func main() {
	bp := brokerproxy.NewBrokerProxy("", true, shared.N_HOST_PB, 1818)
	//bp.ConnectionBroker()

	bp.Exchange_Declare("Direct-X", "direct")
	msg := "Olá, consumidor do publish2!\n"

	fmt.Println("Ready to send a message")
	fmt.Scanln()
	for i := 0; i < 10000; i++ {
		bp.Basic_Publish("Direct-X", "Key1", msg)
	}

	fmt.Scanln()
}
