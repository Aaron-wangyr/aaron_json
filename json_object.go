package aaronjson

import (
	"fmt"
	"reflect"
	"sort"
)

type JsonObject struct {
	jsonNode
	data       map[string]JsonValue
	sortedkeys []string
}

func NewJsonObject() *JsonObject {
	keys := make([]string, 0)
	return &JsonObject{
		jsonNode:   jsonNode{},
		data:       make(map[string]JsonValue),
		sortedkeys: keys,
	}
}

func (jo *JsonObject) AsObject() (*JsonObject, error) {
	return jo, nil
}

func (jo *JsonObject) IsObject() bool {
	return true
}

func (jo *JsonObject) Get(key ...string) (JsonValue, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("no key provided for Get operation")
	}
	
	if len(key) == 1 {
		// Single key, return the value directly
		if value, exists := jo.data[key[0]]; exists {
			return value, nil
		}
		return nil, fmt.Errorf("key '%s' not found in object", key[0])
	}
	
	// Multiple keys, navigate through nested objects
	if value, exists := jo.data[key[0]]; exists {
		if !value.IsObject() {
			return nil, fmt.Errorf("value for key '%s' is not an object, cannot match key path", key[0])
		}
		return value.Get(key[1:]...)
	}
	
	return nil, fmt.Errorf("key '%s' not found in object", key[0])
}

func (jo *JsonObject) GetMap() (map[string]JsonValue, error) {
	if len(jo.data) == 0 {
		return nil, fmt.Errorf("object is empty, cannot return map")
	}
	result := make(map[string]JsonValue, len(jo.data))
	for key, value := range jo.data {
		if value == nil {
			return nil, fmt.Errorf("value for key '%s' is nil, cannot return map", key)
		}
		result[key] = value
	}
	return result, nil
}

func (jo *JsonObject) Set(key string, value JsonValue) (JsonValue, error) {
	if value == nil {
		return nil, fmt.Errorf("cannot set nil value for key '%s'", key)
	}
	jo.data[key] = value
	jo.updateKeys()
	return value, nil
}

func (jo *JsonObject) Remove(key string) (JsonValue, error) {
	if value, exists := jo.data[key]; exists {
		delete(jo.data, key)
		jo.updateKeys()
		return value, nil
	}
	return nil, nil // Key not found, return nil
}

func (jo *JsonObject) Length() (int, error) {
	return len(jo.data), nil
}

func (jo *JsonObject) Keys() ([]string, error) {
	return jo.sortedkeys, nil
}

func (jo *JsonObject) updateKeys() {
	jo.sortedkeys = make([]string, 0, len(jo.data))
	for key := range jo.data {
		jo.sortedkeys = append(jo.sortedkeys, key)
	}
	sort.Strings(jo.sortedkeys) // Keep keys sorted
}

func (jo *JsonObject) String() string {
	if len(jo.data) == 0 {
		return "{}"
	}

	result := "{"
	for i, key := range jo.sortedkeys {
		if i > 0 {
			result += ", "
		}
		value := jo.data[key]
		
		// Format the value based on its type
		var valueStr string
		if _, ok := value.(*JsonString); ok {
			valueStr = fmt.Sprintf("\"%s\"", escapeString(value.String()))
		} else {
			valueStr = value.String()
		}
		
		result += fmt.Sprintf("%q: %s", key, valueStr)
	}
	result += "}"
	return result
}

func (jo *JsonObject) Unmarshal(v interface{}) error {
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
	case reflect.Map:
		return jo.unmarshalToMap(rv)
	case reflect.Struct:
		return jo.unmarshalToStruct(rv)
	case reflect.Interface:
		// For interface{}, convert to map[string]interface{}
		result := make(map[string]interface{})
		for key, value := range jo.data {
			var elem interface{}
			if err := value.Unmarshal(&elem); err != nil {
				return fmt.Errorf("failed to unmarshal object field '%s': %v", key, err)
			}
			result[key] = elem
		}
		rv.Set(reflect.ValueOf(result))
		return nil
	default:
		return fmt.Errorf("cannot unmarshal object into %v", rv.Type())
	}
}

func (jo *JsonObject) unmarshalToMap(rv reflect.Value) error {
	mapType := rv.Type()
	keyType := mapType.Key()
	valueType := mapType.Elem()

	// Check if key type is string
	if keyType.Kind() != reflect.String {
		return fmt.Errorf("map key type must be string, got %v", keyType)
	}

	// Create a new map
	newMap := reflect.MakeMap(mapType)

	for key, value := range jo.data {
		mapKey := reflect.ValueOf(key)
		mapValue := reflect.New(valueType)

		if err := value.Unmarshal(mapValue.Interface()); err != nil {
			return fmt.Errorf("failed to unmarshal object field '%s': %v", key, err)
		}

		newMap.SetMapIndex(mapKey, mapValue.Elem())
	}

	rv.Set(newMap)
	return nil
}

func (jo *JsonObject) unmarshalToStruct(rv reflect.Value) error {
	structType := rv.Type()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldValue := rv.Field(i)

		if !fieldValue.CanSet() {
			continue // Skip unexported fields
		}

		// Get the JSON field name (use struct tag if available, otherwise field name)
		jsonFieldName := field.Tag.Get("json")
		if jsonFieldName == "" || jsonFieldName == "-" {
			jsonFieldName = field.Name
		}

		// Handle json:",omitempty" and other tag options
		for commaIndex := 0; commaIndex < len(jsonFieldName); commaIndex++ {
			if jsonFieldName[commaIndex] == ',' {
				jsonFieldName = jsonFieldName[:commaIndex]
				break
			}
		}

		if jsonFieldName == "" {
			jsonFieldName = field.Name
		}

		// Get the value from JSON object
		if jsonValue, exists := jo.data[jsonFieldName]; exists {
			fieldPtr := reflect.New(field.Type)
			if err := jsonValue.Unmarshal(fieldPtr.Interface()); err != nil {
				return fmt.Errorf("failed to unmarshal field '%s': %v", field.Name, err)
			}
			fieldValue.Set(fieldPtr.Elem())
		}
	}

	return nil
}

// PrettyString returns a pretty-printed JSON object
func (jo *JsonObject) PrettyString() string {
	return jo.prettyStringWithIndent(0)
}

// prettyStringWithIndent returns a pretty-printed JSON object with specified indentation
func (jo *JsonObject) prettyStringWithIndent(indent int) string {
	if len(jo.data) == 0 {
		return "{}"
	}

	indentStr := ""
	for i := 0; i < indent; i++ {
		indentStr += "  "
	}
	nextIndentStr := indentStr + "  "

	result := "{\n"
	keys, _ := jo.Keys()
	for i, key := range keys {
		result += nextIndentStr
		result += fmt.Sprintf("\"%s\": ", escapeString(key))

		value := jo.data[key]
		if prettyItem, ok := value.(interface{ prettyStringWithIndent(int) string }); ok {
			result += prettyItem.prettyStringWithIndent(indent + 1)
		} else {
			result += value.PrettyString()
		}

		if i < len(keys)-1 {
			result += ","
		}
		result += "\n"
	}
	result += indentStr + "}"
	return result
}
