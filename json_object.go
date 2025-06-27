package aaronjson

import (
	"fmt"
	"reflect"
	"sort"
)

type JsonObject interface {
	JsonData
}

type JsonObjectImpl struct {
	jsonNode
	data       map[string]JsonData
	sortedkeys []string
}

// Constructor for JsonObject
func NewJsonObject() JsonObject {
	keys := make([]string, 0)
	return &JsonObjectImpl{
		jsonNode:   jsonNode{nodeType: TypeObject},
		data:       make(map[string]JsonData),
		sortedkeys: keys,
	}
}

func (jo *JsonObjectImpl) Get(key string) (JsonData, error) {
	if value, exists := jo.data[key]; exists {
		return value, nil
	}
	return nil, fmt.Errorf("key '%s' not found in object", key)
}

func (jo *JsonObjectImpl) Set(key string, value JsonData) (JsonData, error) {
	if value == nil {
		return nil, fmt.Errorf("cannot set nil value for key '%s'", key)
	}
	jo.data[key] = value
	jo.updateKeys()
	return value, nil
}

func (jo *JsonObjectImpl) Remove(key string) (JsonData, error) {
	if value, exists := jo.data[key]; exists {
		delete(jo.data, key)
		jo.updateKeys()
		return value, nil
	}
	return nil, nil // Key not found, return nil
}

func (jo *JsonObjectImpl) Length() (int, error) {
	return len(jo.data), nil
}

func (jo *JsonObjectImpl) Keys() ([]string, error) {
	return jo.sortedkeys, nil
}

func (jo *JsonObjectImpl) updateKeys() {
	jo.sortedkeys = make([]string, 0, len(jo.data))
	for key := range jo.data {
		jo.sortedkeys = append(jo.sortedkeys, key)
	}
	sort.Strings(jo.sortedkeys) // Keep keys sorted
}

func (jo *JsonObjectImpl) Unmarshal(v interface{}) error {
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

func (jo *JsonObjectImpl) unmarshalToMap(rv reflect.Value) error {
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

func (jo *JsonObjectImpl) unmarshalToStruct(rv reflect.Value) error {
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
