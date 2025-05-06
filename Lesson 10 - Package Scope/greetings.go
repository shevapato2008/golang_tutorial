package main

import (
	"fmt"
)

var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) {
	// This function is not exported
	fmt.Println("Hello", n)
}

func showScore() {
	// This function is exported
	fmt.Println("Your scored this many points:", score)
}
