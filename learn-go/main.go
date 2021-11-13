package main

import (
	"fmt"
	"strings"
)

// Both a and b are int
func multiply(a, b int) int {
	return a * b
}

// Return multiple result
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {

	fmt.Println("Hello World!")

	const name string = "Youngkwon"
	age := 25 // Only possible inside of a func
	// var age int = 25

	fmt.Println(name)
	fmt.Println(age)

	fmt.Println(multiply(2, 3))

	totalLength, upperName := lenAndUpper("Youngkwon")
	fmt.Println(totalLength, upperName)

	repeatMe("youngkwon", "kim", "jake", "alden")
}
