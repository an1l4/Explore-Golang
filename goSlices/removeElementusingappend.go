package main

import "fmt"

func main() {

	var courses = []string{"react", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	//using index variable removing values from slice
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
