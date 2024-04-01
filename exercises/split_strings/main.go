package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "Alexander Supertramp"
	separatedName := strings.Fields(name)
	fmt.Println("the first name is", separatedName[0])

	name2 := "Rick Sanchez"
	fmt.Println("the first name is", strings.Fields(name2)[0])
	fmt.Println("the last name is", strings.Fields(name2)[1])

}
