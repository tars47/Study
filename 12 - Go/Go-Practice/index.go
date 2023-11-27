package main

import (
	"fmt"
	//"math/rand"
)

func init() {
	fmt.Printf("I am in init function \n")
}

func main() {

	/////////////////////////////////////////////////////////////////////////////

	// x := rand.Intn(251)
	// fmt.Printf("the value of x is %v.\n", x)
	// if x >= 0 && x <= 100 {
	// 	fmt.Printf("between 0 and 100.\n")
	// } else if x >= 101 && x <= 200 {
	// 	fmt.Printf("between 101 and 200.\n")
	// } else if x >= 201 && x <= 250 {
	// 	fmt.Printf("between 201 and 250.\n")

	// }

	////////////////////////////////////////////////////////////////////////////////

	// switch {
	// case x >= 0 && x <= 100:
	// 	fmt.Printf("switch - between 0 and 100.\n")
	// case x >= 101 && x <= 200:
	// 	fmt.Printf("switch - between 101 and 200.\n")
	// case x >= 201 && x <= 250:
	// 	fmt.Printf("switch - between 201 and 250.\n")
	// }

	/////////////////////////////////////////////////////////////////////////////

	// for i := 0; i < 100; i++ {
	// 	fmt.Printf("iteration number %v\n", i)
	// }

	/////////////////////////////////////////////////////////////////////////////

	// arr := [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)",
	// 	"Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)",
	// 	"Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
	// 	"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
	// 	"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar",
	// 	"Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)",
	// 	"Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)",
	// 	"Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough",
	// 	"Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different",
	// 	"Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie",
	// 	"Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}

	// fmt.Printf("%v \n %T", len(arr), arr)

	/////////////////////////////////////////////////////////////////////////////

	// xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)",
	// 	"Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)",
	// 	"Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
	// 	"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
	// 	"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar",
	// 	"Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)",
	// 	"Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)",
	// 	"Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough",
	// 	"Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different",
	// 	"Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie",
	// 	"Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}

	// for i, v := range xs {
	// 	fmt.Printf("%v - %v \n", i, v)
	// }

	/////////////////////////////////////////////////////////////////////////////

	/*
			    Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.
		        key                 value
		        `bond_james`        `shaken, not stirred`, `martinis`, `fast cars`
		        `moneypenny_jenny`  `intelligence`, `literature`, `computer science`
		        `no_dr`             `cats`, `ice cream`, `sunsets`

	*/
	mp := map[string][]string{
		`bond_james`:       {`shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`: {`intelligence`, `literature`, `computer science`},
		`no_dr`:            {`cats`, `ice cream`, `sunsets`},
	}

	fmt.Println(mp)
}
