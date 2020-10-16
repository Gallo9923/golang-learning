package main

import (
	"fmt"

	"github.com/Gallo9923/GoRecipies/lib"
	"github.com/Gallo9923/GoRecipies/strutils"
)

func main() {
	useLib()
}

func useLib() {
	fmt.Println("******* Default favorite packages *******")
	lib.PrintFavorites()

	//Add couple of favorites
	lib.Add("Go")
	lib.Add("Language")
	fmt.Println("******* Updated *******")
	lib.PrintFavorites()

	fmt.Println("Size of favorites: ", len(lib.GetAll()))

}

func useStrutils() {
	fmt.Println("Hello, World!")
	str1, str2 := "chris", "GALLO"

	// Convert to uppercase
	fmt.Println("To Upper Case: ", strutils.ToUpperCase(str1))

	// Convert to lowercase
	fmt.Println("To Lower Case:", strutils.ToLowerCase(str2))

	// Convert first letter to uppercase
	fmt.Println("To First Upper", strutils.ToFirstUpper(str1))
}
