package bind

import (
	"fmt"
	"regexp"
	"strings"
)

type Bind struct {
	bindKeys map[string][]string
}

func NewBind() Bind {
	bind := new(Bind)
	bind.bindKeys = make(map[string][]string)
	return *bind
}

func (bind *Bind) SearchQueue(bindKey string, typeExchange string) []string {
	matchs := make([]string, 0)
	cont := true
	for queue, bkQueue := range bind.bindKeys {
		if typeExchange == "topic" || typeExchange == "" {
			found := false
			for i := 0; i < len(bkQueue); i++ {
				if strings.ContainsAny(bkQueue[i], "*") {
					contains := true
					_strs := strings.Split(bkQueue[i], "*")
					for j := 0; j < len(_strs); j++ {
						if !strings.Contains(bindKey, _strs[j]) {
							fmt.Println("deu ruim!!")
							contains = false
							break
						}
					}

					if contains {
						matchs = append(matchs, queue)
						found = true
						fmt.Println("Matched!", "matchs:", matchs)
					}
				} else {
					var result string
					if strings.HasSuffix(bkQueue[i], "#") {
						result = strings.Replace(bkQueue[i], "#", "*", -1)
					} else {
						result = bkQueue[i]
					}

					matched, err := regexp.MatchString(result, bindKey)
					if err == nil {
						if matched {
							matchs = append(matchs, queue)
							found = true
							fmt.Println("Matched:", matched, "matchs:", matchs)
						} else {
							fmt.Println("Matched:", matched, "DEU RUIM")
						}
					}
				}

				if found {
					fmt.Println("Next queue!")
					break
				}
			}
		} else if typeExchange == "direct" {
			for i := 0; i < len(bkQueue); i++ {
				if bkQueue[i] == bindKey {
					matchs = append(matchs, queue)
					cont = false
				}
				break
			}
			if !cont {
				break
			}
		} else if typeExchange == "fanout" {
			matchs = append(matchs, queue)
		} else if typeExchange == "header" {

		}
	}
	return matchs
}

// MÉTODO PARA ADICIONAR UM  BINDKEY  A UMA FILA
func (bind *Bind) BindQueue(queueName string, bindKey string) {
	bind.bindKeys[queueName] = append(bind.bindKeys[queueName], bindKey)
}
