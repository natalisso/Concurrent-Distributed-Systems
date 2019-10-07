package proxies

import (
	"Middleware/Distribution/requestor"
	"Middleware/aux"
	"reflect"
	"shared"
	"math/rand"
)

const N_OBJECTS = 10

type DataBankProxy struct {
	Proxy ClientProxy
}

func NewDataBankProxy() DataBankProxy {
	p := new(DataBankProxy)

	p.Proxy.TypeName = reflect.TypeOf(DataBankProxy{}).String()
	p.Proxy.Host = "localhost"
	//p.Proxy.Port = shared.FindNextAvailablePort()  // TODO
	p.Proxy.Port = shared.CALCULATOR_PORT
	p.Proxy.id = rand.intn(N_OBJECTS)
	return *p
}

func (proxy DataBankProxy) Save(p1 string, p2 int, p3 string) string {

	// prepare invocation
	params := make([]interface{}, 3)
	params[0] = p1
	params[1] = p2
	params[2] = p3
	request := aux.Request{Op: "Save", Params: params}
	inv := aux.Invocation{Host: proxy.Proxy.Host, Port: proxy.Proxy.Port, Request: request, id: proxy.Proxy.Id}

	// invoke requestor
	req := requestor.Requestor{}
	result := req.Invoke(inv).([]interface{})

	return string(result[0].(string))
}

func (proxy DataBankProxy) Search(p1 string) bool {

	// prepare invocation
	params := make([]interface{}, 1)
	params[0] = p1
	request := aux.Request{Op: "Search", Params: params}
	inv := aux.Invocation{Host: proxy.Proxy.Host, Port: proxy.Proxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	result := req.Invoke(inv).([]interface{})

	return bool(result[0].(float64))
}
