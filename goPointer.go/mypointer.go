package main

import "fmt"

func main() {
	var ptr *int //creating a memory or creating a pointer using * symbol

	fmt.Println("The value of ptr is", ptr)

	mynumber := 23

	var num = &mynumber //referencing mynumber using & symbol

	fmt.Println("value of num is ", num) //will get memory address
	fmt.Println("value of num is", *num) //want to see what is inside the num which is num holding refernce of mynumber which is 23
	//*num which give you the actual value not the memory address

	*num = *num + 2 //which means 23 + 2 new value goes to mynumber
	fmt.Println("New Value is ", mynumber)
}
