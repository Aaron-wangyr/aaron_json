package aaronjson

import (
	"testing"
)

func TestNewJsonString(t *testing.T) {
	str := NewJsonString("hello")
	if str == nil {
		t.Error("NewJsonString() returned nil")
	}
	if !str.IsString() {
		t.Error("NewJsonString() should return a string")
	}
	if str.String() != "hello" {
		t.Errorf("NewJsonString().String() = %v, want hello", str.String())
	}
}

func TestJsonStringAsString(t *testing.T) {
	str := NewJsonString("hello world")
	
	val, err := str.AsString()
	if err != nil {
		t.Errorf("AsString() error = %v", err)
	}
	if val != "hello world" {
		t.Errorf("AsString() = %v, want hello world", val)
	}
}

func TestJsonStringAsInt(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    int
		wantErr bool
	}{
		{
			name:    "valid integer string",
			value:   "42",
			want:    42,
			wantErr: false,
		},
		{
			name:    "negative integer string",
			value:   "-10",
			want:    -10,
			wantErr: false,
		},
		{
			name:    "zero string",
			value:   "0",
			want:    0,
			wantErr: false,
		},
		{
			name:    "invalid integer string",
			value:   "hello",
			wantErr: true,
		},
		{
			name:    "float string",
			value:   "3.14",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := NewJsonString(tt.value)
			got, err := str.AsInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("AsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonStringAsFloat(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    float64
		wantErr bool
	}{
		{
			name:    "valid float string",
			value:   "3.14",
			want:    3.14,
			wantErr: false,
		},
		{
			name:    "integer string",
			value:   "42",
			want:    42.0,
			wantErr: false,
		},
		{
			name:    "negative float string",
			value:   "-2.5",
			want:    -2.5,
			wantErr: false,
		},
		{
			name:    "scientific notation",
			value:   "1.23e10",
			want:    1.23e10,
			wantErr: false,
		},
		{
			name:    "invalid float string",
			value:   "hello",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := NewJsonString(tt.value)
			got, err := str.AsFloat()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("AsFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonStringAsBool(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    bool
		wantErr bool
	}{
		{
			name:    "true string",
			value:   "true",
			want:    true,
			wantErr: false,
		},
		{
			name:    "false string",
			value:   "false",
			want:    false,
			wantErr: false,
		},
		{
			name:    "1 string",
			value:   "1",
			want:    true,
			wantErr: false,
		},
		{
			name:    "0 string",
			value:   "0",
			want:    false,
			wantErr: false,
		},
		{
			name:    "invalid bool string",
			value:   "hello",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := NewJsonString(tt.value)
			got, err := str.AsBool()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("AsBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonStringLength(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  int
	}{
		{
			name:  "empty string",
			value: "",
			want:  0,
		},
		{
			name:  "single character",
			value: "a",
			want:  1,
		},
		{
			name:  "multi character",
			value: "hello world",
			want:  11,
		},
		{
			name:  "unicode string",
			value: "你好",
			want:  6, // UTF-8 encoding
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := NewJsonString(tt.value)
			got, err := str.Length()
			if err != nil {
				t.Errorf("Length() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonStringUnmarshal(t *testing.T) {
	str := NewJsonString("hello world")
	
	// Test unmarshaling to string
	var target string
	err := str.Unmarshal(&target)
	if err != nil {
		t.Errorf("Unmarshal() error = %v", err)
	}
	if target != "hello world" {
		t.Errorf("Unmarshal() result = %v, want hello world", target)
	}
	
	// Test unmarshaling to interface{}
	var interfaceTarget interface{}
	err = str.Unmarshal(&interfaceTarget)
	if err != nil {
		t.Errorf("Unmarshal() to interface{} error = %v", err)
	}
	if interfaceTarget != "hello world" {
		t.Errorf("Unmarshal() to interface{} = %v, want hello world", interfaceTarget)
	}
	
	// Test unmarshaling to int (should fail)
	var intTarget int
	err = str.Unmarshal(&intTarget)
	if err == nil {
		t.Error("Unmarshal() to int should return error")
	}
	
	// Test unmarshaling to non-pointer (should fail)
	var badTarget string
	err = str.Unmarshal(badTarget)
	if err == nil {
		t.Error("Unmarshal() should return error for non-pointer target")
	}
	
	// Test unmarshaling to nil (should fail)
	err = str.Unmarshal(nil)
	if err == nil {
		t.Error("Unmarshal() should return error for nil target")
	}
}

func TestJsonStringPrettyString(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "simple string",
			value: "hello",
			want:  `"hello"`,
		},
		{
			name:  "empty string",
			value: "",
			want:  `""`,
		},
		{
			name:  "string with quotes",
			value: `he said "hello"`,
			want:  `"he said \"hello\""`,
		},
		{
			name:  "string with newline",
			value: "hello\nworld",
			want:  `"hello\nworld"`,
		},
		{
			name:  "string with tab",
			value: "hello\tworld",
			want:  `"hello\tworld"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := NewJsonString(tt.value)
			got := str.PrettyString()
			if got != tt.want {
				t.Errorf("PrettyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonStringTypeChecks(t *testing.T) {
	str := NewJsonString("hello")
	
	// Test type checks
	if !str.IsString() {
		t.Error("IsString() should return true")
	}
	if str.IsInt() {
		t.Error("IsInt() should return false")
	}
	if str.IsFloat() {
		t.Error("IsFloat() should return false")
	}
	if str.IsBool() {
		t.Error("IsBool() should return false")
	}
	if str.IsArray() {
		t.Error("IsArray() should return false")
	}
	if str.IsObject() {
		t.Error("IsObject() should return false")
	}
	if str.IsNull() {
		t.Error("IsNull() should return false")
	}
}
