package main

import (
	"errors"
	"fmt"
)

func main() {
	const name string = "Gurusudhan"
	const age int = 24
	var number, err = addNumber(10, 0)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(number)
}

func addNumber(num1 int, num2 int) (int, error) {
	var err error
	if num2 == 0 {
		err = errors.New("Error")
	}
	var add int = num1 + num2
	return add, err
}
