package main

type Student struct {
	Name string
	ID   int
	GPA  float32
}

func checkGPA(a float32) string {
	if a >= 3.6 {
		return "Excellent"
	} else if a >= 3 {
		return "Good"
	} else if a >= 2 {
		return "Fair"
	} else {
		return "Bad"
	}
}
