package main

import (
	"fmt"
)

func main() {
	// Declares a slice called myCourses
	//myCourses := make([]string, 5, 10) // make( <type>, <len>, <capacity> )
	myslice := []string{"Marty", "Michelle", "Marie"}

	fmt.Printf("Length is: %d. \nCapacity is: %d\n", len(myslice), cap(myslice))

	fmt.Printf("Sibling: %s\n", myslice[1]) // index zero-based

	//myslice[3] = "Russell"               // nope - go.panic exception!    beyond end of initial (literal-defined array)
	myslice = append(myslice, "Russell") // fix, use append if you know you are going beyond the cap[acity]
	fmt.Printf("Length is: %d. \nCapacity is: %d\n", len(myslice), cap(myslice))

	fmt.Printf("Sibling: %s\n", myslice[3])

	//myslice[4] = "Brett"
	myslice = append(myslice, "Brett")
	fmt.Printf("Length is: %d. \nCapacity is: %d\n", len(myslice), cap(myslice))

	fmt.Printf("Brother-from-another-mutha: %s\n", myslice[4])

	myslice = append(myslice, "Ron")
	fmt.Printf("Length is: %d. \nCapacity is: %d\n", len(myslice), cap(myslice))

	fmt.Printf("Brother-from-another-mutha: %s\n", myslice[5])

	myslice = append(myslice, "John")
	fmt.Printf("Length is: %d. \nCapacity is: %d\n", len(myslice), cap(myslice))

	fmt.Printf("Brother-from-another-mutha: %s\n", myslice[6])

}

// slices  (and arrays)
// array - like a 'numbered lists' - of the same type (indexed, zero-based); static fixed length (6, or n)
// slices - variable length, flexible (built on top of arrays); they are a reference (pointer);
//          'slice header' :  name, type, offset, length, passed by reference to functions   ;  length can't be longer than array
//
