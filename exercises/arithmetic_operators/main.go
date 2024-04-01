package main

import "fmt"

var (
	packagesToDeliver = 100
	deliveredPackages = 20
	customers         = 4
)

func main() {

	fmt.Printf("i have %d packages to deliver to the customers! \n", packagesToDeliver)

	remainingPackages := packagesToDeliver - deliveredPackages
	fmt.Printf("%d packages have been delivered to the customers and %d packages remaine in stock! \n", deliveredPackages, remainingPackages)

	//Comment out lines 15 and 16 and uncomment lines below.then, run the program to see the result:

	//packagesToDeliver -= deliveredPackages
	//fmt.Printf("%d packages have been delivered to the customers and %d packages remained in the stock! \n", deliveredPackages, packagesToDeliver)

	packagesForEachCustomer := packagesToDeliver / customers
	fmt.Printf("We have %v customers and %d packages will be dedicated to each customer :) \n", customers, packagesForEachCustomer)

}
