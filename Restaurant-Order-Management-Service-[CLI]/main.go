package main

import (
	"fmt"
)

type MenuItem struct {
	ID          int
	Name        string
	Description string
	Price       int
}

var TotalItems int

func main() {
	var customerName string
	fmt.Print("What is your name? ") // prompt
	fmt.Scan(&customerName)
	greet(customerName)
	bengaliMenu, northIndianMenu := getMenu()
	TotalItems = bengaliMenu.Len() + northIndianMenu.Len()
	showMenu(bengaliMenu, northIndianMenu)
	takeOrder(bengaliMenu, northIndianMenu)
	orderProcessing()
	sayBye(customerName)
}
