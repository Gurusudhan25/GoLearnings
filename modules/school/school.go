package main

import "fmt"

type Address struct {
	Street  string
	Door    string
	Country string
}

type School struct {
	Students   map[int8][]Student
	SchoolName string
	Address    Address
}

func (s *School) ShowStudentsOfClass(class int8) {
	fmt.Println(s.Students[class])
}

func (s *School) AddStudent(std Student) {
	if len(s.Students) == 0 {
		s.Students = make(map[int8][]Student)
		s.Students[std.Class] = []Student{std}
	} else {
		s.Students[std.Class] = append(s.Students[std.Class], std)
	}
}

func CreateSchool() *School {
	return &School{
		Address: Address{
			Street:  "trichy",
			Door:    "123",
			Country: "India",
		},
	}
}
