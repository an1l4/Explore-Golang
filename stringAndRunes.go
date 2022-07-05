package main

import "fmt"

func main() {

	var s string = "this is a string"
	fmt.Printf("value=%v , type=%T\n", s, s)
	fmt.Printf("value=%v , type=%T\n", s[0], s[0])

	//converting them into a slice/collection of bytes
	b := []byte(s)
	fmt.Printf("value=%v ,type=%T\n", b, b)

	//Rune (utf-32 character)
	//Type alias for int32
	var r rune = 'a'
	fmt.Printf("value=%v , type=%T\n", r, r)

}
