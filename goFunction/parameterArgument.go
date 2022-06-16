package main

import "fmt"

func familyName(fname string) {  //fname is parameter
	fmt.Println("Hello",fname,"Cooper")
}

func main() {
	familyName("Liam")  //Liam,John,Jane are arguments
	familyName("John")
	familyName("Jane")
}