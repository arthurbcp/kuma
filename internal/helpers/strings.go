package helpers

import (
	"strings"
	"unicode"
)

func ToSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i == 0 {
			result = append(result, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

func ToPascalCase(str string) string {
	if len(str) == 0 {
		return str
	}

	return strings.ToUpper(string(str[0])) + str[1:]
}
