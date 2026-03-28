package main

import "fmt"

var amountApples int = 5

func main() {
	fmt.Printf("The result is: %[1]v - The returned value is of type %[1]T.\n", float64(amountApples)/2)
	fmt.Printf("The result is: %[1]v - The returned value is of type %[1]T.\n", amountApples/2)
	fmt.Printf("The result is: %[1]v - The returned value is of type %[1]T.\n", float64(amountApples/2))
}
