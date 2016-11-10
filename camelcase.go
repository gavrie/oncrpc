package main

import "unicode"

func camelCase(s string, upperInitial bool) string {
	camel := make([]rune, 0, len(s))
	upperNext := upperInitial

	for _, r := range s {
		if r == '_' {
			upperNext = true
			continue
		}
		if upperNext {
			r = unicode.ToUpper(r)
			upperNext = false
		} else {
			r = unicode.ToLower(r)
		}
		camel = append(camel, r)
		if !unicode.IsLetter(r) {
			upperNext = true
		}
	}

	return string(camel)
}
