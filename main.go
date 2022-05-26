package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1]

	words := strings.Fields(str)
	count := len(words)

	fmt.Println(count)
}
