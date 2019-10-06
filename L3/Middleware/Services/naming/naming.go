package naming

import (
	"Middleware/Distribution/proxies"
)

type NamingService struct {
	Repository map[string]proxies.ClientProxy
}

func (naming *NamingService) Register(name string, proxy proxies.ClientProxy) (bool) {
	r := false

	// Verifica se já existe um repositório criado 
	if len(naming.Repository) == 0 {
		naming.Repository = make(map[string]proxies.ClientProxy)
	}
	// Verifica se já existe um serviço com esse nome registrado no repositório
	_, ok := naming.Repository[name]
	if ok {
		// Serviço já registrado 
		r = false 
	} else {
		naming.Repository[name] = proxies.ClientProxy{TypeName: proxy.TypeName, Host: proxy.Host, Port: proxy.Port}
		r = true
	}

	return r
}

func (naming NamingService) Lookup(name string) proxies.ClientProxy {

	return naming.Repository[name]
}

func (naming NamingService) List(name string) map[string]proxies.ClientProxy {

	return naming.Repository
}


