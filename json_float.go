package aaronjson

import (
	"fmt"
	"reflect"
)

type JsonFloat struct {
	jsonNode
	data float64
}

// NewJsonFloat creates a new JsonFloat instance.
func NewJsonFloat(value float64) *JsonFloat {
	return &JsonFloat{
		jsonNode: jsonNode{},
		data:     value,
	}
}

func (jn *JsonFloat) IsFloat() bool {
	return true
}

// String returns the string representation of the number.
func (jn *JsonFloat) String() string {
	return fmt.Sprintf("%f", jn.data)
}

// AsInt returns the number as an integer.
func (jn *JsonFloat) AsInt() (int, error) {
	return 0, fmt.Errorf("cannot convert float %f to int", jn.data)
}

// AsFloat returns the number as a float64.
func (jn *JsonFloat) AsFloat() (float64, error) {
	return jn.data, nil
}

// Unmarshal implementation for JsonFloat
func (jn *JsonFloat) Unmarshal(v interface{}) error {
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
		return fmt.Errorf("cannot unmarshal float into %v", rv.Type())
	}
}

// PrettyString returns a pretty-printed JSON number
func (jn *JsonFloat) PrettyString() string {
	// Check if it's an integer value
	if jn.data == float64(int64(jn.data)) && jn.data > -1e12 && jn.data < 1e12 {
		return fmt.Sprintf("%.0f", jn.data)
	}
	
	// Use %g for general formatting - it will choose scientific notation when appropriate
	return fmt.Sprintf("%g", jn.data)
}
