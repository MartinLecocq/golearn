package main

import "fmt"

func main() {
	var aninteger int = 165_256
	var anotherinteger int = 78
	fmt.Println("Hi")
	fmt.Println(aninteger)
	fmt.Printf("%v:  %#[1]b  %#[1]o  %#[1]x\n", anotherinteger)
}
