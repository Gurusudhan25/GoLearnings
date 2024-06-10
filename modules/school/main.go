package main

import "fmt"

func main() {
	school := CreateSchool()
	fmt.Println(*school)
	// school := School{}
	// Gonna get this in interface
	std1 := Student{
		Name: "gurusudhan",
		Age:  21,
		Address: Address{
			Street:  "Chennai",
			Door:    "1/223",
			Country: "India",
		},
		Class: 12,
	}
	// std2 := Student{
	// 	Name: "sudhann",
	// 	Age:  21,
	// 	Address: Address{
	// 		Street:  "Chennai",
	// 		Door:    "1/223",
	// 		Country: "India",
	// 	},
	// 	Class: 12,
	// }
	school.AddStudent(std1)
	// school.AddStudent(std2)
	fmt.Printf("%x", *school)

}
