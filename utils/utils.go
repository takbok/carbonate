package utils

import (
	"regexp"
	"runtime"
	"strings"
)

func getFullFunctionName(depth int) string {
	function, _, _, _ := runtime.Caller(depth)
	return runtime.FuncForPC(function).Name()
}

func GetCallingPackageName() string {
	name := getFullFunctionName(3)

	holder := strings.Split(name, ".")
	holder = strings.Split(holder[0], "/")
	name = holder[len(holder)-1]

	return name
}

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
