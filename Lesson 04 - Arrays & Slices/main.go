package main

import "fmt"

func main() {

	// 1. Arrays
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}
	fmt.Println("Ages:", ages, ", length: ", len(ages))

	names := [4]string{"John", "Doe", "Smith", "Yoshi"}
	fmt.Println("Names:", names, ", length: ", len(names))
	names[1] = "Luigi"
	fmt.Println("Names:", names, ", length: ", len(names))

	// 2. Slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println("Scores:", scores, ", length: ", len(scores))

	// 3. Slice Ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Koopa")
	fmt.Println("Range One:", rangeOne, ", length: ", len(rangeOne))

}
