package main

import (
	"fmt"
)

func updateName(x string) {
	x = "wedge"
}

func updateNameByPointer(x *string) {
	*x = "wedge"
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs
	name := "tifa"

	updateName(name)

	fmt.Println("memory address of name:", &name) // printing the memory address of name
	// output: memory address of name: 0xc000012060
	// this is dynamic memory address, it will change every time you run the program

	m := &name                                  // m is a pointer to name
	fmt.Println("memory address:", m)           // printing the memory address of name
	fmt.Println("value at memory address:", *m) // dereferencing m to get the value at that memory address
	fmt.Println("memory address of m:", &m)     // printing the memory address of m
	// output: memory address: 0xc000012060
	// output: value at memory address: tifa
	// output: memory address of m: 0xc00005c038
	/*
		|--name---|----m----|
		|  0x001  |  0x002  |
		| "tifa"  | p0x001  |
		|---------|---------|
	*/

	fmt.Println("name before update:", name)
	// output: name before update: tifa
	updateNameByPointer(m)
	fmt.Println("name after update:", name)
	// output: name after update: wedge

}
