package utils

import (
	"regexp"
	"strings"
)

func SplitOnUppercase(str string) string {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	result := re.FindAllString(str, -1)
	return strings.Join(result, " ")
}

func UppercaseFirstLetter(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToUpper(string(str[0])) + str[1:]
}
