package main

import (
	"Middleware/Distribution/invoker"
	"Middleware/Distribution/proxies"
	"Middleware/Services/naming/proxy"
	"fmt"
)

func main() {

	// create a built-in proxy of naming service
	namingProxy := proxy.NamingProxy{}

	// create a proxy of calculator service
	dataBase := proxies.NewDataBaseProxy()
	//converter := proxies.NewConverterProxy()

	// register service in the naming service
	namingProxy.Register("DataBase", dataBase)
	//namingProxy.Register("Converter", converter)

	// control loop passed to middleware
	fmt.Println("Calculator Server running!!")
	calculatorInvoker := invoker.NewDataBaseInvoker()
	//converterInvoker := invoker.NewConverter()

	go calculatorInvoker.Invoke()
	//go converterInvoker.Invoke()

	fmt.Scanln()
}
