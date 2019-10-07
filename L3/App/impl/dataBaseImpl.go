package impl

import "shared"

type user struct {
	name     string
	identity string
	year     int
}

type DataBase struct {
	Id    int
	users []user
}

func (bank *DataBase) InvokeDataBase(req shared.Request) bool {
	var r bool

	op := req.Op
	p1 := req.P1
	p2 := req.P2
	p3 := req.P3

	switch op {
	case "save":
		r = bank.Save(p1, p2, p3)
	case "search":
		r = bank.Search(p2)
	}
	return r
}

func (bank *DataBase) Save(name string, identity string, year int) bool {
	newUser := user{name: name, identity: identity, year: year}
	bank.users = append(bank.users, newUser)

	return true
}

func (bank *DataBase) Search(ind string) bool {

	found := false
	for i := 0; i < len(bank.users); i++ {
		if bank.users[i].identity == ind {
			found = true
		}
	}

	return found
}
