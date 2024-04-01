package main

import "fmt"

func main() {

	x := 2000

	result1 := x > 50 && x < 2020
	result2 := x > 50 && x%2 == 0
	result3 := x%2 == 1

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	x += 5

	result4 := x > 50 && x < 3000
	result5 := x > 50 || x < 2020
	result6 := x > 50 || x%2 == 0

	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)

}
