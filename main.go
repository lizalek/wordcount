package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// for _, arg := range os.Args[1:] {
	// 	fmt.Println(arg)
	// }
	str := os.Args[1]

	words := strings.Fields(str)
	count := len(words)
	// for _, word := range words {
	// 	fmt.Print(word)
	// }
	fmt.Println(count)
}
