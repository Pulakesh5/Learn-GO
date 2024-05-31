// Package weather provides tools to forecast the current weather conditions of various cities.
package weather

// CurrentCondition tells us about the condition of the weather currently.
var CurrentCondition string

// CurrentLocation tells us about the location of the city we are talking about.
var CurrentLocation string

// Forecast returns the weather forecast of the city for which weather conditions were passed.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
