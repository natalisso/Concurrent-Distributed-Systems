package main

import (
	"fmt"
	"Middleware/Distribution/invoker"
)

func main() {

	fmt.Println("Naming server running!!")

	// control loop passed to invoker
	namingInvoker := invoker.NamingInvoker{}
	namingInvoker.Invoke()
}
