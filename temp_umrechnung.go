package main

import "fmt"

// Zusätzliche Funktion aus Punkt 4.2
func fNachC(tempF float64) float64 {
	return (tempF - 32) / 1.8 // Temperatur in Grad Celsius
}

func main() {
	// Deklaration einer Variablen vom Typ float64
	// und implizite Initialisierung mit Null-Wert
	var tempC float64
	// Zusätzliche Variablen aus Punkt 4.2
	var tempF0 float64 = -459.67
	var tempF1 float64 = 0
	var tempF2 float64 = 32
	var tempF3 float64 = 96

	// Deklaration einer Konstanten vom Typ float64
	const nullpunkt float64 = -273.15

	tempC = 10
	fmt.Println("Temperatur zum Start (°C):", tempC)
	tempC = tempC + 12
	fmt.Println("Temperatur nach Änderung (°C):", tempC)

	var tempF float64
	tempF = 32 + tempC*1.8 // Temperatur in Grad Fahrenheit

	// Deklaration einer Variablen und
	// direkte Initialisierung mit einem Wert
	var tempK float64 = tempC - nullpunkt // Temperatur in Kelvin

	/* Ausgabe der Temperatur in den Einheiten:
	- Celsius
	- Fahrenheit
	- Kelvin
	*/
	var tempText string = "Temperatur in "
	fmt.Println()
	fmt.Println(tempText+"°C:", tempC)
	fmt.Println(tempText+"°F:", tempF)
	fmt.Println(tempText+" K:", tempK)

	// Beispiel aus Punkt 4.2
	fmt.Println("\nBeispiel au Punkt 4.2:")
	fmt.Println(tempF0, "°F entsprechen", fNachC(tempF0), "°C.")
	fmt.Println(tempF1, "°F entsprechen", fNachC(tempF1), "°C.")
	fmt.Println(tempF2, "°F entsprechen", fNachC(tempF2), "°C.")
	fmt.Println(tempF3, "°F entsprechen", fNachC(tempF3), "°C.")
	fmt.Println("Andere Version:")
	fmt.Printf("%v °F entsprechen %v °C.\n", tempF1, fNachC(tempF1))
	fmt.Printf("%7.2f °F entsprechen %7.2f °C.\n", tempF0, fNachC(tempF0))
	fmt.Printf("%7.2f °F entsprechen %7.2f °C.\n", tempF1, fNachC(tempF1))
	fmt.Printf("%7.2f °F entsprechen %7.2f °C.\n", tempF2, fNachC(tempF2))
	fmt.Printf("%7.2f °F entsprechen %7.2f °C.\n", tempF3, fNachC(tempF3))
}
