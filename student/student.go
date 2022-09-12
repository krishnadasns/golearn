package student

import "github.com/krishnadasns/golearn/school"

type Student struct {
	Name    string
	Age     int
	Section string
	School  school.School
}

func New(name string, age int, sec string, school school.School) Student {
	return (Student{Name: name, Age: age, Section: sec, School: school})
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) GetAge() int {
	return s.Age
}

func (s *Student) GetSection() string {
	return s.Section
}
