package lib

import (
	"fmt"
)

// PrintFavorites prints all
func PrintFavorites() {
	for _, v := range favorites {
		fmt.Println(v)
	}
}
