package ifaces

import "fmt"

type Dog struct {
	Name  string
	Breed string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("Dog name: %v, breed: %v", d.Name, d.Breed)
}

type Cat struct {
	Name string
	Age  int
}

func (c Cat) Speak() string {
	return fmt.Sprintf("Cat name: %v, age: %v", c.Name, c.Age)
}
