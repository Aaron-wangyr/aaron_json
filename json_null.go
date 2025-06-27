package aaronjson

import (
	"fmt"
	"reflect"
)

type JsonNull struct {
	jsonNode
}

// NewJsonNull creates a new JsonNull instance.
func NewJsonNull() *JsonNull {
	return &JsonNull{
		jsonNode: jsonNode{nodeType: TypeNull},
	}
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

	// For null values, we set the target to its zero value or nil for pointers/interfaces
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		rv.Set(reflect.Zero(rv.Type()))
		return nil
	default:
		// For non-nullable types, set to zero value
		rv.Set(reflect.Zero(rv.Type()))
		return nil
	}
}

// UnmarshalTo is an alias for Unmarshal
func (jn *JsonNull) UnmarshalTo(v interface{}) error {
	return jn.Unmarshal(v)
}

// PrettyString returns a pretty-printed JSON null
func (jn *JsonNull) PrettyString() string {
	return "null"
}