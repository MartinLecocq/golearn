package main

import (
	"fmt"
	"unicode/utf8"
)

// Variable declaration spot
var mystring string = "Tür"

func main() {
	fmt.Println(`Hello world!`)

	// Testing the counting of characters in a word with 2 different functions
	fmt.Printf("Der String %q besteht aus %d Zeichen. Hoppla! Ich meine tatsächlich %d Zeichen.\n",
		mystring, len(mystring), utf8.RuneCountInString(mystring))
	// I now want to print my string in the hexadecimal form
	fmt.Printf("Der String %[1]q wird auf Hexadezimal so geschrieben: %[1]X\n", mystring)

}
