package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world!")

	myVar, err := allAboutFuncs(29, 12, "hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(myVar)
}

// functions
func allAboutFuncs(x, y int, something string) (int, error) {
	fmt.Println(x+y, something)
	// if x+y != 1000 {
	// 	return 0, fmt.Errorf("Oh no error occured")
	// }
	return 200, nil
}

// variables
func allAboutVars() {
	var myVar string
	myVar = "hello"
	fmt.Println(myVar)

	x := 1234   // int
	f := 1.1234 // 64 bit floating point number

	var someNum int

	myRawString = `hello
world`

	myAscii := '%' // prints out a the ASCII number, not treated as a string / char

	myByte := byte(6)

	// Arrays
	myArr := [3]int{1, 2, 3}
	fmt.Println(myArr[0])

	// Slices (have a dynamic size as compared to arrays)
	mySlice := []int{1, 2, 3}
	mySlice = append(mySlice, 4, 5, 6, 7, 8)

	myOtherSlice := make([]int, 5)
	// will look like {0,0,0,0,0}

	// Maps
	// "myKey" ---> "myValue"
	myMap := map[string]string{"myFirstKey": "myFirstValue"}
	myMap["myKey"] = "myValue"
}
