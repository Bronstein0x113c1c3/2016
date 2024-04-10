package command

import (
	"canvas_with_template/entity"
	"fmt"
)

func Command() {
	fmt.Printf("\x1b[2J")
	fmt.Println("Choose your option: ")
	fmt.Println("1. Get user info")
	fmt.Println("2. Course/courses info, quizzes and related announcements")
	fmt.Println("3. All announcements")
	fmt.Println("4. Update the courses again")
	fmt.Println("5. Exit")
	var i int
	fmt.Scanln(&i)
	switch i {
	case 1:
		//user info
		fmt.Printf("\x1b[2J")
		userCommand()
		Command()
		// break
	case 2:
		//course/s info
		fmt.Printf("\x1b[2J")
		courseCommand()
		Command()
	case 3:
		fmt.Printf("\x1b[2J")
		announcementCommand()
		Command()
	case 4:
		//re-update again
		fmt.Printf("\x1b[2J")
		entity.GetCourseList()
		Command()
	default:
		fmt.Println("exiting")
		// return
		break
	}
}
