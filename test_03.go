package main

import "fmt"

func main() {
	var gleitkommazahl_01 float32 = 2.0 / 3.0
	var gleitkommazahl_02 float32 = 0.67
	var ergebnis float32 = gleitkommazahl_02 - gleitkommazahl_01
	fmt.Println(ergebnis)
	// Die Mantisse ist bei float32 zwar viel großer als 8 bit (hier 23 bits). Deshalb ist die Rundungsfehler viel
	// kleiner (0,000093%).

	// Test auf Boolesche Velgeichsoperatoren für Strings
	var bool_ergebnis = "bc" > "ad"
	fmt.Println(bool_ergebnis)
	// Test auf String Datentyp: Unicode Zeichen
	var univar1 string = "Unendlich ist \u221e und Smiley ist \U0001f600"
	var univar2 string = "Griechisch: \u0370\u0371\u037C"
	fmt.Println(univar1)
	fmt.Println(univar2)
}
