package purchase

import "slices"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	requireLicense := []string{"car", "truck"}

	return slices.Contains(requireLicense, kind)
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	choice := option2

	if option1 < option2 {
		choice = option1
	}

	return choice + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	percentageOfPrice := 100

	if age < 3 {
		percentageOfPrice = 80
	} else if age >= 3 && age < 10 {
		percentageOfPrice = 70
	} else if age >= 10 {
		percentageOfPrice = 50
	}

	return originalPrice * (float64(percentageOfPrice) / 100)
}
