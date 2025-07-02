package aaronjson

import (
	"testing"
)

func TestNewJsonObject(t *testing.T) {
	obj := NewJsonObject()
	if obj == nil {
		t.Error("NewJsonObject() returned nil")
	}
	if !obj.IsObject() {
		t.Error("NewJsonObject() should return an object")
	}
	if obj.String() != "{}" {
		t.Errorf("NewJsonObject().String() = %v, want {}", obj.String())
	}
}

func TestJsonObjectSet(t *testing.T) {
	obj := NewJsonObject()
	
	// Test setting string value
	val, err := obj.Set("name", NewJsonString("John"))
	if err != nil {
		t.Errorf("Set() error = %v", err)
	}
	if val == nil {
		t.Error("Set() returned nil value")
	}
	
	// Test setting nil value (should fail)
	_, err = obj.Set("nil_key", nil)
	if err == nil {
		t.Error("Set() should return error when setting nil value")
	}
	
	// Test the object contains the value
	if obj.String() != `{"name": "John"}` {
		t.Errorf("Object string representation = %v, want {\"name\": \"John\"}", obj.String())
	}
}

func TestJsonObjectGet(t *testing.T) {
	obj := NewJsonObject()
	_, _ = obj.Set("name", NewJsonString("John"))
	_, _ = obj.Set("age", NewJsonInt(30))
	
	// Test getting existing key
	val, err := obj.Get("name")
	if err != nil {
		t.Errorf("Get() error = %v", err)
	}
	if val.String() != "John" {
		t.Errorf("Get(\"name\") = %v, want John", val.String())
	}
	
	// Test getting non-existing key
	_, err = obj.Get("nonexistent")
	if err == nil {
		t.Error("Get() should return error for non-existing key")
	}
	
	// Test getting without key
	_, err = obj.Get()
	if err == nil {
		t.Error("Get() should return error when no key provided")
	}
}

func TestJsonObjectRemove(t *testing.T) {
	obj := NewJsonObject()
	_, _ = obj.Set("name", NewJsonString("John"))
	_, _ = obj.Set("age", NewJsonInt(30))
	
	// Test removing existing key
	val, err := obj.Remove("name")
	if err != nil {
		t.Errorf("Remove() error = %v", err)
	}
	if val.String() != "John" {
		t.Errorf("Remove(\"name\") = %v, want John", val.String())
	}
	
	// Test removing non-existing key
	val, err = obj.Remove("nonexistent")
	if err != nil {
		t.Errorf("Remove() error = %v", err)
	}
	if val != nil {
		t.Error("Remove() should return nil for non-existing key")
	}
	
	// Check that key was actually removed
	_, err = obj.Get("name")
	if err == nil {
		t.Error("Key should have been removed")
	}
}

func TestJsonObjectLength(t *testing.T) {
	obj := NewJsonObject()
	
	// Test empty object length
	length, err := obj.Length()
	if err != nil {
		t.Errorf("Length() error = %v", err)
	}
	if length != 0 {
		t.Errorf("Length() = %v, want 0", length)
	}
	
	// Test after adding elements
	_, _ = obj.Set("key1", NewJsonString("value1"))
	_, _ = obj.Set("key2", NewJsonString("value2"))
	
	length, err = obj.Length()
	if err != nil {
		t.Errorf("Length() error = %v", err)
	}
	if length != 2 {
		t.Errorf("Length() = %v, want 2", length)
	}
}

func TestJsonObjectKeys(t *testing.T) {
	obj := NewJsonObject()
	_, _ = obj.Set("zebra", NewJsonString("value1"))
	_, _ = obj.Set("apple", NewJsonString("value2"))
	_, _ = obj.Set("banana", NewJsonString("value3"))
	
	keys, err := obj.Keys()
	if err != nil {
		t.Errorf("Keys() error = %v", err)
	}
	
	// Keys should be sorted
	expected := []string{"apple", "banana", "zebra"}
	if len(keys) != len(expected) {
		t.Errorf("Keys() length = %v, want %v", len(keys), len(expected))
	}
	
	for i, key := range keys {
		if key != expected[i] {
			t.Errorf("Keys()[%d] = %v, want %v", i, key, expected[i])
		}
	}
}

func TestJsonObjectGetMap(t *testing.T) {
	obj := NewJsonObject()
	_, _ = obj.Set("name", NewJsonString("John"))
	_, _ = obj.Set("age", NewJsonInt(30))
	
	// Test getting map from non-empty object
	m, err := obj.GetMap()
	if err != nil {
		t.Errorf("GetMap() error = %v", err)
	}
	if len(m) != 2 {
		t.Errorf("GetMap() length = %v, want 2", len(m))
	}
	
	// Test getting map from empty object
	emptyObj := NewJsonObject()
	_, err = emptyObj.GetMap()
	if err == nil {
		t.Error("GetMap() should return error for empty object")
	}
}

func TestJsonObjectUnmarshal(t *testing.T) {
	obj := NewJsonObject()
	_, _ = obj.Set("name", NewJsonString("John"))
	_, _ = obj.Set("age", NewJsonInt(30))
	_, _ = obj.Set("active", NewJsonBool(true))
	
	// Test unmarshaling to map
	var m map[string]interface{}
	err := obj.Unmarshal(&m)
	if err != nil {
		t.Errorf("Unmarshal() error = %v", err)
	}
	if len(m) != 3 {
		t.Errorf("Unmarshaled map length = %v, want 3", len(m))
	}
	
	// Test unmarshaling to struct
	type Person struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Active bool   `json:"active"`
	}
	
	var p Person
	err = obj.Unmarshal(&p)
	if err != nil {
		t.Errorf("Unmarshal() error = %v", err)
	}
	if p.Name != "John" {
		t.Errorf("Person.Name = %v, want John", p.Name)
	}
	
	// Test unmarshaling to non-pointer (should fail)
	var badTarget Person
	err = obj.Unmarshal(badTarget)
	if err == nil {
		t.Error("Unmarshal() should return error for non-pointer target")
	}
	
	// Test unmarshaling to nil (should fail)
	err = obj.Unmarshal(nil)
	if err == nil {
		t.Error("Unmarshal() should return error for nil target")
	}
}

func TestJsonObjectPrettyString(t *testing.T) {
	obj := NewJsonObject()
	_, _ = obj.Set("name", NewJsonString("John"))
	_, _ = obj.Set("age", NewJsonInt(30))
	
	pretty := obj.PrettyString()
	if pretty == "" {
		t.Error("PrettyString() returned empty string")
	}
	
	// Test empty object pretty string
	emptyObj := NewJsonObject()
	if emptyObj.PrettyString() != "{}" {
		t.Errorf("Empty object PrettyString() = %v, want {}", emptyObj.PrettyString())
	}
}

func TestJsonObjectNestedGet(t *testing.T) {
	// Create nested object structure
	inner := NewJsonObject()
	_, _ = inner.Set("value", NewJsonString("nested"))
	
	outer := NewJsonObject()
	_, _ = outer.Set("inner", inner)
	
	// Test nested get
	val, err := outer.Get("inner", "value")
	if err != nil {
		t.Errorf("Get() error = %v", err)
	}
	if val.String() != "nested" {
		t.Errorf("Get(\"inner\", \"value\") = %v, want nested", val.String())
	}
}
