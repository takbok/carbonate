package utils

import (
	"regexp"
	"strings"
)

func Hyphenate(s string) string {
	exp := regexp.MustCompile("([A-Z])")
	s = exp.ReplaceAllString(s, "-$1")
	s = strings.ToLower(s)
	s = strings.Trim(s, "- ")

	return s
}

func Camelize(s string) string {
	s = strings.ToLower(s)
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)

	return s
}
