package main

import (
	"fmt"
	"time"

	orderedmap "github.com/wk8/go-ordered-map"
)

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
	time.Sleep(100 * time.Millisecond)
	fmt.Print("************************** Bengali Dishes ****************************\n")
	time.Sleep(100 * time.Millisecond)
	var item MenuItem
	for pair := bengaliMenu.Oldest(); pair != nil; pair = pair.Next() {
		item = pair.Value.(MenuItem)
		fmt.Printf("%-2d. %-30s  ₹%-5d\n   %s\n", item.ID, item.Name, item.Price, item.Description)
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println()
	fmt.Print("************************** North Indian Dishes ****************************\n")
	time.Sleep(100 * time.Millisecond)
	for pair := northIndianMenu.Oldest(); pair != nil; pair = pair.Next() {
		item = pair.Value.(MenuItem)
		fmt.Printf("%-2d. %-30s  ₹%-5d\n   %s\n", item.ID, item.Name, item.Price, item.Description)
		time.Sleep(20 * time.Millisecond)
	}
}
