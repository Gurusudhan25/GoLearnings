package milestone2

import (
	"errors"
	"fmt"
	"math"
)

// practicing functions
func SayMyAge() (int, int32, int64) {
	return FAVORITE_NUMBER, longInt, longLongInt
}

func AddNumber(num1 int, num2 int) (int, error) {
	var err error
	if num2 == 0 {
		err = errors.New("Error")
	}
	var add int = num1 + num2
	return add, err
}

func SayAllMaxValues() {

	// Int
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)

	// Uint
	fmt.Println(math.MaxUint8)
	fmt.Println(math.MaxUint16)
	fmt.Println(math.MaxUint32)
	// This actually overflows int
	// we need to explicty convert into uint or string
	fmt.Println(uint64(math.MaxUint64))

	// Float
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
}
