package main

import "fmt"

func main() {
	var aninteger int = 165_256
	var anotherinteger int = 78
	fmt.Println("Hi")
	fmt.Println(aninteger)
	fmt.Printf("%v:  %#[1]b  %#[1]o  %#[1]x\n\n", anotherinteger)

	var a int = 78
	var b int = 0b1010_0011 // Dezimalzahl 163
	// Ausgabe ohne Präfix
	fmt.Printf("%v: %[1]b %[1]o %[1]x\n", a)
	// Ausgabe mit Präfix
	fmt.Printf("%d: %#[1]b %#[1]o %#[1]x\n", a)
	// alternative Ausgabe
	fmt.Printf("%v: %[1]b %[1]O %[1]X\n", b)
	fmt.Println("\n")

	var c uint8 = 0b00001010
	fmt.Printf("NOT c (%08[1]b): %08[2]b\n", c, ^c)
}
