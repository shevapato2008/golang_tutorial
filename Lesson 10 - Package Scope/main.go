package main

import (
	"fmt"
)

var score = 99.5

func main() {

	sayHello("mario")
	// Output: Hello mario

	for _, v := range points {
		fmt.Print(v, " ")
	}
	fmt.Printf("\n")
	// Output: 20 90 100 45 70

	showScore()
	// Output: Your scored this many points: 99.5
}

// go run main.go greetings.go
