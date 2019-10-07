package main

import (
	"Middleware/Distribution/proxies"
	"Middleware/Services/naming/proxy"
	"fmt"
	// "time"
)

func runClient() {
	// create a built-in proxy of naming service
	namingService := proxy.NamingProxy{}

	// look for a service in naming service
	calculator := namingService.Lookup("DataBase").(proxies.DataBaseProxy)

	// invoke remote operation
	for i := 0; i < 5; i++ {
		fmt.Println(calculator.Save("Nome", "85324", 1929))
	}

}

func main() {
	go runClient()
	fmt.Scanln()
}
