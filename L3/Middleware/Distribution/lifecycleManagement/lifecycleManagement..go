package LifecycleManagement

import (
	"App/impl"
	"Middleware/Services"
)

type LifecycleMan struct {
	objects []*impl.DataBank
}

func (lcm *LifecycleMan) registerObject(newObject impl.DataBank) {
	
	append(lcm.objects, newObject)
}
