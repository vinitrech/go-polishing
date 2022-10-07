package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, isPresent := units[unit]; !isPresent { // returns false if the unit is not in the units map
		return false
	}

	bill[item] += units[unit]

	return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	if _, isPresent := bill[item]; !isPresent {
		return false
	}

	if _, isPresent := units[unit]; !isPresent {
		return false
	}

	if bill[item]-units[unit] < 0 {
		return false
	} else if bill[item]-units[unit] == 0 {
		delete(bill, item)
	} else {
		bill[item] -= units[unit]
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if _, isPresent := bill[item]; !isPresent {
		return 0, false
	}

	return bill[item], true
}
