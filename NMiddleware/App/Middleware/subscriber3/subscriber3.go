package main

import (
	"NMiddleware/Middleware/Distribution/brokerproxy"
	"NMiddleware/shared"
	"log"
	"fmt"
	"os"
)

// CONSUMIDOR
func main() {
	bp := brokerproxy.NewBrokerProxy("", true, shared.N_HOST_PB, 1616)
	//bp.ConnectionBroker()

	bp.Queue_Declare("Q3")
	bp.Queue_Bind("Direct-X", "Q3", "Key_dir")

	f, err := os.Create("./s2")
	if err != nil {
		log.Fatalf("Error")
	}

	for i := 0; i < 10000; i++ {
		// f.WriteString(bp.Basic_Consume("Topic-Q"))
		fmt.Printf("Received: %s\n", bp.Basic_Consume("Q3"))
	}
	f.Close()
}
