package lifecycleManagement

import (
	"App/impl"
	"shared"
)

type LifecycleMan struct {
	pool []impl.DataBase
}

func (lcm *LifecycleMan) NewLifecycleMan() []impl.DataBase {

	var db []impl.DataBase
	for i := 0; i < shared.N_INSTANCES; i++ {
		dataBaseImpl := impl.DataBase{Id: i}
		lcm.pool = append(lcm.pool, dataBaseImpl) // registra no life cycle management
		db = append(db, dataBaseImpl)
	}
	return db
}
