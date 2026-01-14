package main

import "fmt"

func main() {
	var c int8 = -13
	// var cLeft int8 = c << 1
	// var cRight int8 = c >> 1
	fmt.Printf("c:2-Kompl: %08b Dez: %3v\n", uint8(c), c)
	fmt.Println(uint8(c))
}
