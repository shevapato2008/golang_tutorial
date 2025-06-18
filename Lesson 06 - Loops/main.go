package main

import (
	"fmt"
)

func main() {
	x := 0

	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}
	/*
		value of x is: 0
		value of x is: 1
		value of x is: 2
		value of x is: 3
		value of x is: 4
	*/

	names := []string{"mario", "luigi", "yoshi", "peach"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	/*
		mario
		luigi
		yoshi
		peach
	*/

	for index, value := range names {
		fmt.Printf("index: %v, value is %v\n", index, value)
	}
	/*
		index: 0, value is mario
		index: 1, value is luigi
		index: 2, value is yoshi
		index: 3, value is peach
	*/

	for _, value := range names {
		fmt.Printf("value is %v\n", value)
		// this line will not change the original value in the slice,
		// because value is a copy
		value = "new value"
	}
	fmt.Println(names) // original slice remains unchanged
	/*
		value is mario
		value is luigi
		value is yoshi
		value is peach
		[mario luigi yoshi peach]
	*/
}
