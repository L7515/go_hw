package main

import "fmt"

func main() {
	inputString := "d\\n5abc"

	unpackedStr, err := Unpack(inputString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("unpack: ", unpackedStr)
}
