package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    storeUnitScore := map[string]int{
        "quarter_of_a_dozen" : 3,
        "half_of_a_dozen" : 6,
        "dozen" : 12,
        "small_gross" : 120,
        "gross" : 144,
        "great_gross" : 1728,
    }
    return storeUnitScore
    
	panic("Please implement the Units() function")
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    bill := make(map[string]int)
    return bill
	panic("Please implement the NewBill() function")
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    unitScore, exists := units[unit]
    if !exists {
       return false   
    }
        
    bill[item] += unitScore
    return true
        
    
	panic("Please implement the AddItem() function")
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    // 1. Check if the item is even on the bill
	currentBillAmount, itemExists := bill[item]
	if !itemExists {
		return false
	}

	// 2. Check if the unit deduction amount exists in configurations
	unitScore, unitExists := units[unit]
	if !unitExists {
		return false
	}

	// 3. Calculate what the new amount would be
	newAmount := currentBillAmount - unitScore

	// Fail safety guard: If the removal amount makes the bill negative, reject the operation
	if newAmount < 0 {
		return false
	}

	// 4. Update or clear the map entry
	if newAmount == 0 {
		// If no value is left, completely wipe the key from the map
		delete(bill, item)
	} else {
		// Otherwise, update the map with the reduced value
		bill[item] = newAmount
	}

	return true
	panic("Please implement the RemoveItem() function")
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    quantity, exists := bill[item]
    if !exists {
        return 0, false
    }
    return quantity, true
	panic("Please implement the GetItem() function")
}
