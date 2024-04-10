package main

import (
	"canvas_with_template/common"
	// "reflect"
	// "fmt"
)

func main() {
	//load the token
	token, err := common.LoadToken("env/.env")
	if err != nil {
		panic(err)
	}

}
