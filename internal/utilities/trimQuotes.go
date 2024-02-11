package utilities

import "strings"

func TrimQuotes(s string) string {
	return strings.Trim(s, "\"")
}
