package lasagna

func PreparationTime(layers []string, preparation_time_per_layer int) int {
	if preparation_time_per_layer < 1 {
		return len(layers) * 2
	}

	return len(layers) * preparation_time_per_layer
}

func Quantities(layers []string) (noodles int, sauce float64) {

	for _, ingredient := range layers {
		if ingredient == "noodles" {
			noodles += 50
		}
		if ingredient == "sauce" {
			sauce += 0.2
		}
	}

	return
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, numberOfPortions int) (scaledQuantities []float64) {

	for _, quantity := range quantities {
		scaledQuantities = append(scaledQuantities, quantity*float64(numberOfPortions)/2.0)
	}

	return scaledQuantities
}
