package main

import (
	"fmt"     // format package (text formatting)   from standard library, needs no pathing
	"runtime" //
)

func main() {
<<<<<<< HEAD
	fmt.Printf("Hello, World\n")
}
=======
	fmt.Println("Hello, World times", addTwo(2, 3)) // initial cap;  expose Println method "outside of the package", implies l.c. would be internal-only
	fmt.Println("Go O.S. : ", runtime.GOOS)
}

func addTwo(a int, b int) int {
	return a + b
}
>>>>>>> updated helloworld
