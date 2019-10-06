package aux

type Invocation struct {
	Host string
	Port int
	Request Request
}

type Termination struct {
	Rep Reply
}

type AOR struct {
	Host string
	Port int
	Id   int
}

type Request struct {
	Op     string
	Params []interface{}
}

type Reply struct {
	Result []interface{};
}

