package main

import (
	"strings"
)

func main() {
	println(isValid("({{{}}})"))
}
func isValid(s string) bool {
	c := ""
	x := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "[" {
			c += "]"
			x++
		} else if string(s[i]) == "(" {
			c += ")"
			x++
		} else if string(s[i]) == "{" {
			c += "}"
			x++
		} else if x > 0 && string(c[x-1]) == string(s[i]) {
			x--
			c = strings.TrimSuffix(c, string(c[x]))
		} else {
			return false
		}

	}
	if x == 0 && len(s) > 0 {
		return true
	}
	return false
}
