package main

import (
	"fmt"
	"strings"
	"time"
)

func printDashedLine() {
	dashedLine := strings.Repeat("-", 75)
	fmt.Print("+")
	for _, dash := range dashedLine {
		fmt.Print(string(dash))
		time.Sleep(time.Millisecond * 7)
	}
	fmt.Print("+\n")
}
func printGenerateBill() {
	generateBillHeader := "****************************** Generating Bill ******************************"
	fmt.Println()
	for _, char := range generateBillHeader {
		fmt.Print(string(char))
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Println()
}

func greet(customerName string) {
	fmt.Printf("%52s %s\n", "Namaskar ", customerName)
	fmt.Printf("%72s\n", "_/\\_ Welcome to Pancha Byanjan! _/\\_ ")
	fmt.Println()
}

func sayBye(customerName string) {
	fmt.Println()
	fmt.Printf("%17s", " ")
	fmt.Printf("_/\\_ Thank you %v for visiting Pancha Byanjan! _/\\_\n", customerName)
	fmt.Printf("%55s\n", "We hope to see you again!")
	fmt.Printf("%51s\n", "Have a nice day! \n")
}
