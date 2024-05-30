package milestone4

import "fmt"

type Student struct {
	Name string
	Age  int
}

type School struct {
	Students []Student
	Place    string
}

func (s *School) ShowStudent() {
	for i := range s.Students {
		fmt.Println(s.Students[i].Name)
	}
}

func (s *School) AddStudentToSchool(student Student) {
	s.Students = append(s.Students, student)
}

func init() {
	student1 := Student{
		"Gurusudhan", 22,
	}
	student2 := Student{
		"Rajesh", 28,
	}
	school1 := School{}

	fmt.Println(school1)
	school1.AddStudentToSchool(student1)
	school1.AddStudentToSchool(student2)
	school1.ShowStudent()
}
