package milestone4

import "fmt"

// Basic struct
type Vertex struct {
	X int
	Y int
}

func init() {
	var str Vertex = Vertex{1, 2}
	a := fmt.Sprintf("%p, %p, %p", &str.X, &str.Y, &str)
	fmt.Println(a)
}
