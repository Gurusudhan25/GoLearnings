package main

import (
	"fmt"

	"github.com/gurusudhan25/pkg/milestone2"
	"github.com/gurusudhan25/pkg/milestone3"
	"github.com/gurusudhan25/pkg/milestone4"
)

func main() {

	fmt.Println("Hello World!")

	var age = [5]int{29, 23, 21, 24, 22}
	milestone3.SortArray(age)
	milestone4.Milestone4()
}

func ErrorHandling() {
	var number, err = milestone2.AddNumber(10, 0)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(number)
}
