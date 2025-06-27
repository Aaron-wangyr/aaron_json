package aaronjson

import (
	"fmt"
	"strconv"
)

// ParseByte parses a JSON byte slice and returns the corresponding JsonData.
// If parsing fails, it returns nil and an error.
func ParseByte(jsonData []byte) (JsonData, error) {
	data, err := parseJsonByte(jsonData)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Parse parses a JSON string and returns the corresponding JsonData.
// If parsing fails, it returns nil and an error.
func Parse(jsonStr string) (JsonData, error) {
	data, err := parseJsonByte([]byte(jsonStr))
	if err != nil {
		return nil, err
	}
	return data, nil
}

// parseJsonByte is the internal function that parses JSON byte data.
// It returns the parsed JsonData and any error encountered during parsing.
func parseJsonByte(data []byte) (JsonData, error) {
	pos := 0
	pos = skipWhitespace(data, pos)

	if pos >= len(data) {
		return nil, fmt.Errorf("empty JSON data")
	}

	value, _, err := parseValue(data, pos)
	return value, err
}

// parseValue parses any JSON value starting at the given position.
// It determines the type of JSON value and delegates to the appropriate parser.
// Returns the parsed JsonData, the new position, and any error.
func parseValue(data []byte, pos int) (JsonData, int, error) {
	pos = skipWhitespace(data, pos)

	if pos >= len(data) {
		return nil, pos, fmt.Errorf("unexpected end of data")
	}

	switch data[pos] {
	case '{':
		return parseObject(data, pos)
	case '[':
		return parseArray(data, pos)
	case '"':
		return parseString(data, pos)
	case 't', 'f':
		return parseBool(data, pos)
	case 'n':
		return parseNull(data, pos)
	default:
		if (data[pos] >= '0' && data[pos] <= '9') || data[pos] == '-' || data[pos] == '+' {
			return parseNumber(data, pos)
		}
		return nil, pos, fmt.Errorf("invalid JSON data at position %d", pos)
	}
}

// parseObject parses a JSON object starting at the given position.
// It expects the current character to be '{' and parses key-value pairs
// until it finds the matching '}'. Returns the parsed object, new position, and any error.
func parseObject(data []byte, pos int) (JsonData, int, error) {
	if pos >= len(data) || data[pos] != '{' {
		return nil, pos, fmt.Errorf("expected '{' at position %d", pos)
	}
	pos++ // Move past '{'

	obj := NewJsonObject()
	pos = skipWhitespace(data, pos)

	// Check for empty object
	if pos < len(data) && data[pos] == '}' {
		pos++ // Move past '}'
		return NewJsonObject(), pos, nil
	}

	for {
		pos = skipWhitespace(data, pos)

		if pos >= len(data) {
			return nil, pos, fmt.Errorf("unexpected end of data while parsing object")
		}

		// Parse key
		if data[pos] != '"' {
			return nil, pos, fmt.Errorf("expected string key at position %d", pos)
		}

		keyData, newPos, err := parseString(data, pos)
		if err != nil {
			return nil, pos, err
		}
		pos = newPos

		key := keyData.String()
		if key == "" {
			return nil, pos, fmt.Errorf("empty key at position %d", pos)
		}

		// Expect colon
		pos = skipWhitespace(data, pos)
		if pos >= len(data) || data[pos] != ':' {
			return nil, pos, fmt.Errorf("expected ':' after key at position %d", pos)
		}
		pos++ // Move past ':'

		// Parse value
		pos = skipWhitespace(data, pos)
		value, newPos, err := parseValue(data, pos)
		if err != nil {
			return nil, pos, err
		}
		pos = newPos

		obj.Set(key, value)

		// Check for end or comma
		pos = skipWhitespace(data, pos)
		if pos >= len(data) {
			return nil, pos, fmt.Errorf("unexpected end of data while parsing object")
		}

		if data[pos] == '}' {
			pos++
			break
		} else if data[pos] == ',' {
			pos++
		} else {
			return nil, pos, fmt.Errorf("expected ',' or '}' at position %d", pos)
		}
	}

	return obj, pos, nil
}

// parseArray parses a JSON array starting at the given position.
// It expects the current character to be '[' and parses array elements
// until it finds the matching ']'. Returns the parsed array, new position, and any error.
func parseArray(data []byte, pos int) (JsonData, int, error) {
	if pos >= len(data) || data[pos] != '[' {
		return nil, pos, fmt.Errorf("expected '[' at position %d", pos)
	}
	pos++ // Move past '['

	arr := NewJsonArray()
	pos = skipWhitespace(data, pos)

	// Check for empty array
	if pos < len(data) && data[pos] == ']' {
		pos++ // Move past ']'
		return arr, pos, nil
	}

	for {
		pos = skipWhitespace(data, pos)

		if pos >= len(data) {
			return nil, pos, fmt.Errorf("unexpected end of data while parsing array")
		}

		// Parse element
		value, newPos, err := parseValue(data, pos)
		if err != nil {
			return nil, pos, err
		}
		pos = newPos

		arr.Append(value)

		// Check for end or comma
		pos = skipWhitespace(data, pos)
		if pos >= len(data) {
			return nil, pos, fmt.Errorf("unexpected end of data while parsing array")
		}

		if data[pos] == ']' {
			pos++
			break
		} else if data[pos] == ',' {
			pos++
		} else {
			return nil, pos, fmt.Errorf("expected ',' or ']' at position %d", pos)
		}
	}

	return arr, pos, nil
}

// parseString parses a JSON string starting at the given position.
// It expects the current character to be '"' and parses until the closing '"'.
// Returns the parsed string, new position, and any error.
func parseString(data []byte, pos int) (JsonData, int, error) {
	if pos >= len(data) || data[pos] != '"' {
		return nil, pos, fmt.Errorf("expected '\"' at position %d", pos)
	}
	pos++ // Move past opening '"'

	start := pos
	for pos < len(data) {
		switch data[pos] {
		case '"':
			// Found closing quote
			str := string(data[start:pos])
			pos++ // Move past closing '"'
			return NewJsonString(str), pos, nil
		case '\\':
			// Handle escape sequences (simplified)
			pos += 2
		default:
			pos++
		}
	}

	return nil, pos, fmt.Errorf("unterminated string starting at position %d", start-1)
}

// parseNumber parses a JSON number starting at the given position.
// It handles integers, floats, negative numbers, and scientific notation.
// Returns the parsed number as float64, new position, and any error.
func parseNumber(data []byte, pos int) (JsonData, int, error) {
	start := pos

	// Handle negative sign
	if pos < len(data) && data[pos] == '-' {
		pos++
	}

	// Parse integer part
	if pos >= len(data) || (data[pos] < '0' || data[pos] > '9') {
		return nil, pos, fmt.Errorf("invalid number at position %d", start)
	}

	for pos < len(data) && data[pos] >= '0' && data[pos] <= '9' {
		pos++
	}

	// Parse decimal part
	if pos < len(data) && data[pos] == '.' {
		pos++
		for pos < len(data) && data[pos] >= '0' && data[pos] <= '9' {
			pos++
		}
	}

	// Parse exponent part (simplified)
	if pos < len(data) && (data[pos] == 'e' || data[pos] == 'E') {
		pos++
		if pos < len(data) && (data[pos] == '+' || data[pos] == '-') {
			pos++
		}
		for pos < len(data) && data[pos] >= '0' && data[pos] <= '9' {
			pos++
		}
	}

	// Convert to float64
	numStr := string(data[start:pos])
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return nil, pos, fmt.Errorf("invalid number format: %s", numStr)
	}

	return NewJsonNumber(num), pos, nil
}

// parseBool parses a JSON boolean starting at the given position.
// It expects either "true" or "false" and returns the corresponding boolean value.
// Returns the parsed boolean, new position, and any error.
func parseBool(data []byte, pos int) (JsonData, int, error) {
	if pos+3 < len(data) && string(data[pos:pos+4]) == "true" {
		return NewJsonBool(true), pos + 4, nil
	} else if pos+4 < len(data) && string(data[pos:pos+5]) == "false" {
		return NewJsonBool(false), pos + 5, nil
	}
	return nil, pos, fmt.Errorf("invalid boolean at position %d", pos)
}

// parseNull parses a JSON null value starting at the given position.
// It expects the literal "null" and returns a null JsonData node.
// Returns the null value, new position, and any error.
func parseNull(data []byte, pos int) (JsonData, int, error) {
	if pos+3 < len(data) && string(data[pos:pos+4]) == "null" {
		return NewJsonNull(), pos + 4, nil
	}
	return nil, pos, fmt.Errorf("invalid null at position %d", pos)
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
