package main

import "fmt"

func avg(x float64, y float64) float64 {
	return (x + y) / 2
}

func main() {
	var name string = " Hello Tim"
	var age int = 21
	var uint32 gjennomS
	fmt.Println(" Hello world!", name, "and your age is: ", age) //" " is used for string, ' ' is used for char
	gjennomS = avg(6.56, 13.44)
	fmt.Println("Avg value is: ", gjennomS)
}

/*
multiline comment
*/
//run code: go run fileName.go
// make executable file: go build fileName.go
