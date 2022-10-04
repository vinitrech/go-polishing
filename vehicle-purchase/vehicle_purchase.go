package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	switch kind {
	case "truck":
		return true
	case "car":
		return true
	default:
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	append := " is clearly the better choice."
	if option1 < option2 {
		return option1 + append
	}

	return option2 + append
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	resellPrice := 0.0

	if age < 3 {
		resellPrice = originalPrice * 0.80
	} else if age >= 3 && age < 10 {
		resellPrice = originalPrice * 0.70
	} else {
		resellPrice = originalPrice * 0.50
	}

	return resellPrice
}
