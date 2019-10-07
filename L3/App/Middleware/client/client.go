package main

import (
	"Middleware/Services/naming/proxy"
	"Middleware/Distribution/proxies"
	"fmt"
	// "time"
)

func ExecuteExperiment() {
	// create a built-in proxy of naming service
	namingService := proxy.NamingProxy{}

	// look for a service in naming service
	calculator := namingService.Lookup("Calculator").(proxies.CalculatorProxy)

	// invoke remote operation
	for i := 0; i < 5; i++ {
		// t1 := time.Now()
		fmt.Println(calculator.Save("Nome", "12312312312", 29))
		// calculator.Add(1,2)
		// fmt.Println(time.Now().Sub(t1))
	}

}

func main() {
	go ExecuteExperiment()
	fmt.Scanln()
}