package main

import (
	"fmt" // format package (text formatting)   from standard library, needs no pathing
	"os"
	"runtime"
	"strconv"
)

func main() {
	myUsername := os.Getenv("USERNAME") // os pkg; get Windows env var

	fmt.Println("Hello, World times", addTwo(2, 3)) // methods have init cap; Println auto-converts int result
	fmt.Println("Go O.S. : ", runtime.GOOS)         // runtime pkg; check the GO Operating System (GOOS)
	fmt.Println("Thank you. My name is", myUsername)
	fmt.Println(playWithString(" the number is: ", 5))
}

func addTwo(a int, b int) int {
	return a + b
}

func playWithString(theString string, theNumber int) string { // post notation of type, string is return type
	return theString + strconv.Itoa(theNumber) // strconv package: integer to ASCII
}
