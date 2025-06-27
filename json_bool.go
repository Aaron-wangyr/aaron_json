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