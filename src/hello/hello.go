package main

import (
	"fmt" // format package (text formatting)   from standard library, needs no pathing
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Person struct is for describing a person - wow that's annoying compiler shit (must comment, must be above the typedef, must begin with Person (the name of the typedef) WTF)
type Person struct {
	Name        string
	Age         int
	City, Phone string
}

func main() { // no input parms, no return values, called automatically
	myUsername := os.Getenv("USERNAME") // os pkg; get Windows env var
	theOS := runtime.GOOS

	runtime.GOMAXPROCS(2) // increases the parallel thread pool count to two (2), goroutines use a thread pool
	var waitGrp sync.WaitGroup
	waitGrp.Add(1) // part of concurency, had to add the "sync" pkg above, too

	//todo: NYI, implement:  myChannel := make(chan int, 5) // channel, buffer size(5)(optional) channel writer blocks until read occurs, unless buffered like this one at 5, it can buffer 5

	// slice variable length, can be resized like .net lists; slices are built on top of arrays, slices are references to contiguous sections of an array (name, type, starting offset, length)
	mySlice := make([]string, 5, 10)                                                                         // slice example  type, length, capacity (cap -max)
	initializedSlice := []string{"Humpty", "Dumpty", "sat on a wall", "Humpty", "Dumpty", "took a big fall"} // an initialized array == a slice automatically
	sliceOfSlice := initializedSlice[3:6]
	sliceOfSlice = append(sliceOfSlice, "WTF humpty") // dynamically append to a slice, is it effecient?

	// maps : list of unordered, key-value pairs  map[keyType]valueType, keys MUST BE unique
	myDictionary := make(map[string]string)
	myDictionary["Name"] = "Marty"
	myDictionary["SSN"] = "777-88-9999"
	// to create a "real hash table" based on unique hashes of strings, in Go, see https://flaviocopes.com/golang-data-structure-hashtable/

	// concurency
	go func() { // the go keyword makes it a goroutine (concurrent)
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Awake after 5 sec snooze.")
	}() // these parent make the function 'self-executing' using concurrent threads, doled out by the go runtime

	// beginning of output section
	fmt.Println("Hello, World times", addTwo(2, 3)) // methods have init cap; Println auto-converts int result
	fmt.Println("Go O.S. : ", theOS)                // runtime pkg; check the GO Operating System (GOOS)
	fmt.Println("Thank you. My name is", myUsername)
	fmt.Println(playWithString(" the number is: ", 5))
	fmt.Println(converter(theOS, myUsername))
	fmt.Println("Top score: ", getBestScore(13, 10, 13, 17, 14, 16))
	fmt.Println("Secret letter: ", testGoCaseStatement(5))
	fmt.Println("Open file result:", openFile("C:\\dummy.txt"))
	fmt.Println("Filling type Person: ", fillPerson())
	fmt.Printf("Slice1 len: %d Cap: %d \n", len(mySlice), cap(mySlice))
	fmt.Printf("Slice2 len: %d Cap: %d \n", len(initializedSlice), cap(initializedSlice))
	fmt.Printf("SliceOfSlice values: %s %s %s %s\n", sliceOfSlice[0], sliceOfSlice[1], sliceOfSlice[2], sliceOfSlice[3]) // [3] is the appended one
	fmt.Printf("Slice2 len: %d Cap: %d \n", len(sliceOfSlice), cap(sliceOfSlice))                                        //  the capacity is THE SAME as the original array, append() didn't grow it, just filled a cell
	fmt.Println("Println dump sliceOfSlice: ", sliceOfSlice)
	fmt.Println("Maps", myDictionary)

	waitGrp.Wait() // syncronize with all concurrent goroutines completion event
	fmt.Println("Exiting, all goroutines in WaitGroup complete.")
}

func fillPerson() Person { // initializing a struct 3-ways
	var person Person // declaring a struct
	// or
	// person := new(Person) // this one gives us a ptr
	// or
	// person := Person{Name:"Marty" , Age:55, City:"Happy Valley", Phone:"(503)555-1212"}

	person.Name = "Marty"
	person.Age = 55
	person.City = "Happy Valley"
	person.Phone = "(503)555-1212"
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

	for a, i := range finishes { // _ in place of 'a', would be an 'undefined variable', in this case the index position; I print 'a' to show its value
		if i > best { // if statement conditional ( ==, !=, <, <=, >, >=, &&, ||)
			best = i
			// break or continue can be used in for loop
		} //  else if,  else   branches supported
		fmt.Println("a: ", a)
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
