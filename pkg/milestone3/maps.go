package milestone3

import "fmt"

func mapPractice() {

	var student map[string]string = make(map[string]string)
	var student2 = map[int]string{
		1: "Hello",
		2: "World",
	}

	student["name"] = "Gurusudhan"
	student["city"] = "chennai"
	student["age"] = "22"
	fmt.Println(student)
	fmt.Println(student2)

	for i, v := range student {
		println(i, v)
	}
}
