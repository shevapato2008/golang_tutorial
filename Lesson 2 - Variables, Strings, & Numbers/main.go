package main

import "fmt"

var someName1 = "hello" // OK
// someNmae2 := "hello" // ERROR: syntax error: non-declaration statement outside function body

func main() {

	// strings
	var nameOne string = "mario"
	var nameTwo = "luigi" // automatically infers the type
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "browser"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi" // short hand declaration
	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40 // short hand declaration
	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 256
	fmt.Println(numOne, numTwo, numThree)

	// float
	var scoreOne float32 = 25.5
	var scoreTwo float64 = 2524514155.5
	scoreThree := 1.5 // short hand declaration, default type is float64
	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
