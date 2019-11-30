package exchange

import (
	"Concurrent-Distributed-Systems/NMiddleware/Middleware/Distribution/bind"
	"fmt"
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
	// OBS: TEM QUE AJEITAR LOGO AS INFORMAÇÕES DO PACOTE E DAS MENSAGENS
	// TO USANDO OQ TÁ, MAS TEM QUE MUDAR DPS I GUESS
	if ex.Type == "direct" || ex.Type == "" {
		nameQueues = append(nameQueues, ex.Bind.SearchQueue(bindKey, ex.Type)...) 

	} else if ex.Type == "topic" {
		nameQueues = append(nameQueues, ex.Bind.SearchQueue(bindKey, ex.Type)...)
	} else {
		fmt.Println("Invalid Type of Exchange")
	}

	return nameQueues
}
