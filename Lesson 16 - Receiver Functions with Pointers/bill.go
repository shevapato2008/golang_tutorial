package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format the bill
func (b bill) format() string {
	// cycle through the items
	fs := "Bill Breakdown:\n"
	var total float64 = 0
	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v: ...$%0.2f \n", k, v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v: ...$%0.2f \n", "tip", b.tip)

	// total
	fs += fmt.Sprintf("%-25v: ...$%0.2f \n", "total", total + b.tip)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	// update the tip
	b.tip = tip		// DON'T have to dereference b because b is a pointer receiver
	                // b is dereferenced automatically
}

// add item to bill
func (b *bill) addItem(name string, price float64) {
	if b.items == nil {
		b.items = map[string]float64{}
	}
	b.items[name] = price
}