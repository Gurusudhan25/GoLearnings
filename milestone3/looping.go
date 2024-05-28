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
	var numberOfBoxes []int = []int{1, 2, 3, 4}
	numberOfBoxes = append(numberOfBoxes, 18)
	fmt.Printf("Now after appending the len of the array is %v and capacity is %v \n",
		len(numberOfBoxes),
		cap(numberOfBoxes),
	)
}
