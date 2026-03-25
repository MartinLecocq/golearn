package main

import (
	"fmt"
	"unicode/utf8"
)

// Variable declaration spot
var mystring string = "Tür"
var mybyte int = 0

func main() {
	fmt.Println(`Hello world!`)

	// Testing the counting of characters in a word with 2 different functions
	fmt.Printf("Der String %q besteht aus %d Zeichen. Hoppla! Ich meine tatsächlich %d Zeichen.\n",
		mystring, len(mystring), utf8.RuneCountInString(mystring))
	// I now want to print my string in the hexadecimal form and showing an exact byte.
	fmt.Printf("Der String %[1]q wird auf Hexadezimal so geschrieben: %[1]X\n", mystring)
	fmt.Printf("Der Byte %v meines String sieht so auf Hexadezimal aus: %X\n", mybyte, mystring[mybyte])
	// With the following function, I can print the letter as s String via a low:high index.
	// But I still need to know how many bytes my character needs.
	fmt.Printf("Die Buchstabe ist: %v", mystring[mybyte:mybyte+1])
}
