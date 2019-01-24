package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i int64

	if 2 > len(os.Args) || len(os.Args) > 2 {
		fmt.Println("arguments limit exceeded")
		os.Exit(0)
	}

	limit, err := strconv.ParseInt(os.Args[1], 10, 0)

	if err != nil {
		fmt.Println("Input argument must be an integer")
		os.Exit(1)
	}

	fmt.Printf("Entered argument value is %d \n", limit)

	for i = 0; i <= limit; i++ {
		fmt.Printf("Iteration %d = %d \n", i, iterator(i))
	}
	fmt.Println("")
}

func iterator(j int64) int64 {
	if j <= 1 {
		return j
	}

	return iterator(j-1) + iterator(j-2)
}
