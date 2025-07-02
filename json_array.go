package aaronjson

import (
	"fmt"
	"reflect"
)

type JsonArray struct {
	jsonNode
	data []JsonValue
}

// NewJsonArray creates a new JsonArray instance.
func NewJsonArray() *JsonArray {
	return &JsonArray{
		jsonNode: jsonNode{},
		data:     make([]JsonValue, 0),
	}
}

func (array *JsonArray) AsArray() (*JsonArray, error) {
	return array, nil
}

func (array *JsonArray) IsArray() bool {
	return true
}

func (array *JsonArray) GetSlice() ([]JsonValue, error) {
	if array.data == nil {
		return nil, fmt.Errorf("array is nil")
	}
	return array.data, nil
}

func (array *JsonArray) Index(i int) (JsonValue, error) {
	if i < 0 || i >= len(array.data) {
		return nil, ErrIndexOutOfBounds
	}
	return array.data[i], nil
}

func (array *JsonArray) SetByIndex(index int, value JsonValue) (JsonValue, error) {
	if index < 0 || index >= len(array.data) {
		return nil, ErrIndexOutOfBounds
	}
	array.data[index] = value
	return value, nil
}

func (array *JsonArray) Append(value JsonValue) (JsonValue, error) {
	if value == nil {
		return nil, ErrNilValueAppend
	}
	array.data = append(array.data, value)
	return value, nil
}

func (array *JsonArray) RemoveByIndex(index int) (JsonValue, error) {
	if index < 0 || index >= len(array.data) {
		return nil, ErrIndexOutOfBounds
	}
	value := array.data[index]
	array.data = append(array.data[:index], array.data[index+1:]...)
	return value, nil
}

func (array *JsonArray) Length() (int, error) {
	return len(array.data), nil
}

func (array *JsonArray) Unmarshal(v interface{}) error {
	if v == nil {
		return ErrUnmarshalNilInterface
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return ErrUnmarshalTargetNotPointer
	}

	rv = rv.Elem()
	if !rv.CanSet() {
		return ErrUnmarshalTargetNotSettable
	}
	switch rv.Kind() {
	case reflect.Slice:
		return array.unmarshalToSlice(rv)
	case reflect.Array:
		return array.unmarshalToArray(rv)
	case reflect.Interface:
		// For interface{}, convert to []interface{}
		slice := make([]interface{}, len(array.data))
		for i, item := range array.data {
			var elem interface{}
			if err := item.Unmarshal(&elem); err != nil {
				return fmt.Errorf("failed to unmarshal array element at index %d: %v", i, err)
			}
			slice[i] = elem
		}
		rv.Set(reflect.ValueOf(slice))
		return nil
	default:
		return ErrUnmarshalTargetTypeMismatch
	}
}

func (array *JsonArray) unmarshalToSlice(rv reflect.Value) error {
	sliceType := rv.Type()
	elemType := sliceType.Elem()

	// Create a new slice with the same length as our data
	newSlice := reflect.MakeSlice(sliceType, len(array.data), len(array.data))

	for i, item := range array.data {
		elem := reflect.New(elemType)
		if err := item.Unmarshal(elem.Interface()); err != nil {
			return fmt.Errorf("failed to unmarshal array element at index %d: %v", i, err)
		}
		newSlice.Index(i).Set(elem.Elem())
	}

	rv.Set(newSlice)
	return nil
}

func (array *JsonArray) unmarshalToArray(rv reflect.Value) error {
	arrayType := rv.Type()
	arrayLen := arrayType.Len()
	elemType := arrayType.Elem()

	if len(array.data) > arrayLen {
		return fmt.Errorf("array length mismatch: JSON array has %d elements but target array has capacity %d", len(array.data), arrayLen)
	}

	for i, item := range array.data {
		elem := reflect.New(elemType)
		if err := item.Unmarshal(elem.Interface()); err != nil {
			return fmt.Errorf("failed to unmarshal array element at index %d: %v", i, err)
		}
		rv.Index(i).Set(elem.Elem())
	}

	return nil
}

func (array *JsonArray) String() string {
	if len(array.data) == 0 {
		return "[]"
	}

	result := "["
	for i, item := range array.data {
		if i > 0 {
			result += ", "
		}
		
		// Format the value based on its type
		if _, ok := item.(*JsonString); ok {
			result += fmt.Sprintf("\"%s\"", escapeString(item.String()))
		} else {
			result += item.String()
		}
	}
	result += "]"
	return result
}

// PrettyString returns a pretty-printed JSON array
func (array *JsonArray) PrettyString() string {
	return array.prettyStringWithIndent(0)
}

// prettyStringWithIndent returns a pretty-printed JSON array with specified indentation
func (array *JsonArray) prettyStringWithIndent(indent int) string {
	if len(array.data) == 0 {
		return "[]"
	}

	indentStr := ""
	for i := 0; i < indent; i++ {
		indentStr += "  "
	}
	nextIndentStr := indentStr + "  "

	result := "[\n"
	for i, item := range array.data {
		result += nextIndentStr
		if prettyItem, ok := item.(interface{ prettyStringWithIndent(int) string }); ok {
			result += prettyItem.prettyStringWithIndent(indent + 1)
		} else {
			result += item.PrettyString()
		}
		if i < len(array.data)-1 {
			result += ","
		}
		result += "\n"
	}
	result += indentStr + "]"
	return result
}
