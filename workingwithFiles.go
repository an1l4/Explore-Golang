package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "this content needs to go to the file"

	//creating a file
	file, err := os.Create("./myfile.txt")

	checkNilErr(err)

	//write the conetent to our newly created file
	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("the length of the file is :", length) //this will give the length of the content
	defer file.Close()                                 //once done with all operations, we need to close the file
	//use defer while closing the file so it will close the file after all operation
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("databytes is ", databyte) //while reading the file we get the raw data like bytes or kind of json

	//to get the actual data
	fmt.Println("converting databytes :", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)

	}
	//we use err checking multiple times so we can wrap the code in a func and use multiple calls

}
