package milestone3

import "fmt"

func init() {
	moreForLoops()
	slice()
	mapPractice()
}

func SortArray(age [5]int) {

	for i := 0; i < len(age); i++ {
		for j := i + 1; j < len(age); j++ {
			if age[i] > age[j] {
				age[i], age[j] = age[j], age[i]
			}
		}
	}

	fmt.Println(age)
}
