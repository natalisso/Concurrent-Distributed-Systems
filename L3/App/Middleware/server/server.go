package main

import (
	"fmt"
	"Middleware/Distribution/proxies"
	"Middleware/Services/naming/proxy"
	"Middleware/Distribution/invoker"
)

func main() {

	// create a built-in proxy of naming service
	namingProxy := proxy.NamingProxy{}

	// create a proxy of calculator service
	calculator := proxies.NewCalculatorProxy()
	//converter := proxies.NewConverterProxy()

	// register service in the naming service
	namingProxy.Register("Calculator", calculator)
	//namingProxy.Register("Converter", converter)

	// control loop passed to middleware
	fmt.Println("Calculator Server running!!")
	calculatorInvoker := invoker.NewCalculatorInvoker()
	//converterInvoker := invoker.NewConverter()

	go calculatorInvoker.Invoke()
	//go converterInvoker.Invoke()

	fmt.Scanln()
}

