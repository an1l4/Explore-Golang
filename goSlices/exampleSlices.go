package main

import "fmt"

func main() {
	s := make([]string,3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println(s[2])
	fmt.Println(len(s))

	s = append(s,"d")
	s = append(s,"e","f")
	fmt.Println(s)

	c := make([]string,len(s))
	copy(c,s)
	fmt.Println(c)

	l := s[2:5]
	fmt.Println(l)

	l = s[:5]
	fmt.Println(l)

	l = s[2:]
	fmt.Println(l)

	t := []string{"i","j","k"}
	fmt.Println(t)

	twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
		//fmt.Println(i)
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
			//fmt.Println(j)
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}