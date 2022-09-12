package main

import (
	"fmt"

	"github.com/krishnadasns/golearn/school"
	"github.com/krishnadasns/golearn/student"
)

func main() {
	fmt.Printf("Entering Main Function\n")
	s1 := student.New("Das", 26, "X", school.New("LMPS", "Vaikom"))

	fmt.Printf("%s\n", s1.GetName())
	fmt.Printf("%d\n", s1.GetAge())
	fmt.Printf("%s\n", s1.GetSection())
	fmt.Printf("%s\n", s1.School.GetSchoolName())
	fmt.Printf("%s\n", s1.School.GetSchoolAddress())

}
