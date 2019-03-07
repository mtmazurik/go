package main

import (
	"fmt" // format package (text formatting)   from standard library, needs no pathing
	"os"
	"runtime"
	"strconv"
	"strings"
)

// Person struct is for describing a person - wow that's annoying compiler shit (must comment, must be above the typedef, must begin with Person (the name of the typedef) WTF)
type Person struct {
	name        string
	age         int
	city, phone string
}

func main() { // no input parms, no return values, called automatically
	myUsername := os.Getenv("USERNAME") // os pkg; get Windows env var
	theOS := runtime.GOOS

	fmt.Println("Hello, World times", addTwo(2, 3)) // methods have init cap; Println auto-converts int result
	fmt.Println("Go O.S. : ", theOS)                // runtime pkg; check the GO Operating System (GOOS)
	fmt.Println("Thank you. My name is", myUsername)
	fmt.Println(playWithString(" the number is: ", 5))
	fmt.Println(converter(theOS, myUsername))
	fmt.Println("Top score: ", getBestScore(13, 10, 13, 17, 14, 16))
	fmt.Println("Secret letter: ", testGoCaseStatement(5))
	fmt.Println("Open file result:", openFile("C:\\dummy.txt"))
	fmt.Println("Filling type Person: ", fillPerson())
}

func fillPerson() Person {
	var person Person
	person.name = "Marty"
	person.age = 55
	person.city = "Happy Valley"
	person.phone = "(503)555-1212"
	return person
}

func addTwo(a int, b int) int {
	return a + b
}

func playWithString(theString string, theNumber int) string { // post notation of type, string is return type
	return theString + strconv.Itoa(theNumber) // strconv package: integer to ASCII
}

func converter(os, username string) (retOS, retUsername string, err error) { // multiple return values
	convertedOS := strings.Title(os)
	username = strings.ToUpper(username)
	err = nil
	return convertedOS, username, err
}

func getBestScore(finishes ...int) int { // variadic function (don't know number of variables coming in) uses elipsis ...

	best := finishes[0] // 1st item in the list

	for _, i := range finishes {
		if i > best { // if statement conditional ( ==, !=, <, <=, >, >=, &&, ||)
			best = i
		} //  else if,  else   branches supported
	}
	return best
}
func testGoCaseStatement(index int) string {
	switch index {
	case 1:
		return "z"
	case 2:
		return "x"
	case 3:
		return "e"
	case 4:
		return "l"
	case 5:
		return "q"
	default:
		return "b"
	}
}
func openFile(filename string) error {
	file, err := os.Open(filename)
	file.Close()
	return err
}
