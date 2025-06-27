package aaronjson

import (
	"fmt"
	"reflect"
)

type JsonNumber struct {
	jsonNode
	data float64
}

// NewJsonNumber creates a new JsonNumber instance.
func NewJsonNumber(value float64) *JsonNumber {
	return &JsonNumber{
		jsonNode: jsonNode{nodeType: TypeNumber},
		data:     value,
	}
}

// String returns the string representation of the number.
func (jn *JsonNumber) String() string {
	return fmt.Sprintf("%f", jn.data)
}

// Type returns the type of the JSON data.
func (jn *JsonNumber) Type() JsonType {
	return TypeNumber
}

// AsInt returns the number as an integer.
func (jn *JsonNumber) AsInt() (int, error) {
	return int(jn.data), nil
}

// AsFloat returns the number as a float64.
func (jn *JsonNumber) AsFloat() (float64, error) {
	return jn.data, nil
}

// Unmarshal implementation for JsonNumber
func (jn *JsonNumber) Unmarshal(v interface{}) error {
	if v == nil {
		return fmt.Errorf("cannot unmarshal into nil interface")
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("unmarshal target must be a pointer")
	}

	rv = rv.Elem()
	if !rv.CanSet() {
		return fmt.Errorf("unmarshal target cannot be set")
	}

	switch rv.Kind() {
	case reflect.Float32, reflect.Float64:
		rv.SetFloat(jn.data)
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		rv.SetInt(int64(jn.data))
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		rv.SetUint(uint64(jn.data))
		return nil
	case reflect.Interface:
		rv.Set(reflect.ValueOf(jn.data))
		return nil
	default:
		return fmt.Errorf("cannot unmarshal number into %v", rv.Type())
	}
}

// UnmarshalTo is an alias for Unmarshal
func (jn *JsonNumber) UnmarshalTo(v interface{}) error {
	return jn.Unmarshal(v)
}

// PrettyString returns a pretty-printed JSON number
func (jn *JsonNumber) PrettyString() string {
	// Check if it's an integer value
	if jn.data == float64(int64(jn.data)) {
		return fmt.Sprintf("%.0f", jn.data)
	}
	return fmt.Sprintf("%g", jn.data)
}