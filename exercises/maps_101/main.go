package main

import (
	"fmt"
	"strconv"
)

func main() {

	var profile map[string]string
	profile = make(map[string]string)

	profile["FirstName: "] = "Reyhane"
	profile["LastName: "] = "Rangamiz"
	profile["Age"] = strconv.FormatUint(uint64(24), 10)

	fmt.Print(profile)
}
