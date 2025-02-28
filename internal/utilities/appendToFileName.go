package utilities

import "strings"

func AppendToFile(filename string, content string) string {
	pathParts := strings.Split(filename, ".")

	return pathParts[0] + content + "." + pathParts[1]
}
