package aaronjson

import "fmt"

func parseJsonByte(data []byte) (JsonData, error) {
	pos := 0
	pos = skipWhitespace(data, pos)
	var value JsonData
	var err error
	if pos < len(data) {
		switch data[pos] {
		case '{':
			value = &JsonObject{}
		case '[':
			value = &JsonArray{}
		case '"':
			value = &JsonString{}
		case 't', 'f':
			if data[pos] == 't' && pos+3 < len(data) && data[pos+1] == 'r' && data[pos+2] == 'u' && data[pos+3] == 'e' {
				value = &JsonBool{*JSONTrue, true}
				pos += 3 // Move past "true"
			} else if data[pos] == 'f' && pos+4 < len(data) && data[pos+1] == 'a' && data[pos+2] == 'l' && data[pos+3] == 's' && data[pos+4] == 'e' {
				value = &JsonBool{*JSONFalse, false}
				pos += 4 // Move past "false"
			} else {
				return nil, fmt.Errorf("Invalid JSON boolean at position %d: %s", pos, string(data[pos:]))
			}
			pos++ // Move past the boolean character
		case 'n':
			value = &JsonNull{}
		default:
			if (data[pos] >= '0' && data[pos] <= '9') || data[pos] == '-' || data[pos] == '+' {
				value = &JsonNumber{}
			} else {
				return nil, fmt.Errorf("Invalid JSON data at position %d: %s", pos, string(data[pos:]))
			}
		}
		value, pos, err = value.parse(data, pos)
		if err != nil {
			return nil, fmt.Errorf("Error parsing JSON data at position %d: %v", pos, err)
		}
	} else {
		return nil, fmt.Errorf("Empty JSON data")
	}
}

func skipWhitespace(data []byte, pos int) int {
	for pos < len(data) && (data[pos] == ' ' || data[pos] == '\n' || data[pos] == '\r' || data[pos] == '\t') {
		pos++
	}
	return pos
}
