package aaronjson

import (
	"fmt"
	"reflect"
)

type JsonArray interface {
	JsonData
}

type JsonArrayImpl struct {
	jsonNode
	data []JsonData
}

// NewJsonArray creates a new JsonArray instance.
func NewJsonArray() JsonArray {
	return &JsonArrayImpl{
		jsonNode: jsonNode{nodeType: TypeArray},
		data:     make([]JsonData, 0),
	}
}

func (ja *JsonArrayImpl) Index(i int) (JsonData, error) {
	if i < 0 || i >= len(ja.data) {
		return nil, nil // Index out of bounds, return nil
	}
	return ja.data[i], nil
}

func (ja *JsonArrayImpl) SetByIndex(index int, value JsonData) (JsonData, error) {
	if index < 0 || index >= len(ja.data) {
		return nil, nil // Index out of bounds, return nil
	}
	ja.data[index] = value
	return value, nil
}

func (ja *JsonArrayImpl) Append(value JsonData) (JsonData, error) {
	if value == nil {
		return nil, nil // Cannot append nil value
	}
	ja.data = append(ja.data, value)
	return value, nil
}

func (ja *JsonArrayImpl) RemoveByIndex(index int) (JsonData, error) {
	if index < 0 || index >= len(ja.data) {
		return nil, nil // Index out of bounds, return nil
	}
	value := ja.data[index]
	ja.data = append(ja.data[:index], ja.data[index+1:]...)
	return value, nil
}

func (ja *JsonArrayImpl) Length() (int, error) {
	return len(ja.data), nil
}

func (ja *JsonArrayImpl) Unmarshal(v interface{}) error {
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
	case reflect.Slice:
		return ja.unmarshalToSlice(rv)
	case reflect.Array:
		return ja.unmarshalToArray(rv)
	case reflect.Interface:
		// For interface{}, convert to []interface{}
		slice := make([]interface{}, len(ja.data))
		for i, item := range ja.data {
			var elem interface{}
			if err := item.Unmarshal(&elem); err != nil {
				return fmt.Errorf("failed to unmarshal array element at index %d: %v", i, err)
			}
			slice[i] = elem
		}
		rv.Set(reflect.ValueOf(slice))
		return nil
	default:
		return fmt.Errorf("cannot unmarshal array into %v", rv.Type())
	}
}

func (ja *JsonArrayImpl) unmarshalToSlice(rv reflect.Value) error {
	sliceType := rv.Type()
	elemType := sliceType.Elem()
	
	// Create a new slice with the same length as our data
	newSlice := reflect.MakeSlice(sliceType, len(ja.data), len(ja.data))
	
	for i, item := range ja.data {
		elem := reflect.New(elemType)
		if err := item.Unmarshal(elem.Interface()); err != nil {
			return fmt.Errorf("failed to unmarshal array element at index %d: %v", i, err)
		}
		newSlice.Index(i).Set(elem.Elem())
	}
	
	rv.Set(newSlice)
	return nil
}

func (ja *JsonArrayImpl) unmarshalToArray(rv reflect.Value) error {
	arrayType := rv.Type()
	arrayLen := arrayType.Len()
	elemType := arrayType.Elem()
	
	if len(ja.data) > arrayLen {
		return fmt.Errorf("array length mismatch: JSON array has %d elements but target array has capacity %d", len(ja.data), arrayLen)
	}
	
	for i, item := range ja.data {
		elem := reflect.New(elemType)
		if err := item.Unmarshal(elem.Interface()); err != nil {
			return fmt.Errorf("failed to unmarshal array element at index %d: %v", i, err)
		}
		rv.Index(i).Set(elem.Elem())
	}
	
	return nil
}