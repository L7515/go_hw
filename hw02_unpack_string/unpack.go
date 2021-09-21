package main

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var prevStack []string
	var decoded []string
	var resString strings.Builder

	for _, val := range input {
		if unicode.IsDigit(val) {
			// also, we could use n, err := strconv.Atoi(string(val)) instead
			n := int(val - '0')

			if len(prevStack) == 0 {
				return "", ErrInvalidString
			}

			switch {
			case n == 0:
				decoded = decoded[:len(decoded)-1]
			case len(decoded) > 1 && decoded[len(decoded)-2] == "\\":
				decoded = decoded[:len(decoded)-2]
			default:
				decoded = decoded[:len(decoded)-1]
			}

			decoded = append(decoded, strings.Repeat(strings.Join(prevStack, ""), n))
			prevStack = nil
		} else {
			if len(prevStack) > 0 && prevStack[len(prevStack)-1] != "\\" {
				prevStack = nil
			}

			prevStack = append(prevStack, string(val))
			decoded = append(decoded, string(val))
		}
	}
	resString.WriteString(strings.Join(decoded, ""))
	return resString.String(), nil
}
