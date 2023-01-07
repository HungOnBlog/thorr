package utils

import "regexp"

func MatchRegex(regex string, valueToMatch string) bool {
	return regexp.MustCompile(regex).MatchString(valueToMatch)
}
