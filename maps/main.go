package main

import "fmt"

func main() {
	fmt.Println("In this lesson, we're approaching maps!")

	colors := map[string]string{
		"red":      "#ff0000",
		"green":    "#00ff00",
		"blue":     "#0000ff",
		"bordeaux": "#990033",
		"yellow":   "#ffff00",
	}

	colors["white"] = "#ffffff"

	fmt.Println(colors)

	delete(colors, "bordeaux")
	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%v : %v \n", color, hex)
	}
}
