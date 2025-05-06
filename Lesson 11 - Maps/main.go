package main

import "fmt"

func main() {
	manu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(manu)
	// Output: map[pie:7.99 salad:6.99 soup:4.99 toffee pudding:3.55]
	fmt.Println(manu["pie"])
	// Output: 7.99

	// loop through the map
	for k, v := range manu {
		fmt.Println(k, "-", v)
	}
	// Output:
	// soup - 4.99
	// pie - 7.99
	// salad - 6.99
	// toffee pudding - 3.55

	// ints as key type
	phonebook := map[int]string{
		267584967: "mario",
		984751242: "luigi",
		812473243: "peach",
	}
	fmt.Println(phonebook)
	// Output: map[267584967:mario 812473243:peach 984751242:luigi]
	fmt.Println(phonebook[267584967])
	// Output: mario

	// update the map
	phonebook[984751242] = "bowser"
	fmt.Println(phonebook)
	// Output: map[267584967:mario 812473243:peach 984751242:bowser]
	phonebook[812473243] = "yoshi"
	fmt.Println(phonebook)
	// Output: map[267584967:mario 812473243:yoshi 984751242:bowser]
}
