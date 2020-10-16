package main

import (
	"fmt"

	"github.com/Gallo9923/strutils"
)

func main() {
	fmt.Println("Hello, World!")
	str1, str2 := "chris", "GALLO"

	// Convert to uppercase
	fmt.Println("To Upper Case: ", strutils.ToUpperCase(str1))

	// Convert to lowercase
	fmt.Println("To Lower Case:", strutils.ToLowerCase(str2))

	// Convert first letter to uppercase
	fmt.Println("To First Upper", strutils.ToFirstUpper(str1))
}
