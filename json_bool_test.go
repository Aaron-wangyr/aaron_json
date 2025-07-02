package aaronjson

import (
	"testing"
)

func TestNewJsonBool(t *testing.T) {
	boolTrue := NewJsonBool(true)
	if boolTrue == nil {
		t.Error("NewJsonBool(true) returned nil")
	}
	if !boolTrue.IsBool() {
		t.Error("NewJsonBool(true) should return a bool")
	}

	boolFalse := NewJsonBool(false)
	if boolFalse == nil {
		t.Error("NewJsonBool(false) returned nil")
	}
	if !boolFalse.IsBool() {
		t.Error("NewJsonBool(false) should return a bool")
	}
}

func TestJsonBoolGlobalConstants(t *testing.T) {
	// Test global constants
	if JSON_BOOL_TRUE == nil {
		t.Error("JSON_BOOL_TRUE should not be nil")
	}
	if JSON_BOOL_FALSE == nil {
		t.Error("JSON_BOOL_FALSE should not be nil")
	}
	
	if JSON_BOOL_TRUE.String() != "true" {
		t.Errorf("JSON_BOOL_TRUE.String() = %v, want true", JSON_BOOL_TRUE.String())
	}
	if JSON_BOOL_FALSE.String() != "false" {
		t.Errorf("JSON_BOOL_FALSE.String() = %v, want false", JSON_BOOL_FALSE.String())
	}
}

func TestJsonBoolString(t *testing.T) {
	tests := []struct {
		name  string
		value bool
		want  string
	}{
		{
			name:  "true value",
			value: true,
			want:  "true",
		},
		{
			name:  "false value",
			value: false,
			want:  "false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewJsonBool(tt.value)
			got := b.String()
			if got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonBoolPrettyString(t *testing.T) {
	tests := []struct {
		name  string
		value bool
		want  string
	}{
		{
			name:  "true value",
			value: true,
			want:  "true",
		},
		{
			name:  "false value",
			value: false,
			want:  "false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewJsonBool(tt.value)
			got := b.PrettyString()
			if got != tt.want {
				t.Errorf("PrettyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonBoolTypeChecks(t *testing.T) {
	b := NewJsonBool(true)
	
	// Test type checks
	if !b.IsBool() {
		t.Error("IsBool() should return true")
	}
	if b.IsString() {
		t.Error("IsString() should return false")
	}
	if b.IsInt() {
		t.Error("IsInt() should return false")
	}
	if b.IsFloat() {
		t.Error("IsFloat() should return false")
	}
	if b.IsArray() {
		t.Error("IsArray() should return false")
	}
	if b.IsObject() {
		t.Error("IsObject() should return false")
	}
	if b.IsNull() {
		t.Error("IsNull() should return false")
	}
}

func TestJsonBoolAsBool(t *testing.T) {
	// Test true
	boolTrue := NewJsonBool(true)
	val, err := boolTrue.AsBool()
	if err != nil {
		t.Errorf("AsBool() error for true = %v", err)
	}
	if val != true {
		t.Errorf("AsBool() for true = %v, want true", val)
	}
	
	// Test false
	boolFalse := NewJsonBool(false)
	val, err = boolFalse.AsBool()
	if err != nil {
		t.Errorf("AsBool() error for false = %v", err)
	}
	if val != false {
		t.Errorf("AsBool() for false = %v, want false", val)
	}
}

func TestJsonBoolAsString(t *testing.T) {
	// Test true
	boolTrue := NewJsonBool(true)
	val, err := boolTrue.AsString()
	if err != nil {
		t.Errorf("AsString() error for true = %v", err)
	}
	if val != "true" {
		t.Errorf("AsString() for true = %v, want true", val)
	}
	
	// Test false
	boolFalse := NewJsonBool(false)
	val, err = boolFalse.AsString()
	if err != nil {
		t.Errorf("AsString() error for false = %v", err)
	}
	if val != "false" {
		t.Errorf("AsString() for false = %v, want false", val)
	}
}

func TestJsonBoolAsInt(t *testing.T) {
	// Test true
	boolTrue := NewJsonBool(true)
	val, err := boolTrue.AsInt()
	if err != nil {
		t.Errorf("AsInt() error for true = %v", err)
	}
	if val != 1 {
		t.Errorf("AsInt() for true = %v, want 1", val)
	}
	
	// Test false
	boolFalse := NewJsonBool(false)
	val, err = boolFalse.AsInt()
	if err != nil {
		t.Errorf("AsInt() error for false = %v", err)
	}
	if val != 0 {
		t.Errorf("AsInt() for false = %v, want 0", val)
	}
}

func TestJsonBoolAsFloat(t *testing.T) {
	// Test true
	boolTrue := NewJsonBool(true)
	val, err := boolTrue.AsFloat()
	if err != nil {
		t.Errorf("AsFloat() error for true = %v", err)
	}
	if val != 1.0 {
		t.Errorf("AsFloat() for true = %v, want 1.0", val)
	}
	
	// Test false
	boolFalse := NewJsonBool(false)
	val, err = boolFalse.AsFloat()
	if err != nil {
		t.Errorf("AsFloat() error for false = %v", err)
	}
	if val != 0.0 {
		t.Errorf("AsFloat() for false = %v, want 0.0", val)
	}
}

func TestJsonBoolUnmarshal(t *testing.T) {
	boolTrue := NewJsonBool(true)
	
	// Test unmarshaling to bool
	var target bool
	err := boolTrue.Unmarshal(&target)
	if err != nil {
		t.Errorf("Unmarshal() error = %v", err)
	}
	if target != true {
		t.Errorf("Unmarshal() result = %v, want true", target)
	}
	
	// Test unmarshaling to interface{}
	var interfaceTarget interface{}
	err = boolTrue.Unmarshal(&interfaceTarget)
	if err != nil {
		t.Errorf("Unmarshal() to interface{} error = %v", err)
	}
	if interfaceTarget != true {
		t.Errorf("Unmarshal() to interface{} = %v, want true", interfaceTarget)
	}
	
	// Test unmarshaling to string (should fail)
	var stringTarget string
	err = boolTrue.Unmarshal(&stringTarget)
	if err == nil {
		t.Error("Unmarshal() to string should return error")
	}
	
	// Test unmarshaling to non-pointer (should fail)
	var badTarget bool
	err = boolTrue.Unmarshal(badTarget)
	if err == nil {
		t.Error("Unmarshal() should return error for non-pointer target")
	}
	
	// Test unmarshaling to nil (should fail)
	err = boolTrue.Unmarshal(nil)
	if err == nil {
		t.Error("Unmarshal() should return error for nil target")
	}
}
