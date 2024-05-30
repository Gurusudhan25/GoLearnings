package milestone4

import "fmt"

// Basic struct
type Vertex struct {
	X int
	Y int
}

type User struct {
	Name string
	Age  int
	City string
}

func init() {
	var str Vertex = Vertex{1, 2}
	// Address of str.X and str are same
	a := fmt.Sprintf("%p, %p, %p", &str.X, &str.Y, &str)
	fmt.Println(a)
	arrayOfStruct()
}

func arrayOfStruct() {
	// Creating new student
	var User1 User = User{
		"Guru",
		22,
		"Chennai",
	}

	var User2 User = User{
		"sudhan",
		21,
		"Chennai",
	}

	var students []User = []User{User1, User2}

	for i := range students {
		fmt.Println(students[i].Name)
	}
}
