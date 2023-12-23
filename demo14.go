package main

import (
	"fmt"
)

type People struct {
	name string
	age  int
}

type Student struct {
	People
	class  string
	source float32
}

func (s *Student) printStudent(a int) {
	fmt.Printf("%v %v", s, a)
}

func main14() {
	var student Student
	student.name = "dbp"
	student.class = "软件2202"
	student.age = 19
	student.source = 100

	updateStudentName(&student)
	student.printStudent(111)
}

func updateStudentName(student *Student) {
	student.name = "sb"
}
