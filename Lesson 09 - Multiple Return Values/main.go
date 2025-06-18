package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	// Split the name into first and last name
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	// Get the initials
	var initials []string
	for _, name := range names {
		initials = append(initials, string(name[:1]))
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fn1, sn1 := getInitials("Mario Bros")
	fmt.Printf("Initials are %s.%s\n", fn1, sn1)
	// output: Initials are M.B

	fn2, sn2 := getInitials("cloud strife")
	fmt.Printf("Initials are %s.%s\n", fn2, sn2)
	// output: Initials are C.S

	fn3, sn3 := getInitials("barret")
	fmt.Printf("Initials are %s.%s\n", fn3, sn3)
	// output: Initials are B._
}
