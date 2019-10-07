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
	dtBase := namingService.Lookup("DataBase").(proxies.DataBaseProxy)

	// invoke remote operation
	for i := 0; i < 100000; i++ {
		rep := dtBase.Save("Nome", "8532-4", 1929)
		fmt.Println(rep)
	}

}

func main() {
	go runClient()
	fmt.Scanln()
}
