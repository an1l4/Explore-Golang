package main

import "fmt"

func familyName(fname string, age int) {
	fmt.Println("Hello",fname,"Cooper you are",age,"Now")
}

func main() {
	familyName("Liam",10)
	familyName("John",20)
	familyName("Jane",30)
}