package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	inputStr := "Hello, OTUS!"
	reversedStr := stringutil.Reverse(inputStr)

	fmt.Println(reversedStr)
}
