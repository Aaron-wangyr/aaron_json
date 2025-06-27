package aaronjson

import (
	"fmt"
	"reflect"
)

type JsonBool struct {
	jsonNode
	data bool
}

// NewJsonBool creates a new JsonBool instance.
func NewJsonBool(value bool) *JsonBool {
	return &JsonBool{
		jsonNode: jsonNode{nodeType: TypeBool},
		data:     value,
	}
}

// String returns the string representation of the boolean.
func (jb *JsonBool) String() string {
	if jb.data {
		return "true"
	}
	return "false"
}

// Type returns the type of the JSON data.
func (jb *JsonBool) Type() JsonType {
	return TypeBool
}

// AsBool returns the boolean value.
func (jb *JsonBool) AsBool() (bool, error) {
	return jb.data, nil
}

// Unmarshal implementation for JsonBool
func (jb *JsonBool) Unmarshal(v interface{}) error {
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
	case reflect.Bool:
		rv.SetBool(jb.data)
		return nil
	case reflect.Interface:
		rv.Set(reflect.ValueOf(jb.data))
		return nil
	default:
		return fmt.Errorf("cannot unmarshal bool into %v", rv.Type())
	}
}

// UnmarshalTo is an alias for Unmarshal
func (jb *JsonBool) UnmarshalTo(v interface{}) error {
	return jb.Unmarshal(v)
}

// PrettyString returns a pretty-printed JSON boolean
func (jb *JsonBool) PrettyString() string {
	if jb.data {
		return "true"
	}
	return "false"
}