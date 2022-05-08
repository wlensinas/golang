package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 11; i++ {
		if numbers[i]%2 == 0 {
			fmt.Printf("%v is even \n", numbers[i])
		} else {
			fmt.Printf("%v is odd \n", numbers[i])
		}
	}
}
