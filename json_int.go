package aaronjson

import (
	"fmt"
	"reflect"
)

type JsonInt struct {
	jsonNode
	data float64
}

// NewJsonInt creates a new JsonInt instance.
func NewJsonInt(value float64) *JsonInt {
	return &JsonInt{
		jsonNode: jsonNode{},
		data:     value,
	}
}

func (jn *JsonInt) IsInt() bool {
	return true
}

// String returns the string representation of the number.
func (jn *JsonInt) String() string {
	return fmt.Sprintf("%f", jn.data)
}

// AsInt returns the number as an integer.
func (jn *JsonInt) AsInt() (int, error) {
	return int(jn.data), nil
}

// AsInt returns the number as a float64.
func (jn *JsonInt) AsFloat() (float64, error) {
	return 0, fmt.Errorf("cannot convert int %f to float64", jn.data)
}

// Unmarshal implementation for JsonInt
func (jn *JsonInt) Unmarshal(v interface{}) error {
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
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		rv.SetInt(int64(jn.data))
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		rv.SetUint(uint64(jn.data))
		return nil
	case reflect.Float32, reflect.Float64:
		rv.SetFloat(jn.data)
		return nil
	case reflect.Interface:
		rv.Set(reflect.ValueOf(int(jn.data)))
		return nil
	default:
		return fmt.Errorf("cannot unmarshal int into %v", rv.Type())
	}
}

// PrettyString returns a pretty-printed JSON number
func (jn *JsonInt) PrettyString() string {
	// Check if it's an integer value
	if jn.data == float64(int64(jn.data)) {
		return fmt.Sprintf("%.0f", jn.data)
	}
	return fmt.Sprintf("%g", jn.data)
}
