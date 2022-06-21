package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //we have to follow the same format 01-02-2006

	createDate := time.Date(2018, time.March, 27, 03, 00, 00, 00, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
}
