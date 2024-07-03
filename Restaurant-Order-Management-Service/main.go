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

type MenuItem struct {
	ID          int
	Name        string
	Description string
	Price       int
}

var TotalItems int

func getMenu() (*orderedmap.OrderedMap, *orderedmap.OrderedMap) { //(map[int]MenuItem, map[int]MenuItem) {
	bengaliMenu := []MenuItem{
		{1, "Machher Jhol", "Traditional Bengali fish curry cooked with potatoes and tomatoes.", 300},
		{2, "Ilish Bhapa", "Steamed Hilsa fish cooked in mustard and coconut paste.", 450},
		{3, "Chingri Malai Curry", "Prawns cooked in a rich coconut milk gravy.", 400},
		{4, "Mangsher Jhol", "Traditional Bengali mutton curry cooked with potatoes.", 250},
		{5, "Aloor Dom", "Spicy potato curry.", 200},
		{6, "Shukto", "A mix of vegetables cooked in a bitter and creamy sauce.", 200},
		{7, "Posto", "Poppy seed paste curry with potatoes.", 180},
		{8, "Luchi", "Deep-fried flatbreads.", 100},
		{9, "Mishti Doi", "Sweetened yogurt.", 120},
		{10, "Rasgulla", "Soft, spongy cottage cheese balls in syrup.", 100},
		{11, "Kosha Mangsho", "Spicy slow-cooked mutton curry.", 380},
		{12, "Cholar Dal", "Split Bengal gram lentils cooked with coconut.", 150},
		{13, "Doi Maach", "Fish cooked in a yogurt-based gravy.", 320},
		{14, "Bhapa Chingri", "Steamed prawns in mustard sauce.", 400},
		{15, "Beguni", "Batter-fried eggplant slices.", 80},
		{16, "Mochar Ghonto", "Banana flower curry.", 180},
		{17, "Kumro Bharta", "Mashed pumpkin with spices.", 150},
		{18, "Chingri Paturi", "Prawns marinated in spices and steamed in banana leaves.", 420},
		{19, "Patishapta", "Rice flour crepes filled with sweetened coconut and jaggery.", 150},
		{20, "Sandesh", "Delicate sweets made from fresh cottage cheese.", 100},
	}

	northIndianMenu := []MenuItem{
		{21, "Butter Chicken", "Chicken cooked in a creamy tomato gravy.", 350},
		{22, "Palak Paneer", "Cottage cheese cubes in a spinach-based gravy.", 280},
		{23, "Rogan Josh", "Aromatic lamb curry.", 400},
		{24, "Dal Makhani", "Creamy lentil dish cooked with butter and spices.", 250},
		{25, "Biryani", "Fragrant rice cooked with meat and aromatic spices.", 300},
		{26, "Tandoori Roti", "Whole wheat flatbread cooked in a tandoor.", 30},
		{27, "Naan", "Soft leavened flatbread cooked in a tandoor.", 40},
		{28, "Paneer Tikka", "Marinated cottage cheese cubes grilled to perfection.", 300},
		{29, "Gulab Jamun", "Deep-fried dough balls soaked in sugar syrup.", 80},
		{30, "Kulfi", "Traditional Indian ice cream.", 120},
		{31, "Chicken Tikka Masala", "Chicken marinated in spices and yogurt, cooked in a tomato-based sauce.", 350},
		{32, "Chole Bhature", "Spicy chickpeas served with deep-fried bread.", 200},
		{33, "Aloo Gobi", "Potatoes and cauliflower cooked with spices.", 180},
		{34, "Bhindi Masala", "Spiced okra curry.", 200},
		{35, "Matar Paneer", "Cottage cheese and peas cooked in a tomato-based sauce.", 250},
		{36, "Lamb Korma", "Lamb cooked in a rich, creamy sauce with nuts.", 420},
		{37, "Malai Kofta", "Deep-fried vegetable balls in a creamy sauce.", 280},
		{38, "Pani Puri", "Hollow dough balls filled with spicy water, tamarind chutney, and potatoes.", 100},
		{39, "Kheer", "Rice pudding with milk, sugar, and nuts.", 120},
		{40, "Jalebi", "Deep-fried spiral-shaped sweets soaked in sugar syrup.", 100},
	}

	// Convert slices to maps
	// bengaliMenuMap := make(map[int]MenuItem)
	bengaliMenuMap := orderedmap.New()
	for _, item := range bengaliMenu {
		// bengaliMenuMap[item.ID] = item
		bengaliMenuMap.Set(item.ID, item)
	}

	// northIndianMenuMap := make(map[int]MenuItem)
	northIndianMenuMap := orderedmap.New()
	for _, item := range northIndianMenu {
		// northIndianMenuMap[item.ID] = item
		northIndianMenuMap.Set(item.ID, item)
	}

	return bengaliMenuMap, northIndianMenuMap
}

