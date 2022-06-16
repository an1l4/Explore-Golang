package main

import (
	"fmt"
	"time"
)

func main() {
	i:=2
	fmt.Println("write",i,"as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2 :
		fmt.Println("two")
	case 3 :
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday :
		fmt.Println("its a weekend")
		default :
		fmt.Println("Its a weekday")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12 :
		fmt.Println("Its before noon")
		default :
		fmt.Println("Its after noon")
	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool :
			fmt.Println("i am boolean")
		case int :
			fmt.Println("i am int")
			default :
			fmt.Printf("i dont know type %T \n",t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}