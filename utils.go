package aaronjson

import "fmt"

// escapeString escapes special characters in JSON strings
func escapeString(s string) string {
	result := ""
	for _, r := range s {
		switch r {
		case '"':
			result += "\\\""
		case '\\':
			result += "\\\\"
		case '\b':
			result += "\\b"
		case '\f':
			result += "\\f"
		case '\n':
			result += "\\n"
		case '\r':
			result += "\\r"
		case '\t':
			result += "\\t"
		default:
			if r < 32 {
				result += fmt.Sprintf("\\u%04x", r)
			} else {
				result += string(r)
			}
		}
	}
	return result
}
