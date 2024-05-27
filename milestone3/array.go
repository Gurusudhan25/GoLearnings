package milestone3

import "fmt"

func init() {
	SortArray()
}

func SortArray() {
	var age = [5]int{29, 23, 21, 24, 22}

	for i := 0; i < len(age); i++ {
		for j := i + 1; j < len(age); j++ {
			if age[i] > age[j] {
				age[i], age[j] = age[j], age[i]
			}
		}
	}

	fmt.Println(age)
}
