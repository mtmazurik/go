// slices  (and arrays)
// array - like a 'numbered lists' - of the same type (indexed, zero-based); static fixed length (6, or n)
// slices - variable length, flexible (built on top of arrays); they are a reference (pointer); 
//          'slice header' :  name, type, offset, length, passed by reference to functions   ;  length can't be longer than array
//
package main

import (
	"fmt"
)

func main() {
	// Declares a slice called myCourses
	myCourses := make([]string, 5, 10)   // make( <type>, <len>, <cap> )

	fmt.Printf("Length is: %d. \nCapacity is: %d", len(myCourses), cap(myCourses))
}