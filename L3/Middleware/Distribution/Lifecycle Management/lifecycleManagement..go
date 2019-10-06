package LifecycleManagement

import (
	"App/impl"
	"Middleware/Services"
)

type LifecycleMan struct {
	tasks []*impl.DataBank
}

func (lcm *LifecycleMan) createObject() {

	dataBankImpl := impl.DataBank{}
	append(lcm.tasks, newObj)
}
