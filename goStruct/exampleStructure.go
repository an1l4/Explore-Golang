package main

import "fmt"

//struct - 1st letter should be caps if its exporting
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	anila := User{"Anila", "example@gmail.com", true, 17}
	fmt.Println(anila)

	//%+v using to get more details in struct
	fmt.Printf("Anilas Details are %+v\n", anila)

	//get particular data
	fmt.Printf("Anilas email is : %v and Anila s age is %v \n", anila.Email, anila.Age)

}
