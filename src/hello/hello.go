package main

import (
	"fmt" // format package (text formatting)   from standard library, needs no pathing
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() { // no input parms, no return values, called automatically
	myUsername := os.Getenv("USERNAME") // os pkg; get Windows env var
	theOS := runtime.GOOS

	fmt.Println("Hello, World times", addTwo(2, 3)) // methods have init cap; Println auto-converts int result
	fmt.Println("Go O.S. : ", theOS)                // runtime pkg; check the GO Operating System (GOOS)
	fmt.Println("Thank you. My name is", myUsername)
	fmt.Println(playWithString(" the number is: ", 5))
	fmt.Println(converter(theOS, myUsername))
	fmt.Println("Top score: ", getBestScore(13, 10, 13, 17, 14, 16))
}

func addTwo(a int, b int) int {
	return a + b
}

func playWithString(theString string, theNumber int) string { // post notation of type, string is return type
	return theString + strconv.Itoa(theNumber) // strconv package: integer to ASCII
}

func converter(os, username string) (retOS, retUsername string) { // multiple return values
	convertedOS := strings.Title(os)
	username = strings.ToUpper(username)
	return convertedOS, username
}

func getBestScore(finishes ...int) int { // variadic function (don't know number of variables coming in) uses elipsis ...

	best := finishes[0] // 1st item in the list

	for _, i := range finishes {
		if i > best {
			best = i
		}
	}
	return best
}
