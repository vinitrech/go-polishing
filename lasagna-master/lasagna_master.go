package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averageTimePerLayer int) int {
	if averageTimePerLayer == 0 {
		return len(layers) * 2
	}

	return len(layers) * averageTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodlesCounter := 0
	sauceCounter := 0.0

	for _, layer := range layers {
		if layer == "sauce" {
			sauceCounter += 0.2
		} else if layer == "noodles" {
			noodlesCounter += 50
		}
	}

	return noodlesCounter, sauceCounter
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scaledQuantities int) []float64 {
	scaledRecipe := []float64{}

	for _, value := range quantities {
		scaledRecipe = append(scaledRecipe, (value/2)*float64(scaledQuantities))
	}

	return scaledRecipe
}
