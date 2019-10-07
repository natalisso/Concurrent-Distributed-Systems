package impl

import "shared"

type user struct {
	name     string
	identity string
	age      int
}

type DataBank struct {
	id    int
	users []user
}

func (bank *DataBank) InvocaCalculadora(req shared.Request) []interface{} {
	var saved string
	var wasFound bool
	var userFounded user

	op := req.Op
	p1 := req.P1
	p2 := req.P2
	p3 := req.p3

	switch op {
	case "save":
		saved = bank.Save(p1, p2, p3)
	case "search":
		r = bank.Search(p2)
	}
	return r
}

func (bank *DataBank) Save(name string, identity string, age int) string {
	newUser := user{name: name, identity: identity, age: age}
	append(bank.users, newUser)

	return "User saved successfully"
}

func (bank *DataBank) Search(ind string) bool {

	found := false
	for i := 0; i < len(bank.users); i++ {
		if bank.users[i].identity == ind {
			found = true
		}
	}

	return found
}
