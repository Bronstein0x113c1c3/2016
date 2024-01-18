package main

import "C"
import (
	"fmt"
	"log"
	"net"
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

//export BeAServer
func BeAServer(port int) {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return
	}
	log.Println("I'm waiting for someone....")
	for {
		newConn, err := conn.Accept()
		if err != nil {
			panic(err)
		}
		go func(c net.Conn) {
			for {
				c.Write([]byte("Bonjour!"))
			}
		}(newConn)
	}

}

func main() {

}
