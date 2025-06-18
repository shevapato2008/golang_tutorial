package main

import (
	"fmt"
)

func main() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)
	/*
		true
		false
		false
		true
	*/

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is 40 or older")
	}
	/*
		Age is 40 or older
	*/

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		fmt.Printf("index: %v, value: %v\n", index, value)
	}
	/*
		index: 0, value: mario
		continuing at pos 1
		index: 2, value: yoshi
		index: 3, value: peach
		index: 4, value: bowser
	*/

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("index: %v, value: %v\n", index, value)
	}
	/*
		index: 0, value: mario
		continuing at pos 1
		index: 2, value: yoshi
		breaking at pos 3
	*/
}
