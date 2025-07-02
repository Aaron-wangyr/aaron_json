package aaronjson

import (
	"testing"
)

func TestNewJsonFloat(t *testing.T) {
	num := NewJsonFloat(3.14)
	if num == nil {
		t.Error("NewJsonFloat() returned nil")
	}
	if !num.IsFloat() {
		t.Error("NewJsonFloat() should return a float")
	}
}

func TestJsonFloatAsFloat(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "positive float",
			value: 3.14,
			want:  3.14,
		},
		{
			name:  "negative float",
			value: -2.5,
			want:  -2.5,
		},
		{
			name:  "zero",
			value: 0.0,
			want:  0.0,
		},
		{
			name:  "scientific notation",
			value: 1.23e10,
			want:  1.23e10,
		},
		{
			name:  "small decimal",
			value: 0.001,
			want:  0.001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := NewJsonFloat(tt.value)
			got, err := num.AsFloat()
			if err != nil {
				t.Errorf("AsFloat() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("AsFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonFloatAsInt(t *testing.T) {
	num := NewJsonFloat(3.14)
	
	// JsonFloat.AsInt() should return error according to the implementation
	_, err := num.AsInt()
	if err == nil {
		t.Error("AsInt() should return error for JsonFloat")
	}
}

func TestJsonFloatString(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  string
	}{
		{
			name:  "positive float",
			value: 3.14,
			want:  "3.140000",
		},
		{
			name:  "negative float",
			value: -2.5,
			want:  "-2.500000",
		},
		{
			name:  "zero",
			value: 0.0,
			want:  "0.000000",
		},
		{
			name:  "integer-like float",
			value: 42.0,
			want:  "42.000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := NewJsonFloat(tt.value)
			got := num.String()
			if got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonFloatPrettyString(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  string
	}{
		{
			name:  "float value",
			value: 3.14,
			want:  "3.14",
		},
		{
			name:  "negative float",
			value: -2.5,
			want:  "-2.5",
		},
		{
			name:  "zero",
			value: 0.0,
			want:  "0",
		},
		{
			name:  "integer-like float",
			value: 42.0,
			want:  "42",
		},
		{
			name:  "scientific notation small",
			value: 0.000001,
			want:  "1e-06",
		},
		{
			name:  "scientific notation large",
			value: 1000000000000.0,
			want:  "1e+12",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := NewJsonFloat(tt.value)
			got := num.PrettyString()
			if got != tt.want {
				t.Errorf("PrettyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonFloatTypeChecks(t *testing.T) {
	num := NewJsonFloat(3.14)
	
	// Test type checks
	if !num.IsFloat() {
		t.Error("IsFloat() should return true")
	}
	if num.IsString() {
		t.Error("IsString() should return false")
	}
	if num.IsInt() {
		t.Error("IsInt() should return false")
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

func TestJsonFloatUnmarshal(t *testing.T) {
	num := NewJsonFloat(3.14)
	
	// Test unmarshal - currently returns nil per implementation
	var target float64
	err := num.Unmarshal(&target)
	if err != nil {
		t.Errorf("Unmarshal() error = %v", err)
	}
	// Note: The current implementation of JsonFloat.Unmarshal() returns nil without doing anything
	// This might be incomplete implementation
}
