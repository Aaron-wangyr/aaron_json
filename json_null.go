package aaronjson

import (
	"fmt"
	"reflect"
)

var JSON_NULL = NewJsonNull()

type JsonNull struct {
	jsonNode
}

// NewJsonNull creates a new JsonNull instance.
func NewJsonNull() *JsonNull {
	return &JsonNull{
		jsonNode: jsonNode{},
	}
}

func (jn *JsonNull) IsNull() bool {
	return true
}

// String returns the string representation of the JSON null.
func (jn *JsonNull) String() string {
	return "null"
}

// PrettyString returns a pretty-printed JSON null
func (jn *JsonNull) PrettyString() string {
	return "null"
}

// Unmarshal implementation for JsonNull
func (jn *JsonNull) Unmarshal(v interface{}) error {
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
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map:
		rv.Set(reflect.Zero(rv.Type()))
		return nil
	default:
		return fmt.Errorf("cannot unmarshal null into %v", rv.Type())
	}
}
