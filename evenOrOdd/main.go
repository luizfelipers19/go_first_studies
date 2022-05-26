package main

import "fmt"

func main() {
	fmt.Println("Test Print")

	var numbers [11]int
	for i := range numbers {
		numbers[i] = i
		if numbers[i]%2 == 0 {
			fmt.Println(numbers[i], " is even.")
		} else {
			fmt.Println(numbers[i], "is odd.")
		}
	}

	fmt.Println(numbers)

}
