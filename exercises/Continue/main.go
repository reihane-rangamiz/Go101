package main

import "fmt"

func main() {

	sum := 0

	for i := 0; i < 100; i++ {

		if i == 50 {
			continue
		}
		sum += i
	}

	fmt.Printf("The final value is: %v", sum)
}
