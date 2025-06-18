package main

import "fmt"

func main() {
	mybill := newBill("Mario's bill")
	fmt.Println(mybill)
	// $ go run main.go bill.go
	// output: {Mario's bill map[] 0}

	fmt.Println(mybill.format())
}
