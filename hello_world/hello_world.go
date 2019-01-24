package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(a, b int) int {
	return a + b
}

func voidAdd(x, y int) {
	fmt.Printf("Void addition of %d and %d = %d \n", x, y, add(x, y))
}

func main() {
	var a int = 1
	var b int = 5
	var factor int = 2
	fmt.Println("Hello12, 世界")
	fmt.Println("My rand number is", rand.Intn(10))
	fmt.Printf("Square root %g\n", 1/math.Sqrt(2))
	fmt.Printf("Approximate Pi value is %g\n", math.Pi)
	fmt.Printf("Addition of %d and %d equals %d\n", a, b, add(a, b))
	voidAdd(a+factor, b+factor)
}
