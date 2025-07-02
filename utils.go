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

// skipWhitespace advances the position past any whitespace characters.
// It skips spaces, tabs, newlines, and carriage returns.
// Returns the new position after skipping whitespace.
func skipWhitespace(data []byte, pos int) int {
	for pos < len(data) && (data[pos] == ' ' || data[pos] == '\n' || data[pos] == '\r' || data[pos] == '\t') {
		pos++
	}
	return pos
}

// isValidNumberTerminator checks if a character can validly follow a number
func isValidNumberTerminator(c byte) bool {
	return c == ',' || c == ']' || c == '}' || c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

// isValidBoolTerminator checks if a character can validly follow a boolean
func isValidBoolTerminator(c byte) bool {
	return c == ',' || c == ']' || c == '}' || c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

// isValidNullTerminator checks if a character can validly follow a null
func isValidNullTerminator(c byte) bool {
	return c == ',' || c == ']' || c == '}' || c == ' ' || c == '\t' || c == '\n' || c == '\r'
}