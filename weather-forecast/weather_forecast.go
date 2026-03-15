// Package weather provides tools to get the forecast.
package weather

var (
	// CurrentCondition represents the weather codition right now as a string.
	CurrentCondition string
	// CurrentLocation represents where the forecast is for as a string.
	CurrentLocation string
)

// Forecast returns a formatted forecast description including the CurrentLocation and CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
