package structfile

import "C"
import "fmt"

//export Student
type Student struct {
	Name string
	ID   int
}

//export New
func New() {
	fmt.Println("sdfsdfdsfs")
}
