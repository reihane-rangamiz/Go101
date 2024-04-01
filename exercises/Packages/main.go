package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// 1.
	currentTime := time.Now()
	fmt.Println("The time now is: ", currentTime)
	//fmt.Println("The time now is: ", time.Now())

	// 2.
	randInt := rand.Intn(100)
	fmt.Println("The random integer is: ", randInt)
	//fmt.Println("The random integer is: ", rand.Intn(100))

	// 3.
	sqrt := math.Sqrt(9)
	fmt.Println("The square of 9 is: ", sqrt)
	//fmt.Println("The square of 9 is: ", math.Sqrt(9))

}
