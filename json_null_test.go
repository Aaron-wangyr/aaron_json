package aaronjson

import (
	"testing"
)

func TestNewJsonNull(t *testing.T) {
	null := NewJsonNull()
	if null == nil {
		t.Error("NewJsonNull() returned nil")
	}
	if !null.IsNull() {
		t.Error("NewJsonNull() should return a null")
	}
}

func TestJsonNullGlobalConstant(t *testing.T) {
	// Test global constant
	if JSON_NULL == nil {
		t.Error("JSON_NULL should not be nil")
	}

	if JSON_NULL.String() != "null" {
		t.Errorf("JSON_NULL.String() = %v, want null", JSON_NULL.String())
	}

	if !JSON_NULL.IsNull() {
		t.Error("JSON_NULL.IsNull() should return true")
	}
}

func TestJsonNullString(t *testing.T) {
	null := NewJsonNull()
	got := null.String()
	want := "null"
	if got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

func TestJsonNullPrettyString(t *testing.T) {
	null := NewJsonNull()
	got := null.PrettyString()
	want := "null"
	if got != want {
		t.Errorf("PrettyString() = %v, want %v", got, want)
	}
}

func TestJsonNullTypeChecks(t *testing.T) {
	null := NewJsonNull()

	// Test type checks
	if !null.IsNull() {
		t.Error("IsNull() should return true")
	}
	if null.IsString() {
		t.Error("IsString() should return false")
	}
	if null.IsInt() {
		t.Error("IsInt() should return false")
	}
	if null.IsFloat() {
		t.Error("IsFloat() should return false")
	}
	if null.IsBool() {
		t.Error("IsBool() should return false")
	}
	if null.IsArray() {
		t.Error("IsArray() should return false")
	}
	if null.IsObject() {
		t.Error("IsObject() should return false")
	}
}

func TestJsonNullAsString(t *testing.T) {
	null := NewJsonNull()
	_, err := null.AsString()
	if err == nil {
		t.Error("AsString() should return error for null")
	}
}

func TestJsonNullAsInt(t *testing.T) {
	null := NewJsonNull()
	_, err := null.AsInt()
	if err == nil {
		t.Error("AsInt() should return error for null")
	}
}

func TestJsonNullAsFloat(t *testing.T) {
	null := NewJsonNull()
	_, err := null.AsFloat()
	if err == nil {
		t.Error("AsFloat() should return error for null")
	}
}

func TestJsonNullAsBool(t *testing.T) {
	null := NewJsonNull()
	_, err := null.AsBool()
	if err == nil {
		t.Error("AsBool() should return error for null")
	}
}

func TestJsonNullAsObject(t *testing.T) {
	null := NewJsonNull()
	_, err := null.AsObject()
	if err == nil {
		t.Error("AsObject() should return error for null")
	}
}

func TestJsonNullAsArray(t *testing.T) {
	null := NewJsonNull()
	_, err := null.AsArray()
	if err == nil {
		t.Error("AsArray() should return error for null")
	}
}

func TestJsonNullGet(t *testing.T) {
	null := NewJsonNull()
	_, err := null.Get("key")
	if err == nil {
		t.Error("Get() should return error for null")
	}
}

func TestJsonNullGetMap(t *testing.T) {
	null := NewJsonNull()
	_, err := null.GetMap()
	if err == nil {
		t.Error("GetMap() should return error for null")
	}
}

func TestJsonNullGetSlice(t *testing.T) {
	null := NewJsonNull()
	_, err := null.GetSlice()
	if err == nil {
		t.Error("GetSlice() should return error for null")
	}
}

func TestJsonNullUnmarshal(t *testing.T) {
	null := NewJsonNull()

	// Test unmarshaling to pointer - should set to nil
	stringPtr := new(string)
	*stringPtr = "test"

	err := null.Unmarshal(&stringPtr)
	if err != nil {
		t.Errorf("Unmarshal() error = %v", err)
	}
	if stringPtr != nil {
		t.Error("Unmarshal() should set pointer to nil")
	}

	// Test unmarshaling to interface{} - should set to nil
	var interfaceTarget interface{} = "test"

	err = null.Unmarshal(&interfaceTarget)
	if err != nil {
		t.Errorf("Unmarshal() to interface{} error = %v", err)
	}
	if interfaceTarget != nil {
		t.Error("Unmarshal() should set interface{} to nil")
	}

	// Test unmarshaling to non-pointer value (should fail)
	var badTarget string
	err = null.Unmarshal(badTarget)
	if err == nil {
		t.Error("Unmarshal() should return error for non-pointer target")
	}

	// Test unmarshaling to nil (should fail)
	err = null.Unmarshal(nil)
	if err == nil {
		t.Error("Unmarshal() should return error for nil target")
	}
}
