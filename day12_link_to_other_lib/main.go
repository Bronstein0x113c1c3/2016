package main

import "C"
import (
	"fmt"
	"reflect"
)

// //export MyStruct
// type MyStruct struct {
// 	Field1 int
// 	Field2 string
// }

//export TestS
func TestS() {
	type Student struct {
		Name string
		ID   int
	}
	s := &Student{
		Name: "Sdfsdfsdffrerwe",
		ID:   1,
	}
	fmt.Println(s)
}

//export CalculateFibonacci
func CalculateFibonacci(i int) int {
	if i <= 1 {
		return 1
	}
	return CalculateFibonacci(i-1) + CalculateFibonacci(i-2)
}

//export GetName
func GetName(s string) {
	fmt.Println(s)
}

//export GetStudentInfo
func GetStudentInfo() (string, int) {
	return "Monika", 1
}

//export R
func R(i interface{}) {
	fmt.Println(i)
}

func main() {
	fields := []reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(0),
		},
	}
	newType := reflect.StructOf(fields)
	newObject := reflect.New(newType).Elem()
	newObject.FieldByName("Name").SetString("K")
	newObject.FieldByName("Age").SetInt(1)
	fmt.Println(newObject)
}
