package aaronjson

import (
	"fmt"
	"reflect"
)

// Marshal converts a Go value to JsonValue.
// It supports structs, maps, slices, arrays, and basic types.
func Marshal(v interface{}) (JsonValue, error) {
	if v == nil {
		return NewJsonNull(), nil
	}

	return marshalValue(reflect.ValueOf(v))
}

// marshalValue is the internal function that handles the conversion
// of reflect.Value to JsonValue based on the value's type.
func marshalValue(rv reflect.Value) (JsonValue, error) {
	// Handle pointers by dereferencing them
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return NewJsonNull(), nil
		}
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Bool:
		return NewJsonBool(rv.Bool()), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return NewJsonInt(float64(rv.Int())), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return NewJsonInt(float64(rv.Uint())), nil

	case reflect.Float32, reflect.Float64:
		return NewJsonFloat(rv.Float()), nil

	case reflect.String:
		return NewJsonString(rv.String()), nil

	case reflect.Slice, reflect.Array:
		return marshalSlice(rv)

	case reflect.Map:
		return marshalMap(rv)

	case reflect.Struct:
		return marshalStruct(rv)

	case reflect.Interface:
		if rv.IsNil() {
			return NewJsonNull(), nil
		}
		return marshalValue(rv.Elem())

	default:
		return nil, fmt.Errorf("unsupported type: %v", rv.Type())
	}
}

// marshalSlice converts a slice or array to JsonArray
func marshalSlice(rv reflect.Value) (JsonValue, error) {
	arr := NewJsonArray()

	for i := 0; i < rv.Len(); i++ {
		elem, err := marshalValue(rv.Index(i))
		if err != nil {
			return nil, fmt.Errorf("failed to marshal array element at index %d: %v", i, err)
		}
		_, _ = arr.Append(elem)
	}

	return arr, nil
}

// marshalMap converts a map to JsonObject
func marshalMap(rv reflect.Value) (JsonValue, error) {
	// Only support maps with string keys
	if rv.Type().Key().Kind() != reflect.String {
		return nil, fmt.Errorf("only maps with string keys are supported")
	}

	obj := NewJsonObject()

	for _, key := range rv.MapKeys() {
		keyStr := key.String()
		value := rv.MapIndex(key)

		jsonValue, err := marshalValue(value)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal map value for key '%s': %v", keyStr, err)
		}

		_, _ = obj.Set(keyStr, jsonValue)
	}

	return obj, nil
}

// marshalStruct converts a struct to JsonObject
func marshalStruct(rv reflect.Value) (JsonValue, error) {
	obj := NewJsonObject()
	structType := rv.Type()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldValue := rv.Field(i)

		// Skip unexported fields
		if !fieldValue.CanInterface() {
			continue
		}

		// Get the JSON field name from struct tag or use field name
		jsonFieldName := getJsonFieldName(field)
		if jsonFieldName == "-" {
			continue // Skip fields marked with json:"-"
		}

		// Check for omitempty tag
		omitEmpty := hasOmitEmptyTag(field)
		if omitEmpty && isEmptyValue(fieldValue) {
			continue
		}

		jsonValue, err := marshalValue(fieldValue)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal struct field '%s': %v", field.Name, err)
		}

		_, _ = obj.Set(jsonFieldName, jsonValue)
	}

	return obj, nil
}

// getJsonFieldName extracts the JSON field name from struct tag
func getJsonFieldName(field reflect.StructField) string {
	tag := field.Tag.Get("json")
	if tag == "" {
		return field.Name
	}

	// Handle json:",omitempty" format
	for i, c := range tag {
		if c == ',' {
			return tag[:i]
		}
	}

	return tag
}

// hasOmitEmptyTag checks if the field has the omitempty tag
func hasOmitEmptyTag(field reflect.StructField) bool {
	tag := field.Tag.Get("json")
	return containsOmitEmpty(tag)
}

// containsOmitEmpty checks if the tag contains "omitempty"
func containsOmitEmpty(tag string) bool {
	// Check for exact match
	if tag == "omitempty" {
		return true
	}

	// Check for ",omitempty" pattern
	omitEmptyTag := ",omitempty"
	if len(tag) >= len(omitEmptyTag) {
		for i := 0; i <= len(tag)-len(omitEmptyTag); i++ {
			if tag[i:i+len(omitEmptyTag)] == omitEmptyTag {
				return true
			}
		}
	}

	return false
}

// isEmptyValue checks if a value is considered empty for omitempty
func isEmptyValue(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return rv.IsNil()
	}
	return false
}
