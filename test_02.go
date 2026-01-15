package main

import "fmt"

func main() {
	var c int8 = -13
	var d = uint8(c)
	var clinks = c << 4
	// var cLeft int8 = c << 1
	// var cRight int8 = c >> 1
	fmt.Printf("c: %08b Dez: %3v \n", c, c)
	fmt.Printf("c:2-Kompl: %08b Dez: %3v\n", uint8(c), c)
	// Test of shifting c 4 times to the left, to see the overflow
	fmt.Printf("c << 4: %08b Dez: %3v\n", clinks, clinks)
	// Seeing the value of c, when it's type is turned in uint8
	fmt.Println(uint8(c))
	// Testing if the same happens when we use another variable
	fmt.Printf("d: %08b Dez: %3v \n", d, d)

	// Testing with another value
	fmt.Println("\n")
	var e int8 = -8
	fmt.Printf("e: %08b Dez: %3v \n", uint8(e), e)
	var f int8 = e << 4
	fmt.Printf("e << 4: %08b Dez: %3v \n", uint8(f), f)
}
