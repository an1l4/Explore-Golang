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

	anila.GetStatus()
	anila.ChangeEmail()
	fmt.Printf("Anilas email is : %v and Anila s age is %v \n", anila.Email, anila.Age)

}

//if exporting name of function starts with caps
func (u User) GetStatus() { //method(function + struct)
	fmt.Println("Is user active", u.Status)

}

func (u User) ChangeEmail() {
	u.Email = "test@dev.go"
	fmt.Println("New email is ", u.Email)
}
