package proxy

import (
	"Middleware/Distribution/proxies"
	"Middleware/Distribution/requestor"
	"Middleware/aux"
	"Middleware/repository"
	"shared"
)

type NamingProxy struct{}

func (NamingProxy) Register(p1 string, proxy interface{}) bool {

	// prepare invocation
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = proxy
	namingproxy := proxies.ClientProxy{Host: "", Port: shared.NAMING_PORT, Id: 0}
	request := aux.Request{Op: "Register", Params: params}
	inv := aux.Invocation{Host: namingproxy.Host, Port: namingproxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	return ter[0].(bool)
}

func (NamingProxy) Lookup(p1 string) interface{} {
	// prepare invocation
	params := make([]interface{}, 1)
	params[0] = p1
	namingproxy := proxies.ClientProxy{Host: "", Port: shared.NAMING_PORT, Id: 0}
	request := aux.Request{Op: "Lookup", Params: params}
	inv := aux.Invocation{Host: namingproxy.Host, Port: namingproxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	// process reply
	proxyTemp := ter[0].(map[string]interface{})
	clientProxyTemp := proxies.ClientProxy{TypeName: proxyTemp["TypeName"].(string), Host: proxyTemp["Host"].(string), Port: int(proxyTemp["Port"].(float64))}
	clientProxy := repository.CheckRepository(clientProxyTemp)

	return clientProxy
}
