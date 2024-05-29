package milestone2

import "fmt"

// declaration
// var
var age int = 24
var year int = 2024

// Const
const FAVORITE_NUMBER int = 25

// Int other types
// var someNumber int8 = 128 -> Error int8 -128 to 127
var smallInt int8 = 127

// var medium32768Int int16 = 32768 -> Error int8 -32768 to 32768
var mediumInt int16 = 32767

// var longInt int32 = 2147483648 -> Error int8 -2147483648 to 2147483647
var longInt int32 = 2147483647

// var longInt int32 = 9223372036854775808 -> Error int8 -9223372036854775808 to 9223372036854775807
var longLongInt int64 = 9223372036854775807

// Uint other types - Is only for positive intgers
// var someNumber int8 = 256 -> Error int8 upto 255
var smallUint uint8 = 255

// var medium32768uint uint16 = 65536 -> Error uint8 upto 65535
var mediumUint uint16 = 65535

// var longuint uint32 = 4294967296 -> Error uint8 upto 4294967295
var longUint uint32 = 4294967295

// var longuint uint32 = 18446744073709551616 -> Error uint8 upto 18446744073709551615
var longLongUint uint64 = 18446744073709551615

// Float
// Float32 and Float64
var smallFloat float32 = 3.4028234663852886e+38
var longFloat float64 = 1.7976931348623157e+308

/*
	  init() will be triggered before the main() function,
		and follow this order of execution: init() will run only once.
		init() will run after global variable initialization of each package and before main().
		init() will only run if the package is imported
*/
func init() {
	fmt.Println("Guru from init")
	fmt.Println(smallFloat)
	fmt.Println(longFloat)

	myName := "Gurusudhan"
	fmt.Println(myName)
}
