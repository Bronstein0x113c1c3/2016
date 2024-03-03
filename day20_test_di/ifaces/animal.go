package ifaces

type Animal interface {
	Speak() string
}

type Result struct {
	Animal Animal
}

func (r Result) Speak() string {
	return r.Animal.Speak()
}
