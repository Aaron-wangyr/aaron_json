package aaronjson

import (
	"testing"
)

func TestNewJsonInt(t *testing.T) {
	num := NewJsonInt(42)
	if num == nil {
		t.Error("NewJsonInt() returned nil")
	}
	if !num.IsInt() {
		t.Error("NewJsonInt() should return an int")
	}
}

func TestJsonIntAsInt(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  int
	}{
		{
			name:  "positive integer",
			value: 42,
			want:  42,
		},
		{
			name:  "negative integer",
			value: -10,
			want:  -10,
		},
		{
			name:  "zero",
			value: 0,
			want:  0,
		},
		{
			name:  "float converted to int",
			value: 3.14,
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := NewJsonInt(tt.value)
			got, err := num.AsInt()
			if err != nil {
				t.Errorf("AsInt() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("AsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonIntAsFloat(t *testing.T) {
	num := NewJsonInt(42)
	
	// JsonInt.AsFloat() should return error according to the implementation
	_, err := num.AsFloat()
	if err == nil {
		t.Error("AsFloat() should return error for JsonInt")
	}
}

func TestJsonIntString(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  string
	}{
		{
			name:  "positive integer",
			value: 42,
			want:  "42.000000",
		},
		{
			name:  "negative integer",
			value: -10,
			want:  "-10.000000",
		},
		{
			name:  "zero",
			value: 0,
			want:  "0.000000",
		},
		{
			name:  "float value",
			value: 3.14,
			want:  "3.140000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := NewJsonInt(tt.value)
			got := num.String()
			if got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonIntPrettyString(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  string
	}{
		{
			name:  "integer value",
			value: 42,
			want:  "42",
		},
		{
			name:  "negative integer",
			value: -10,
			want:  "-10",
		},
		{
			name:  "zero",
			value: 0,
			want:  "0",
		},
		{
			name:  "float value",
			value: 3.14,
			want:  "3.14",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := NewJsonInt(tt.value)
			got := num.PrettyString()
			if got != tt.want {
				t.Errorf("PrettyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonIntTypeChecks(t *testing.T) {
	num := NewJsonInt(42)
	
	// Test type checks
	if !num.IsInt() {
		t.Error("IsInt() should return true")
	}
	if num.IsString() {
		t.Error("IsString() should return false")
	}
	if num.IsFloat() {
		t.Error("IsFloat() should return false")
	}
	if num.IsBool() {
		t.Error("IsBool() should return false")
	}
	if num.IsArray() {
		t.Error("IsArray() should return false")
	}
	if num.IsObject() {
		t.Error("IsObject() should return false")
	}
	if num.IsNull() {
		t.Error("IsNull() should return false")
	}
}

func TestJsonIntUnmarshal(t *testing.T) {
	num := NewJsonInt(42)
	
	// Test unmarshal - currently returns nil per implementation
	var target int
	err := num.Unmarshal(&target)
	if err != nil {
		t.Errorf("Unmarshal() error = %v", err)
	}
	// Note: The current implementation of JsonInt.Unmarshal() returns nil without doing anything
	// This might be incomplete implementation
}
