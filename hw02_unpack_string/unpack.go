package main

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func join(strs []string) string {
	var sb strings.Builder

	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}

func Unpack(input string) (string, error) {
	var prevStack []string
	var decoded []string
	var resString strings.Builder

	for _, val := range input {
		n, err := strconv.Atoi(string(val))

		if err == nil {
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

			decoded = append(decoded, strings.Repeat(join(prevStack), n))

			prevStack = nil
		} else {
			if len(prevStack) > 0 && prevStack[len(prevStack)-1] != "\\" {
				prevStack = nil
			}

			prevStack = append(prevStack, string(val))
			decoded = append(decoded, string(val))
		}
	}

	resString.WriteString(join(decoded))

	return resString.String(), nil
}
