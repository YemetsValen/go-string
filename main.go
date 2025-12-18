package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func unpackString(s string) (string, error) {
	var result []rune
	i := 0

	for i < len(s) {
		r := rune(s[i])

		if r == '\\' {
			i++
			if i >= len(s) {
				return "", fmt.Errorf("некорректная строка: экранирование в конце строки")
			}
			r = rune(s[i])
		} else {
			if unicode.IsDigit(r) {
				if i == 0 || unicode.IsDigit(rune(s[i-1])) {
					return "", fmt.Errorf("некорректная строка: число в начале или подряд идущие цифры")
				}
			}
		}

		if i+1 < len(s) && unicode.IsDigit(rune(s[i+1])) {
			numStr := ""
			j := i + 1
			for j < len(s) && unicode.IsDigit(rune(s[j])) {
				numStr += string(s[j])
				j++
			}
			count, _ := strconv.Atoi(numStr)
			for k := 0; k < count; k++ {
				result = append(result, r)
			}
			i = j
			continue
		}

		result = append(result, r)
		i++
	}

	return string(result), nil
}
