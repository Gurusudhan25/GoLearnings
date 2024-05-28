package main

import (
	"errors"
	"fmt"

	"github.com/gurusudhan25/milestone2"
	"github.com/gurusudhan25/milestone3"
)

func main() {

	fmt.Println("Hello World!")

	var age = [5]int{29, 23, 21, 24, 22}
	milestone3.SortArray(age)
}

// Unit testing function
func Hello(name string) (string, error) {
	var err error
	if len(name) == 2 {
		err = errors.New("short name")
	}
	return name, err
}

func ErrorHandling() {
	var number, err = milestone2.AddNumber(10, 0)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(number)
}
