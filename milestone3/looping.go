package milestone3

import "fmt"

func PrintNumbers() {
	for i := 0; i < 11; i++ {
		fmt.Println((i))
	}
}

func moreForLoops() {
	i := 0
	for {
		if i > 10 {
			break
		}

		fmt.Println(i)
		i++
	}
}

func slice() {
	var numberOfBoxes []int = []int{100, 200, 300, 400}
	numberOfBoxes = append(numberOfBoxes, 18)

	var example2 [2]uint8 = [2]uint8{1, 2}

	example2[1] = 12
	// example2[1] = -12 error only posite numbers allowed
	// example2[2] = 12 invalid argument: index 2 out of bounds [0:2]

	fmt.Printf("Now after appending the len of the array is %v and capacity is %v \n",
		len(numberOfBoxes),
		cap(numberOfBoxes),
	)

	for i := range numberOfBoxes {
		fmt.Println(numberOfBoxes[i])
	}
}
