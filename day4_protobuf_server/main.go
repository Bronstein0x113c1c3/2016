package main

import (
	"day4/person"
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	p1 := &person.Person{
		Name:       "Bronstein",
		Occupation: "Dev",
		Email: []*person.Email{
			&person.Email{
				Kind:    "Work",
				Address: "blahblah@gsaddsf.com",
			},
			&person.Email{
				Kind:    "Play",
				Address: "Bronstein3211",
			},
		},
	}
	data, err := proto.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data))
	// p2 := new(person.Person)
	// err = proto.Unmarshal(data, p2)
	// fmt.Println(p2)

}
