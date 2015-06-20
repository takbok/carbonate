// The Brahma Utilities package contains commonly used functions for
// manipulating URLs and accessing package name
package utils

import (
	"regexp"
	"runtime"
	"strings"
)

// Get the name of a function till a certain depth
func getFullFunctionName(depth int) string {
	function, _, _, _ := runtime.Caller(depth)
	return runtime.FuncForPC(function).Name()
}

// This method can be used to retrive the name of the source package
// ie, it returns the name of the package from which it was invoked
func GetCallingPackageName() string {
	name := getFullFunctionName(3)

	holder := strings.Split(name, ".")
	holder = strings.Split(holder[0], "/")
	name = holder[len(holder)-1]

	return name
}

// Takes in a CamelCase string and return a hypherated string
// Eg. input - ThisIsASampleString
//     output - this-is-a-sample-string
// Used to map methods to URLs
func Hyphenate(s string) string {
	exp := regexp.MustCompile("([A-Z])")
	s = exp.ReplaceAllString(s, "-$1")
	s = strings.ToLower(s)
	s = strings.Trim(s, "- ")

	return s
}

// Takes a hyphenated strings and returns a CamelCase string
// Eg. input - this-is-a-sample-string
//     output - ThisIsASampleString
// Used to map URLs to methods
func Camelize(s string) string {
	s = strings.ToLower(s)
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "", -1)

	return s
}
