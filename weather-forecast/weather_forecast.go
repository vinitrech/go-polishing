// Package weather takes a location and a condition and returns the current weather condition.
package weather

// CurrentCondition used to search for weather information.
var CurrentCondition string

// CurrentLocation used to search for weather information.
var CurrentLocation string

// Forecast returns the current condition of the city's weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
