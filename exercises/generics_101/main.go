package main

import "fmt"

type CType interface {
	int | int32 | uint | uint32
}

// in this code, by having a generic function, you will remove some of the input's type limits and make it more general !

func sum1[T CType](A, B T) T {
	return A + B
}

// and here you remove the quantity limits, means you can have 1 input or a thousands --> (nums ...T)

func sum2[T CType](nums ...T) T {
	var result T
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {
	fmt.Println(sum1(10, 4))

	// create a list of uint
	numbers := []uint{20, 22, 43, 84, 95, 1000}

	// use the sum2 function to get the summation
	fmt.Println(sum2(numbers...))
}
