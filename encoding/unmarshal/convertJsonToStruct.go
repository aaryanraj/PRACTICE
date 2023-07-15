package main

import (
	"encoding/json"
	"fmt"
)

type Courses struct {
	CourseName string `json:"course-name"`
	CourseID   int    `json:"course-id"`
}

type User struct {
	Name    string  `json:"name"`
	EmpID   string  `json:"id"`
	Courses Courses `json:"courses"`
}

func main() {
	var u User

	jsonStr := `{"name":"Aaryan","id":"1787130","courses":{"course-name":"Kubernetes","course-id":12345}}`

	err := json.Unmarshal([]byte(jsonStr), &u)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u)
	}
}
