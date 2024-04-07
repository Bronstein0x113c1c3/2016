package common

const (
	StudentInfoEndpoint      = "https://troy.instructure.com/api/v1/users/self"
	CourseInfoEndpoint       = "https://troy.instructure.com/api/v1/courses"
	QuizInfoEndpoint         = "https://troy.instructure.com/api/v1/courses/%v/quizzes/"
	AnnouncementInfoEndpoint = "https://troy.instructure.com/api/v1/announcements?context_codes[]=course_%v"
)
