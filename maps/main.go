package main

import "fmt"

// maps are like dictionary (Key, Value)

func main() {

	// option 1 to decalre a map
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
		"white": "#fffff",
	}
	fmt.Println(colors)

	// option 2
	Kolors := make(map[int]string)
	Kolors[10] = "#fffff"
	fmt.Println(Kolors)
	delete(Kolors, 10) // delete value from maps
	fmt.Println(Kolors)

	// iterating over maps
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
