# **GoLearnings**
<!-- --- -->
<!-- ![Preview](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/2560px-Go_Logo_Blue.svg.png) -->
<!-- --- -->

## `Milestone1`

### Why we use GO
---
 - Statically typed and strongly typed
 - Compiled language produces exe
 - Concurrency
 - lightweight threads, called goroutines
 - Simplicity
 - More GO libraries

### Real-life applications where Go is extensively used
---
 - Web Servers and APIs
 - Microservices
 - Command-Line Tools (CLIs)
 - Cloud Native Applications
 - Network Programming


## `Milestone 2`

### GO Tools
---
- **Gopls**  - Language server protocal auto detection and remove unused
- **Gofmt**  - Go formatting files
- **Goleak** - Will find go routine leaks
- **Gomod**  - root of your Go
- **Gowork** -  managing multi-module workspaces

### Main() vs init()
---
 - init() will be triggered before the main function
- and follow this order of execution: init will run only once
- init() will run after global variable initialization of each package and before main()
- init() will only run if the package is imported

### Int and max values
---

```
var smallInt int8 = 127
// var someNumber int8 = 128 -> Error int8 -128 to 127
```
```
var mediumInt int16 = 32767
// var medium32768Int int16 = 32768 -> Error int8 -32768 to 32768
```
```
var longInt int32 = 2147483647
// var longInt int32 = 2147483648 -> Error int8 -2147483648 to 2147483647
```
```
var longLongInt int64 = 9223372036854775807
// var longInt int32 = 9223372036854775808 -> Error int8 -9223372036854775808 to 9223372036854775807
```

### Uint and max values
---
```
var smallUint uint8 = 255
// var someNumber uint8 = 256 -> Error int8 upto 255
```
```
// var medium32768uint uint16 = 65536 -> Error uint8 upto 65535
var mediumUint uint16 = 65535
```
```
// var longuint uint32 = 4294967296 -> Error uint8 upto 4294967295
var longUint uint32 = 4294967295
```
```
// var longuint uint32 = 18446744073709551616 -> Error uint8 upto 18446744073709551615
var longLongUint uint64 = 18446744073709551615
```

### Float
---

```
var smallFloat float32 = 3.4028234663852886e+38
var longFloat float64 = 1.7976931348623157e+308
```

### Bool
---

```
var isHot bool = false
var CanEat bool = true
```

### Exporting and Importing
---
The primary rule is that an identifier's first letter must be uppercase to be considered exported. Names starting with lowercase letters are not exported and are restricted to the package where they are defined

```
// Exported names (accessible from other packages)
const Pi = 3.14159
func GetMessage() string {
  return "Hello, world!"
}

// Not exported (only accessible within this package)
var secretNumber = 42
func internalFunction() {
  fmt.Println("This is internal")
}
```

### Rune and String
---

| Rune |  String  |
|:-----|:--------|
| Rune represents a single unicode character   | String is sequence of bytes |
| A rune represents a single character   |  string can represent a sequence of characters  |
| Runes have a fixed-width integer representation  | strings use a variable-length encoding scheme like UTF-8 |
| `int32`  | `[]byte` |

### Print

Printf - Formatted printing
```
fmt.Printf("My name is %s and age is %v \n",
  "Gurusudhan",
  25,
)
```
Println - Print new lines
```
fmt.Println("Hello")
fmt.Println("World")

// Output
Hello
World
```
Print - Print in same line
```
fmt.Println("Hello")
fmt.Println("World")

// Output
Hello World
```

