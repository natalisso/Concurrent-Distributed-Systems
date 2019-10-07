package repository

import (
	"Middleware/Distribution/proxies"
	"reflect"
)

func CheckRepository(proxy proxies.ClientProxy) interface{} {
	var clientProxy interface{}

	switch proxy.TypeName {
	case reflect.TypeOf(proxies.DataBaseProxy{}).String():
		dataBaseProxy := proxies.NewDataBaseProxy()
		dataBaseProxy.Proxy.TypeName = proxy.TypeName
		dataBaseProxy.Proxy.Host = proxy.Host
		dataBaseProxy.Proxy.Port = proxy.Port
		clientProxy = dataBaseProxy
	}

	return clientProxy
}
