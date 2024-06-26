package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)
    units["quarter_of_a_dozen"] = 3
    units["half_of_a_dozen"] = 6
    units["dozen"] = 12
    units["small_gross"] = 120
    units["gross"] = 144
    units["great_gross"] = 1728
    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, check := units[unit]
    if(!check) {
        return false
    } else {
        _ , check := bill[item]
        if(!check) {
            bill[item] = units[unit]
        } else {
            bill[item] += units[unit]
        }
        return true
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, checkUnit := units[unit]
    _, checkItem := bill[item]
    if(!checkUnit || !checkItem) {
        return false
    } else {
        count, _ := bill[item]
        if(count<units[unit]) {
			return false
        } else {
            bill[item] -= units[unit]
            if (bill[item]==0) {
                delete(bill, item)
            }
            return true
        }
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itm, checkItem := bill[item]
    if(!checkItem) {
        return 0, false
    } else {
        return itm, true
    }
}