func showMenu(bengaliMenu *orderedmap.OrderedMap, northIndianMenu *orderedmap.OrderedMap) {
	fmt.Print("\n\n************************** MENUCARD **************************\n\n")
	fmt.Print("************************** Bengali Dishes ****************************\n")
	var item MenuItem
	for pair := bengaliMenu.Oldest(); pair != nil; pair = pair.Next() {
		item = pair.Value.(MenuItem)
		fmt.Printf("%-2d. %-30s  ₹%-5d\n   %s\n", item.ID, item.Name, item.Price, item.Description)
	}
	fmt.Println()
	fmt.Print("************************** North Indian Dishes ****************************\n")
	for pair := northIndianMenu.Oldest(); pair != nil; pair = pair.Next() {
		item = pair.Value.(MenuItem)
		fmt.Printf("%-2d. %-30s  ₹%-5d\n   %s\n", item.ID, item.Name, item.Price, item.Description)
	}
}

func printDashedLine() {
	dashedLine := strings.Repeat("-", 110)
	fmt.Print("+")
	for _, dash := range dashedLine {
		fmt.Print(string(dash))
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Print("+\n")
}
func printGenerateBill() {
	generateBillHeader := "*********************************************** Generating Bill ***********************************************"
	fmt.Println()
	for _, char := range generateBillHeader {
		fmt.Print(string(char))
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Println()
}
func printBill(order map[MenuItem]int) {
	printGenerateBill()

	subTotalExclGST := 0
	itemSubTotal := 0
	printDashedLine()
	fmt.Printf("%-15s %-30s %-20s %-20s %-20s\n", "Item ID", "Item Name", "Price(₹)", "Quantity", "Total Price(₹)")
	printDashedLine()

	for orderItem := range order {
		itemSubTotal = orderItem.Price * order[orderItem]
		subTotalExclGST += itemSubTotal
		fmt.Printf("%-15d %-30s  ₹%-20d  %-20v  ₹%-20d\n", orderItem.ID, orderItem.Name, orderItem.Price, order[orderItem], itemSubTotal)
		time.Sleep(time.Second * 1)
	}
	printDashedLine()

	fmt.Printf("%67s Subtotal (excluding GST): ₹%5.2f\n", " ", float64(subTotalExclGST))
	totalBillinclGST := float64(subTotalExclGST) * 1.05
	fmt.Printf("%67s Total Bill (including GST): ₹%2.2f\n", " ", totalBillinclGST)

	printDashedLine()
}

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

		if choice == "Q" || choice == "q" {
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
	fmt.Print("Do you wish to change your order?\n")

	choice, _ = reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice[0] == 'y' || choice[0] == 'Y' {
		modifyOrder(&order)
		printBill(order)
	}

	fmt.Print("\nWe will retun with your food in no time.\n\n")
}

func main() {
	fmt.Println("************************** Welcome to Bengal Bay **************************")
	bengaliMenu, northIndianMenu := getMenu()
	TotalItems = bengaliMenu.Len() + northIndianMenu.Len()
	showMenu(bengaliMenu, northIndianMenu)
	takeOrder(bengaliMenu, northIndianMenu)
}
