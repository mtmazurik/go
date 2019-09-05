// maps are unordered lists; key-value pairs (like dictionary or hash-tables)
//      are by-reference
//      keys must be unique
package main

import (
	"fmt"
)

func main() {

	mydict := make(map[string]string) // [key]value

	mydict["marty@martymazurik.com"] = "Marty Mazurik"

	fmt.Printf("Lookup map value: %s\n", mydict["marty@martymazurik.com"])
	fmt.Printf("Length is: %d. \n", len(mydict))

}
