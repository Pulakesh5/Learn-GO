package main

import (
	"fmt"
	"time"
)

func printBill(order map[MenuItem]int) {
	printGenerateBill()

	subTotalExclGST := 0
	itemSubTotal := 0

	printDashedLine()
	fmt.Printf("%-10s %-20s %-15s %-10s %-10s\n", "Item ID", "Item Name", "Price(₹)", "Quantity", "Total Price(₹)")
	printDashedLine()

	for orderItem := range order {
		itemSubTotal = orderItem.Price * order[orderItem]
		subTotalExclGST += itemSubTotal
		fmt.Printf("%-10d %-20s  ₹%-15d  %-10v  ₹%-10d\n", orderItem.ID, orderItem.Name, orderItem.Price, order[orderItem], itemSubTotal)
		time.Sleep(time.Second * 1)
	}

	printDashedLine()
	fmt.Printf("%-10s %-50s ₹%-10.2f\n", " ", "Subtotal (excluding GST):", float64(subTotalExclGST))
	totalBillinclGST := float64(subTotalExclGST) * 1.05
	fmt.Printf("%-10s %-50s ₹%-10.2f\n", " ", "Total Bill (including GST):", totalBillinclGST)
	printDashedLine()
}
