package main

import "fmt"

func main() {
	var jediName string
	var favoriteDroid string

	fmt.Println("🌌 Welcome, brave Jedi! May the Force be with you.")
	fmt.Println("✨What's your Jedi name, master of the Force?")
	fmt.Scan(&jediName)

	fmt.Println("✨Tell me the name of your favorite droid:")
	fmt.Scanln(&favoriteDroid)

	fmt.Printf("\nA Jedi named %s, a guardian of peace and justice, with their loyal droid %s. May the Force guide you on your journey!", jediName, favoriteDroid)

}
