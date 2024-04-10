package entity

import (
	"canvas_with_template/common"
	"canvas_with_template/instructure"
	"fmt"
)

type Course instructure.Course_Info_Instructure

var CourseList []Course

func GetCourseList(token string) {
	if err := common.Get(common.CourseInfoEndpoint, token, &CourseList); err != nil {
		panic(fmt.Sprintf("cannot get the courses: %v", err))
	}
	// return courseList
}
func (c Course) String() string {
	return fmt.Sprintf("%-10v|%-51v|%-50v|", c.ID, c.Name, c.CourseCode)
}

func (c *Course) GetAnnouncements() []Announcement {
	
}
func (c *Course) GetQuizzes() []Quiz {

}	

// func (c )
