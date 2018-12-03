package stringsutils

import "strings"

// ReplaceMsg replace a set of placeholders with values
func ReplaceMsg(message string, placeHolders []string, values []string) string {
	for i, k := range placeHolders {
		message = strings.Replace(message, k, values[i], -1)
	}
	return message
}
