package main

import (
	"fmt"
)

func updateName(x string) {
	x = "wedge"
}

func updateNameV2(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 4.99
}

func main() {
	// group A types [Non-Pointer Values] -> strings, ints, bools, floats, arrays, structs
	name := "tifa"

	updateName(name)

	fmt.Println(name)
	// output: tifa, not wedge because we passed a copy of the value

	name = updateNameV2(name)
	fmt.Println(name)
	// output: wedge, because we returned the updated value

	// group B types [Pointer Wrapper Values] -> slices, maps, functions
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}
	updateMenu(menu)
	fmt.Println(menu)
	// output: map[ice cream:3.99 pie:5.95 coffee:4.99]
	// the map is updated because we passed a reference to the map, not a copy
	// maps are reference types in Go, so they are passed by reference
	// this means that when we update the map, the changes are reflected outside the function

}
