package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	orderedmap "github.com/wk8/go-ordered-map"
)

func updateOrder(update string, order *map[MenuItem]int) {
	details := strings.Split(update, " ")
	itemId, _ := strconv.Atoi(details[0])
	itemQuantity, _ := strconv.Atoi(details[1])
	bengaliMenu, northIndianMenu := getMenu()

	if itemId < 0 || itemId > TotalItems {
		fmt.Println("Invalid Item ID")
	} else if itemId <= bengaliMenu.Len() {
		item, _ := bengaliMenu.Get(itemId)
		menuItem := item.(MenuItem)
		if itemQuantity <= 0 {
			delete(*order, menuItem)
			fmt.Printf("%s removed from your order.\n", menuItem.Name)
		} else {
			(*order)[menuItem] = itemQuantity
			fmt.Printf("\n%s updated in your order.\n", menuItem.Name)
			fmt.Printf("Prensent order of %s is : %d\n", menuItem.Name, (*order)[menuItem])
		}

	} else {
		item, _ := northIndianMenu.Get(itemId)
		menuItem := item.(MenuItem)
		if itemQuantity <= 0 {
			delete(*order, item.(MenuItem))
			fmt.Printf("%s removed from your order.\n", menuItem.Name)
		} else {
			(*order)[item.(MenuItem)] = itemQuantity
			fmt.Printf("\n%s updated in your order.\n", menuItem.Name)
			fmt.Printf("Prensent order of %s is : %d\n", menuItem.Name, (*order)[menuItem])
		}

	}
}

func deleteItem(update string, order *map[MenuItem]int) {
	details := strings.Split(update, " ")
	itemId, _ := strconv.Atoi(details[0])
	bengaliMenu, northIndianMenu := getMenu()
	var item interface{}
	var menuItem MenuItem

	if itemId < 0 || itemId > TotalItems {
		fmt.Printf("Invalide item ID\n")
	} else if itemId <= bengaliMenu.Len() {
		item, _ = bengaliMenu.Get(itemId)
		menuItem = item.(MenuItem)
		delete(*order, menuItem)
		fmt.Printf("%s removed from your order.\n", menuItem.Name)
	} else {
		item, _ = northIndianMenu.Get(itemId)
		menuItem = item.(MenuItem)
		delete(*order, menuItem)
		fmt.Printf("%s removed from your order.\n", menuItem.Name)
	}
}

func modifyOrder(order *map[MenuItem]int) {
	reader := bufio.NewReader(os.Stdin)
	var choice string
	var orderUpdate string

	for {
		fmt.Println("\nPress '1' to update item quantity.")
		fmt.Println("Press '2' to delete an item from the order list.")
		fmt.Println("Press F to finalize the order")

		choice, _ = reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice[0] == 'F' || choice[0] == 'f' {
			break
		}

		// choice, _ = reader.ReadString('\n')
		// choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			fmt.Println("Enter the item Id followed by updated quantity :")
			orderUpdate, _ = reader.ReadString('\n')
			orderUpdate = strings.TrimSpace(orderUpdate)
			updateOrder(orderUpdate, order)
		case "2":
			fmt.Println("Enter the item ID you want to delete from the order:")
			orderUpdate, _ = reader.ReadString('\n')
			orderUpdate = strings.TrimSpace(orderUpdate)
			deleteItem(orderUpdate, order)
		}

	}

}

func takeOrder(bengaliMenu *orderedmap.OrderedMap, northIndianMenu *orderedmap.OrderedMap) {
	order := make(map[MenuItem]int)
	reader := bufio.NewReader(os.Stdin)
	var orderID int
	var choice string
	var item MenuItem
	TotalItems := (bengaliMenu.Len() + northIndianMenu.Len())
	fmt.Print("\n")

	for {
		fmt.Printf("\nEnter your item of choice [To Exit enter Q]: ")
		choice, _ = reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice[0] == 'Q' || choice[0] == 'q' {
			break
		}

		orderID, _ = strconv.Atoi(choice)
		if orderID <= 0 || orderID > TotalItems {
			fmt.Println("Invalid Choice")
			continue
		}

		if orderID <= 20 {
			menuItem, _ := bengaliMenu.Get(orderID)
			item = menuItem.(MenuItem)
		} else {
			menuItem, _ := northIndianMenu.Get(orderID)
			item = menuItem.(MenuItem)
		}

		fmt.Printf("How many %s do you want?: ", item.Name)
		quantityInput, _ := reader.ReadString('\n')
		quantityInput = strings.TrimSpace(quantityInput)
		quantity, err := strconv.Atoi(quantityInput)
		if err != nil || quantity < 0 {
			fmt.Println("Invalid quantity")
			continue
		}
		order[item] = quantity
		fmt.Printf("You just ordered %d %s\n", quantity, item.Name)
	}
	printBill(order)
	fmt.Print("\nDo you wish to change your order?\n")

	choice, _ = reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice[0] == 'y' || choice[0] == 'Y' {
		modifyOrder(&order)
		printBill(order)
	} else {
		fmt.Println("\nThank you for your order")
		// printBill(order)
	}

	fmt.Print("We will retun with your food in no time.\n\n")
}

func orderProcessing() {
	fmt.Print("Your order is being processed. Please wait ")

	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Print(". \n")
	time.Sleep(1 * time.Millisecond)

	fmt.Println("Table served!")
	time.Sleep(1 * time.Second)
	fmt.Println("Enjoy your food")
}
