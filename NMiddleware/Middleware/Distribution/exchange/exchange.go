package exchange

import (
	"NMiddleware/Middleware/Distribution/bind"
	"fmt"
	//"fmt"
)

type Exchange struct {
	Name    string
	Type    string
	Durable bool
	Bind    bind.Bind
}

func NewExchange(typ string, durable bool) Exchange {
	ex := new(Exchange)
	ex.Bind = bind.NewBind()
	ex.Type = typ
	ex.Durable = durable
	return *ex
}

// Find
func (ex Exchange) FindQueues(bindKey string) []string {
	// Retorna os nomes das filas
	var nameQueues []string
	if ex.Type == "" || ex.Type == "direct" || ex.Type == "topic" || ex.Type == "fanout" || ex.Type == "header" {
		nameQueues = append(nameQueues, ex.Bind.SearchQueue(bindKey, ex.Type)...)
	} else {
		fmt.Println("Invalid Type of Exchange")
	}

	return nameQueues
}
