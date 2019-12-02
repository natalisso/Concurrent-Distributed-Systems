package main

import (
	"NMiddleware/Middleware/Distribution/broker"
)

func main() {

	broker := broker.NewBroker()

	broker.Manager()

}
