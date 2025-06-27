package aaronjson

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type JsonString struct {
	jsonNode
	data string
}

// Constructor
func NewJsonString(value string) *JsonString {
	return &JsonString{
		jsonNode: jsonNode{nodeType: TypeString},
		data:     value,
	}
}

// JsonValue interface methods
func (js *JsonString) String() string {
	return js.data
}

func (js *JsonString) Type() JsonType {
	return TypeString
}

// JsonData interface methods - Get/Index operations
func (js *JsonString) Get(key string) (JsonData, error) {
	return nil, errors.New("cannot get property from string")
}

func (js *JsonString) Index(i int) (JsonData, error) {
	return nil, errors.New("cannot index string")
}

// Type conversion methods
func (js *JsonString) AsString() (string, error) {
	return js.data, nil
}

func (js *JsonString) AsInt() (int, error) {
	if i, err := strconv.Atoi(js.data); err == nil {
		return i, nil
	}
	return 0, fmt.Errorf("cannot convert '%s' to int", js.data)
}

func (js *JsonString) AsFloat() (float64, error) {
	if f, err := strconv.ParseFloat(js.data, 64); err == nil {
		return f, nil
	}
	return 0, fmt.Errorf("cannot convert '%s' to float", js.data)
}

func (js *JsonString) AsBool() (bool, error) {
	if b, err := strconv.ParseBool(js.data); err == nil {
		return b, nil
	}
	return false, fmt.Errorf("cannot convert '%s' to bool", js.data)
}

// Collection methods
func (js *JsonString) Length() (int, error) {
	return len(js.data), nil
}

// Unmarshal implementation for JsonString
func (js *JsonString) Unmarshal(v interface{}) error {
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
	case reflect.String:
		rv.SetString(js.data)
		return nil
	case reflect.Interface:
		rv.Set(reflect.ValueOf(js.data))
		return nil
	default:
		return fmt.Errorf("cannot unmarshal string into %v", rv.Type())
	}
}

// UnmarshalTo is an alias for Unmarshal
func (js *JsonString) UnmarshalTo(v interface{}) error {
	return js.Unmarshal(v)
}

// PrettyString returns a pretty-printed JSON string with proper escaping
func (js *JsonString) PrettyString() string {
	return fmt.Sprintf("\"%s\"", escapeString(js.data))
}
