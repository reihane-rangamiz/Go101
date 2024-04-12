package main

import "fmt"

var (
	name string
	age  int
)

func main() {
	/*

		//// 1.

			for {

				fmt.Println("Please Enter Your Name:")
				fmt.Scan(&name)

				fmt.Println("Please Enter Your Age:")
				fmt.Scan(&age)

				fmt.Printf("your name is %v and your age is %d", name, age)

			}
		/*

		///// 2.

		// foods := []string{"pizza"}
		// foods = append(foods, "hamburger", "salad")


		//First method:
		//uncomment the section below to see the result

			for _, food := range foods {
				fmt.Println("Food: ", food)
			}
	*/

	/*
		// Second method:
		//If you need to take track of the index:

			for i, food := range foods {
				fmt.Printf("Food %v : %v\n", i, food)
			}
	*/

	var shapes = [3]string{"square", "circle", "triangle"}

	for i, shape := range shapes {
		fmt.Printf("This is %v and its index in the array is  %v\n", shape, i)

	}

}
