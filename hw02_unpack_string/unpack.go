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
	var resString strings.Builder

	for _, val := range input {
		decoded := string(val)
		n, err := strconv.Atoi(string(val))

		if err == nil {
			if len(prevStack) == 0 {
				return "", ErrInvalidString
			}

			repeatCount := n - 1
			if repeatCount < 0 {
				repeatCount = 0
			}

			decoded = strings.Repeat(join(prevStack), repeatCount)

			prevStack = nil
		} else {
			if len(prevStack) > 0 && prevStack[len(prevStack)-1] != "\\" {
				prevStack = nil
			}
			prevStack = append(prevStack, string(val))
		}

		if decoded != "" {
			resString.WriteString(decoded)
		}
	}

	return resString.String(), nil
}
