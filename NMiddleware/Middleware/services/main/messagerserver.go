package main

import (
	"NMiddleware/Middleware/Distribution/broker"
	"fmt"
)

func main() {

	broker := broker.NewBroker()
	fmt.Println("Messager Server is Running!!")

	broker.Manager()

}
