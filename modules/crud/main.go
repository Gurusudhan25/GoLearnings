package main

import (
	"fmt"
	"sync"
)

func Reader(wg *sync.WaitGroup) {
	fmt.Println("Hello World")
	wg.Done()
}

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go reader(&wg)
// 	wg.Wait()

// 	user1 := createUser1()
// 	user2 := createUser2()
// 	// Pointers
// 	user1.ShowAge()
// 	user1.ShowAges()
// 	user1.AddAges(2)
// 	user1.ShowAge()
// 	user1.ShowAges()
// 	user1.AddAges(2)
// 	fmt.Println("he")
// 	user2.ShowAge()
// 	user2.ShowAges()
// 	user2.AddAges(2)
// 	user2.ShowAge()
// 	user2.ShowAges()
// 	user2.AddAges(2)

// 	age := 30

// 	// Declare a pointer variable that stores the memory address of 'age'
// 	var agePtr *int = &age // '&' address-of operator

// 	fmt.Println("Value of age:", age)           // Prints 30
// 	fmt.Println("Memory address of age:", &age) // Prints memory address of 'age'
// 	fmt.Println("Value of agePtr:", agePtr)     // Prints memory address of 'age' (stored in agePtr)

//		// Dereference the pointer to access the actual value
//		*agePtr = 42 // Modifies the value at the memory location pointed to by agePtr (which is 'age')
//		fmt.Println("Value of age after modification:", age)
//		fmt.Println("Value of agePtr:", agePtr)
//		*agePtr = 100
//		fmt.Println("Value of age after modification:", age)
//	}
//
// type myInt
type Mine interface {
	checkNumber() bool
}

func doubleValue(num Mine) {
	// number *= 2 // Dereference the pointer to modify the value at the memory address
	num.checkNumber()
	// if int(num) {
	// }
}

type Num struct {
	val int
}

func (n *Num) checkNumber() bool {
	return true
}

func main() {
	var value Mine = &Num{10}
	doubleValue(value)                   // Pass the address of 'value' using '&'
	fmt.Println("Doubled value:", value) // Prints 20 (value has been modified)
}
