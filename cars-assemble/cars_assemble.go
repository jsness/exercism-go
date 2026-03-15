package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	fraction := successRate / 100.0

	return float64(productionRate) * fraction
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)

	return int(workingCarsPerHour / 60.0)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	extraCars := carsCount % 10
	groupedCars := carsCount - extraCars
	totalGroups := groupedCars / 10

	groupCost := 95000
	singleCost := 10000

	return uint(totalGroups*groupCost + extraCars*singleCost)
}
