package proxies

import (
	"Middleware/Distribution/requestor"
	"Middleware/aux"
	"fmt"
	"math/rand"
	"reflect"
	"shared"
)

type DataBaseProxy struct {
	Proxy ClientProxy
}

func NewDataBaseProxy() DataBaseProxy {
	p := new(DataBaseProxy)

	p.Proxy.TypeName = reflect.TypeOf(DataBaseProxy{}).String()
	p.Proxy.Host = "localhost"
	//p.Proxy.Port = shared.FindNextAvailablePort()  // TODO
	p.Proxy.Port = shared.CALCULATOR_PORT
	p.Proxy.Id = rand.Intn(shared.N_INSTANCES)
	return *p
}

func (proxy DataBaseProxy) Save(p1 string, p2 string, p3 int) bool {

	// prepare invocation
	params := make([]interface{}, 3)
	params[0] = p1
	params[1] = p2
	params[2] = p3
	request := aux.Request{Op: "Save", Params: params}
	inv := aux.Invocation{Host: proxy.Proxy.Host, Port: proxy.Proxy.Port, Request: request, Id: proxy.Proxy.Id}

	// invoke requestor
	req := requestor.Requestor{}
	fmt.Println("Vou pro invoke")
	result := req.Invoke(inv).([]interface{})

	return result[0].(bool)
}

func (proxy DataBaseProxy) Search(p1 string) bool {

	// prepare invocation
	params := make([]interface{}, 1)
	params[0] = p1
	request := aux.Request{Op: "Search", Params: params}
	inv := aux.Invocation{Host: proxy.Proxy.Host, Port: proxy.Proxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	result := req.Invoke(inv).([]interface{})

	return result[0].(bool)
}
