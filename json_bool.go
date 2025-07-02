package aaronjson

import (
	"fmt"
	"reflect"
)

var JSON_BOOL_TRUE = NewJsonBool(true)
var JSON_BOOL_FALSE = NewJsonBool(false)

type JsonBool struct {
	jsonNode
	data bool
}

// NewJsonBool creates a new JsonBool instance.
func NewJsonBool(value bool) *JsonBool {
	return &JsonBool{
		jsonNode: jsonNode{},
		data:     value,
	}
}

func (jb *JsonBool) IsBool() bool {
	return true
}

func (jb *JsonBool) AsBool() (bool, error) {
	return jb.data, nil
}

func (jb *JsonBool) AsString() (string, error) {
	if jb.data {
		return "true", nil
	}
	return "false", nil
}

func (jb *JsonBool) AsInt() (int, error) {
	if jb.data {
		return 1, nil
	}
	return 0, nil
}

func (jb *JsonBool) AsFloat() (float64, error) {
	if jb.data {
		return 1.0, nil
	}
	return 0.0, nil
}

// String returns the string representation of the boolean.
func (jb *JsonBool) String() string {
	if jb.data {
		return "true"
	}
	return "false"
}

// PrettyString returns a pretty-printed JSON boolean
func (jb *JsonBool) PrettyString() string {
	if jb.data {
		return "true"
	}
	return "false"
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
