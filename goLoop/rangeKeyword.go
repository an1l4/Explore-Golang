package main

import "fmt"

func main() {
	fruits := [3]string{"apple","orange","banana"}

	for index, value := range fruits {
		fmt.Printf("%v \t %v \n",index,value)
	}

	//using slice
nums := []int{2,3,4}
sum := 0
for _,value := range nums {
	sum += value
	fmt.Println(value)
}
fmt.Println("sum :",sum)

for index, value := range nums {
	if value == 3 {
		fmt.Println("index :",index)
	}
}
//using map

kvs := map[string]string{"a": "apple","b":"banana"}
for index,value := range kvs {
	fmt.Printf("%s -> %s \n",index,value)
}
for index := range kvs {
	fmt.Println("keys :",index)
}
for index, value := range "go" {
	fmt.Println(index,value)
}

	
}
//syntx for index,value :=range arrayname/slicename/mapname

