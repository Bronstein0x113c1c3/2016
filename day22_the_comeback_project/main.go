package main

import (
	"canvas_with_template/common"
	"canvas_with_template/instructure"
	"fmt"
	"reflect"
	// "reflect"
	// "fmt"
)

func main() {
	//load the token
	token, err := common.LoadToken("env/.env")
	if err != nil {
		panic(err)
	}
	var c []instructure.Course_Info_Instructure
	err = common.Get(common.CourseInfoEndpoint, token, &c)
	if err != nil {
		panic(err)
	}
	x := c[0].CreatedAt
	fmt.Println(x)
	// common.Beautify(c[0], 0)
	common.BeautifyReflection(reflect.ValueOf(c[0]), 0)
}
