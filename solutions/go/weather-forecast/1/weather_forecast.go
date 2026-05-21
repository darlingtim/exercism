// Package weather enables you to know the current weather condition in your city or location.
package weather




var (
    // CurrentCondition fetches the weather condition.
	CurrentCondition string
    // CurrentLocation is used to collect your current. location.
	CurrentLocation  string
)
// Forecast function displays the weather condition of your current location to you.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
