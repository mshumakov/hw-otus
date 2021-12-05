package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")
var builder strings.Builder

func Unpack(s string) (string, error) {
	var lastChar string
	var lastRune rune

	for il1, rl1 := range s {
		if (unicode.IsDigit(rl1) && il1 == 0) || (unicode.IsDigit(lastRune) && unicode.IsDigit(rl1)) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(rl1) {
			if rl1 == 48 { // Проверка на нуль =D
				tmpS := builder.String()
				builder.Reset()

				builder.WriteString(tmpS[:len(tmpS)-1])
				continue
			}

			if count, err := strconv.Atoi(string(rl1)); err == nil {
				builder.WriteString(strings.Repeat(lastChar, count-1))

				lastChar = string(rl1)
				lastRune = rl1

				continue
			}
		}

		builder.WriteString(strings.Repeat(string(rl1), 1))

		lastChar = string(rl1)
		lastRune = rl1
	}

	result := builder.String()
	builder.Reset()

	return result, nil
}
