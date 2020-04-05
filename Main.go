package main

import (
	"fmt"
	"strconv"
)

var j int = 42 // scoped at the package
// var J int = 25 // scoped globally, exported at the package

var (
	name   string = "Simon"
	number int    = 30
)

func main() {
	// var i int
	// i = 42
	fmt.Println(j)
	// fmt.Println(J)
	var j float32 = 26 // block scope
	k := 99            // cant make it float32, has to be float64
	fmt.Printf("%v, %T", j, j)
	fmt.Printf("%v, %T\n", k, k)

	// YOu can also convert types
	var someNumber float32 = 42.5
	var intNUmber int = int(someNumber) // You will lose the .5 info

	var someString string
	someString = string(intNUmber)
	fmt.Printf("%v, %T\n", someString, someString) // converting number to string, looks for unicode 42. u need strconv package
	someString = strconv.Itoa(intNUmber)           // Itoa stands for Int To Ascii
	fmt.Printf("%v, %T\n", someString, someString) // converting number to string, looks for unicode 42. u need strconv package

	// DAY 2: https://youtu.be/YS4e4q9oBaU?t=3442
}
