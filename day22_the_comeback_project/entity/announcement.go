package entity

import (
	"canvas_with_template/instructure"
	"fmt"
)

type Announcement instructure.Announcement_Info_Instructure

var AnnouncementList []Announcement

func (a *Announcement) String() string {
	res := ""
	res += fmt.Sprintf("%-15v: %v \n", "ID", a.ID)
	res += fmt.Sprintf("%-15v: %v \n", "Author", a.Author.DisplayName)
	res += fmt.Sprintf("%-15v: %v \n", "Content", a.Message)
	return res
}
