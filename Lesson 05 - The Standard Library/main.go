package main

import (
	"fmt"
	"sort"
	"strings"
)

// golang standard library: https://pkg.go.dev/std

func main() {
	// strings
	greeting := "Hello there friends!"

	fmt.Println(strings.Contains(greeting, "Hello")) // true
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	fmt.Println("original string:", greeting)

	// sort
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages)
	fmt.Println(ages)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)

	// search
	index := sort.SearchInts(ages, 30)
	fmt.Println(index) // 2

	fmt.Println(sort.SearchStrings(names, "bowser")) // 0
}
