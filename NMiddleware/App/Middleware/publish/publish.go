package main

import (
	"NMiddleware/Middleware/Distribution/brokerproxy"
	"NMiddleware/shared"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func randFloats(min, max float64) float64 {
	res := min + rand.Float64()*(max-min)
	return res
}

// PRODUTOR
func main() {
	bp := brokerproxy.NewBrokerProxy("", true, shared.N_HOST_PB, shared.N_PORT_SB)

	bp.Exchange_Declare("Fanout-X", "fanout")
	msg := "Off\n"

	// Abre arquivo de saida
	nameDataBase := "./dataBaseFanout" + ".csv"
	// fmt.Println("ARQ = ",nameDataBase)

	dataBase, err := os.Create(nameDataBase)
	if err != nil {
		panic(err)
	}

	defer dataBase.Close()

	if _, err := dataBase.Write([]byte("data\n")); err != nil {
		panic(err)
	}

	fmt.Println("Ready to send a message")
	fmt.Scanln()
	for i := 0; i < 10000; i++ {
		t1 := time.Now()

		bp.Basic_Publish("Fanout-X", "Key1.edu.com", msg)

		t2 := time.Now()
		deltaTime := float64(t2.Sub(t1).Nanoseconds()) / 1E6
		if _, err := dataBase.Write([]byte(strconv.FormatFloat(deltaTime, 'f', 6, 64) + "\n")); err != nil {
			panic(err)
		}
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(randFloats(8, 12)) * time.Millisecond)
	}
	//fmt.Scanln()
}
