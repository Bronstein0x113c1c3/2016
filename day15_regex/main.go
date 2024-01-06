package main

import (
	"fmt"
	"regexp"
)

func main() {
	// pattern, compileErr := regexp.Compile(`\w+oat`)
	// description := "A boat for one person"
	// question := "Is that a goat?"
	// preference := "I like oats"
	// if compileErr == nil {
	// 	fmt.Println("Description:", pattern.MatchString(description))
	// 	fmt.Println("Question:", pattern.MatchString(question))
	// 	fmt.Println("Preference:", pattern.MatchString(preference))
	// } else {
	// 	fmt.Println("Error:", compileErr)
	// }
	// test_string := "boatsfsfd blahblahboat df goat xf boat"
	// fmt.Println(pattern.FindAllStringIndex(test_string, -1))
	// fmt.Println(pattern.Split(test_string, -1))
	pattern := regexp.MustCompile("A ([A-Za-z]*) for ([A-z]*) person")
	desc := "Kayak. A boat for one person. A kayak for two person. "
	subs := pattern.FindAllStringSubmatch(desc, -1)
	for idx1, s := range subs {
		for idx2, j := range s {
			fmt.Printf("%v - %v - %v \n", idx1, idx2, j)
		}
	}

}
