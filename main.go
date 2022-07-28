package main

// fmt allow us to print on screen
//import "fmt"
import (
	"fmt"
)

// There are 2 important point
// Its important to have a package name called main, this way go binary knows this is an executable file
// And its must have func main(), coz its an entry point of our executable application
// This way go compiler knows its an executable file

func printmessage() string {
	return "Hello Go World From Method !"
}

func main() {
	fmt.Println("Hello Go World !!")
	fmt.Println(printmessage())
}
