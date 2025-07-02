package aaronjson

import (
	"fmt"
	"strconv"
)

// ParseByte parses a JSON byte slice and returns the corresponding JsonValue.
// If parsing fails, it returns nil and an error.
func ParseByte(jsonData []byte) (JsonValue, error) {
	return parseJsonByte(jsonData)
}

// Parse parses a JSON string and returns the corresponding JsonValue.
// If parsing fails, it returns nil and an error.
func Parse(jsonStr string) (JsonValue, error) {
	data, err := parseJsonByte([]byte(jsonStr))
	if err != nil {
		return nil, err
	}
	return data, nil
}

// parseJsonByte is the internal function that parses JSON byte data.
// It returns the parsed JsonValue and any error encountered during parsing.
func parseJsonByte(data []byte) (JsonValue, error) {
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
// Returns the parsed JsonValue, the new position, and any error.
func parseValue(data []byte, pos int) (JsonValue, int, error) {
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
func parseObject(data []byte, pos int) (JsonValue, int, error) {
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

		// Get the raw string value for the key (not the JSON-formatted string)
		key, err := keyData.AsString()
		if err != nil {
			return nil, pos, fmt.Errorf("failed to get string value for key: %v", err)
		}
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

		_, _ = obj.Set(key, value)

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
func parseArray(data []byte, pos int) (JsonValue, int, error) {
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

		_, _ = arr.Append(value)

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
func parseString(data []byte, pos int) (JsonValue, int, error) {
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
// Returns the parsed number as JsonValue, new position, and any error.
func parseNumber(data []byte, pos int) (JsonValue, int, error) {
	start := pos

	// Check for empty data
	if pos >= len(data) {
		return nil, pos, fmt.Errorf("unexpected end of data while parsing number")
	}

	// Handle optional minus sign
	switch data[pos] {
	case '-':
		pos++
		if pos >= len(data) {
			return nil, pos, fmt.Errorf("invalid number: minus sign without digits at position %d", start)
		}
	case '+':
		// JSON doesn't allow leading plus signs
		return nil, pos, fmt.Errorf("invalid number: leading plus sign at position %d", start)
	}

	// Check if we have at least one digit
	if pos >= len(data) || (data[pos] < '0' || data[pos] > '9') {
		return nil, pos, fmt.Errorf("invalid number: no digits at position %d", start)
	}

	isFloat := false
	hasExponent := false

	// Parse integer part
	if data[pos] == '0' {
		// Leading zero - next character must not be a digit (unless it's a decimal point or exponent)
		pos++
		if pos < len(data) && data[pos] >= '0' && data[pos] <= '9' {
			return nil, pos, fmt.Errorf("invalid number: leading zero followed by digit at position %d", start)
		}
	} else {
		// Parse digits (1-9 followed by 0-9*)
		for pos < len(data) && data[pos] >= '0' && data[pos] <= '9' {
			pos++
		}
	}

	// Check for decimal point
	if pos < len(data) && data[pos] == '.' {
		isFloat = true
		pos++

		// Must have at least one digit after decimal point
		if pos >= len(data) || data[pos] < '0' || data[pos] > '9' {
			return nil, pos, fmt.Errorf("invalid number: decimal point without fractional digits at position %d", start)
		}

		// Parse fractional digits
		for pos < len(data) && data[pos] >= '0' && data[pos] <= '9' {
			pos++
		}
	}

	// Check for exponent
	if pos < len(data) && (data[pos] == 'e' || data[pos] == 'E') {
		isFloat = true
		hasExponent = true
		pos++

		// Optional sign after exponent
		if pos < len(data) && (data[pos] == '+' || data[pos] == '-') {
			pos++
		}

		// Must have at least one digit after exponent
		if pos >= len(data) || data[pos] < '0' || data[pos] > '9' {
			return nil, pos, fmt.Errorf("invalid number: exponent without digits at position %d", start)
		}

		// Parse exponent digits
		for pos < len(data) && data[pos] >= '0' && data[pos] <= '9' {
			pos++
		}
	}

	// Extract the number string
	numStr := string(data[start:pos])

	// Validate that we stopped at an appropriate character
	if pos < len(data) {
		c := data[pos]
		if !isValidNumberTerminator(c) {
			return nil, pos, fmt.Errorf("invalid character '%c' after number at position %d", c, pos)
		}
	}

	// Parse the number based on type
	if isFloat || hasExponent {
		// Parse as float
		val, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return nil, pos, fmt.Errorf("invalid float number '%s' at position %d: %v", numStr, start, err)
		}

		// Check for overflow/underflow
		if val == 0 && numStr != "0" && numStr != "-0" && numStr != "0.0" && numStr != "-0.0" {
			// Check if it's actually zero or underflow
			hasNonZero := false
			for i := 0; i < len(numStr); i++ {
				if numStr[i] >= '1' && numStr[i] <= '9' {
					hasNonZero = true
					break
				}
			}
			if hasNonZero {
				return nil, pos, fmt.Errorf("number underflow: '%s' at position %d", numStr, start)
			}
		}

		return NewJsonFloat(val), pos, nil
	} else {
		// Parse as integer
		val, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return nil, pos, fmt.Errorf("invalid integer number '%s' at position %d: %v", numStr, start, err)
		}
		return NewJsonInt(val), pos, nil
	}
}

func parseBool(data []byte, pos int) (JsonValue, int, error) {
	if pos+4 <= len(data) && string(data[pos:pos+4]) == "true" {
		newPos := pos + 4
		// Check that we're at a valid word boundary
		if newPos < len(data) && !isValidBoolTerminator(data[newPos]) {
			return nil, pos, fmt.Errorf("invalid boolean value at position %d", pos)
		}
		return NewJsonBool(true), newPos, nil
	} else if pos+5 <= len(data) && string(data[pos:pos+5]) == "false" {
		newPos := pos + 5
		// Check that we're at a valid word boundary
		if newPos < len(data) && !isValidBoolTerminator(data[newPos]) {
			return nil, pos, fmt.Errorf("invalid boolean value at position %d", pos)
		}
		return NewJsonBool(false), newPos, nil
	}
	return nil, pos, fmt.Errorf("invalid boolean value at position %d", pos)
}

func parseNull(data []byte, pos int) (JsonValue, int, error) {
	if pos+4 <= len(data) && string(data[pos:pos+4]) == "null" {
		newPos := pos + 4
		// Check that we're at a valid word boundary
		if newPos < len(data) && !isValidNullTerminator(data[newPos]) {
			return nil, pos, fmt.Errorf("invalid null value at position %d", pos)
		}
		return NewJsonNull(), newPos, nil
	}
	return nil, pos, fmt.Errorf("invalid null value at position %d", pos)
}
