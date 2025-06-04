package aaronjson

import "fmt"


func Marshal(obj interface{}) JsonObject {
	return nil
}

func Parse(data []byte) (JsonObject, error) {
	var val JsonObject
	if len(data) == 0 {
		return nil, fmt.Errorf(ErrEmptyData)
	}
	trimedData,err := trimData(data)
	if err != nil {
		return nil, err
	}
	if trimedData[0] == '{' {
		val = &JsonMap{}
		if len(trimedData) == 2 && trimedData[1] == '}' {
			// Special case for empty object
			return val, nil
		} 
		if err := val.parse(trimedData); err != nil {
			return nil, fmt.Errorf(ErrInvalidJsonFormat, err.Error())
		}
		
		

	} else if trimedData[0] == '[' {
		val = &JsonArray{}
		if len(trimedData) == 2 && trimedData[1] == ']' {
			// Special case for empty array
			return val, nil
		}
	} else {
		return nil, fmt.Errorf(ErrInvalidJsonFormat, fmt.Sprintf("expected '{' or '[', got '%c'", trimedData[0]))
	}
	return nil, nil
}

func trimData(data []byte) ([]byte,error) {
	start := 0
	for start < len(data) && (data[start] == ' ' || data[start] == '\n' || data[start] == '\t') {
		start++
	}
	if start == len(data) {
		return nil, fmt.Errorf(ErrEmptyData)
	}
	startChar := data[start]
	if startChar != '{' && startChar != '[' {
		return nil, fmt.Errorf(ErrInvalidJsonFormat, fmt.Sprintf("expected '{' or '[', got '%c'", startChar))
	}
	end := len(data) - 1
	for end >= start && (data[end] == ' ' || data[end] == '\n' || data[end] == '\t') {
		end--
	}
	if end < start {
		return nil, fmt.Errorf(ErrInvalidJsonFormat, "no valid JSON content found")
	}
	if data[end] != '}' && data[end] != ']' {
		return nil, fmt.Errorf(ErrInvalidJsonFormat, fmt.Sprintf("expected '}' or ']', got '%c'", data[end]))
	}
	if startChar == '{' && data[end] != '}' {
		return nil, fmt.Errorf(ErrInvalidJsonFormat, fmt.Sprintf("expected '}' for object, got '%c'", data[end]))
	}
	if startChar == '[' && data[end] != ']' {
		return nil, fmt.Errorf(ErrInvalidJsonFormat, fmt.Sprintf("expected ']' for array, got '%c'", data[end]))
	}
	return data[start : end+1], nil
}