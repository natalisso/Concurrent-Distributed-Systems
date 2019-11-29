package bind

import "fmt"

// Bind ...
type Bind struct {
	bindKeys map[string][]string
	// Mapeia o nome da fila para um slice que tem as strings que são os seus bind key, dessa fila.
	// Pode ter mais de um bind key? deixei assim por enquanto!!!
}

func NewBind() Bind {
	bind := new(Bind)
	bind.bindKeys = make(map[string][]string)
	return *bind
}

// Chegou uma mensagem do Produtor e o exchange é Topic
// Vou procurar as filas que possam dar MATCH com o seu Routine Key
// E retornar um slice c/ os nomes de todas as filas que deram match
func (bind *Bind) SearchQueue(bindKey string) []string {
	matchs := make([]string, 0)
	for queue, bkQueue := range bind.bindKeys {
		for i := 0; i < len(bkQueue); i++ {
			// Aqui irá ser feito a comparação de string e todo trabalho chato
			// Para verificar se der match, utiliza bkQueue[i]; Caso dê
			// faz matchs = append(matchs, bkQueue[i]) e da um break, pra ir p proxima queue
			fmt.Println("%s", queue)
		}
	}
	return matchs
}
