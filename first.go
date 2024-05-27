package main

import (
	"fmt"

	"github.com/gurusudhan25/milestone2"
	"github.com/gurusudhan25/milestone3"
)

func main() {
	const name string = "Gurusudhan"
	const age int = 24
	var number, err = milestone2.AddNumber(10, 0)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(number)
	fmt.Println(name)
	fmt.Println(age)

	fmt.Println(milestone2.SayMyAge())
	milestone2.SayAllMaxValues()

	fmt.Println(milestone3.CheckAboveAge(17))
	fmt.Println(milestone3.CheckAboveAge(19))
	fmt.Println(milestone3.CheckAboveAge(18))

	// Looping
	milestone3.PrintNumbers()
}

// gopls -> Language server protocal auto detection and remove unused
// gofmt -> Go formatting files
// goleak -> Will find go routine leaks
// go.mod -> root of your Go
// go.work -> managing multi-module workspaces
