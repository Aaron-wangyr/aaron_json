package aaronjson

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		jsonStr  string
		want     string
		wantErr  bool
	}{
		{
			name:    "parse simple string",
			jsonStr: `"hello"`,
			want:    "hello",
			wantErr: false,
		},
		{
			name:    "parse integer",
			jsonStr: `42`,
			want:    "42.000000",
			wantErr: false,
		},
		{
			name:    "parse float",
			jsonStr: `3.14`,
			want:    "3.140000",
			wantErr: false,
		},
		{
			name:    "parse boolean true",
			jsonStr: `true`,
			want:    "true",
			wantErr: false,
		},
		{
			name:    "parse boolean false",
			jsonStr: `false`,
			want:    "false",
			wantErr: false,
		},
		{
			name:    "parse null",
			jsonStr: `null`,
			want:    "null",
			wantErr: false,
		},
		{
			name:    "parse empty object",
			jsonStr: `{}`,
			want:    "{}",
			wantErr: false,
		},
		{
			name:    "parse empty array",
			jsonStr: `[]`,
			want:    "[]",
			wantErr: false,
		},
		{
			name:    "parse simple object",
			jsonStr: `{"name": "John", "age": 30}`,
			want:    `{"age": 30.000000, "name": "John"}`,
			wantErr: false,
		},
		{
			name:    "parse simple array",
			jsonStr: `[1, 2, 3]`,
			want:    "[1.000000, 2.000000, 3.000000]",
			wantErr: false,
		},
		{
			name:    "parse nested object",
			jsonStr: `{"person": {"name": "John", "age": 30}, "city": "New York"}`,
			want:    `{"city": "New York", "person": {"age": 30.000000, "name": "John"}}`,
			wantErr: false,
		},
		{
			name:    "parse array with mixed types",
			jsonStr: `[1, "hello", true, null]`,
			want:    `[1.000000, "hello", true, null]`,
			wantErr: false,
		},
		{
			name:    "parse with whitespace",
			jsonStr: ` { "key" : "value" } `,
			want:    `{"key": "value"}`,
			wantErr: false,
		},
		{
			name:    "empty data",
			jsonStr: "",
			wantErr: true,
		},
		{
			name:    "invalid JSON",
			jsonStr: `{invalid}`,
			wantErr: true,
		},
		{
			name:    "unterminated string",
			jsonStr: `"unterminated`,
			wantErr: true,
		},
		{
			name:    "invalid number",
			jsonStr: `123.`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.jsonStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.String() != tt.want {
				t.Errorf("Parse() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestParseByte(t *testing.T) {
	tests := []struct {
		name     string
		jsonData []byte
		want     string
		wantErr  bool
	}{
		{
			name:     "parse byte slice",
			jsonData: []byte(`{"key": "value"}`),
			want:     `{"key": "value"}`,
			wantErr:  false,
		},
		{
			name:     "empty byte slice",
			jsonData: []byte{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseByte(tt.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.String() != tt.want {
				t.Errorf("ParseByte() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestParseObject(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		wantErr bool
	}{
		{
			name:    "valid object",
			jsonStr: `{"name": "John", "age": 30}`,
			wantErr: false,
		},
		{
			name:    "empty object",
			jsonStr: `{}`,
			wantErr: false,
		},
		{
			name:    "nested object",
			jsonStr: `{"outer": {"inner": "value"}}`,
			wantErr: false,
		},
		{
			name:    "object with array",
			jsonStr: `{"items": [1, 2, 3]}`,
			wantErr: false,
		},
		{
			name:    "missing closing brace",
			jsonStr: `{"key": "value"`,
			wantErr: true,
		},
		{
			name:    "missing colon",
			jsonStr: `{"key" "value"}`,
			wantErr: true,
		},
		{
			name:    "missing comma",
			jsonStr: `{"key1": "value1" "key2": "value2"}`,
			wantErr: true,
		},
		{
			name:    "empty key",
			jsonStr: `{"": "value"}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.jsonStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParseArray(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		wantErr bool
	}{
		{
			name:    "valid array",
			jsonStr: `[1, 2, 3]`,
			wantErr: false,
		},
		{
			name:    "empty array",
			jsonStr: `[]`,
			wantErr: false,
		},
		{
			name:    "nested array",
			jsonStr: `[[1, 2], [3, 4]]`,
			wantErr: false,
		},
		{
			name:    "array with objects",
			jsonStr: `[{"key": "value"}, {"key2": "value2"}]`,
			wantErr: false,
		},
		{
			name:    "missing closing bracket",
			jsonStr: `[1, 2, 3`,
			wantErr: true,
		},
		{
			name:    "missing comma",
			jsonStr: `[1 2 3]`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.jsonStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParseString(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		want    string
		wantErr bool
	}{
		{
			name:    "simple string",
			jsonStr: `"hello"`,
			want:    "hello",
			wantErr: false,
		},
		{
			name:    "empty string",
			jsonStr: `""`,
			want:    "",
			wantErr: false,
		},
		{
			name:    "string with escape",
			jsonStr: `"hello\nworld"`,
			want:    `hello\nworld`,
			wantErr: false,
		},
		{
			name:    "unterminated string",
			jsonStr: `"hello`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.jsonStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.String() != tt.want {
				t.Errorf("Parse() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestParseNumber(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		wantErr bool
	}{
		{
			name:    "positive integer",
			jsonStr: `123`,
			wantErr: false,
		},
		{
			name:    "negative integer",
			jsonStr: `-123`,
			wantErr: false,
		},
		{
			name:    "zero",
			jsonStr: `0`,
			wantErr: false,
		},
		{
			name:    "float",
			jsonStr: `123.45`,
			wantErr: false,
		},
		{
			name:    "negative float",
			jsonStr: `-123.45`,
			wantErr: false,
		},
		{
			name:    "scientific notation",
			jsonStr: `1.23e10`,
			wantErr: false,
		},
		{
			name:    "scientific notation negative",
			jsonStr: `1.23e-10`,
			wantErr: false,
		},
		{
			name:    "leading plus sign (invalid)",
			jsonStr: `+123`,
			wantErr: true,
		},
		{
			name:    "leading zero with digit (invalid)",
			jsonStr: `01`,
			wantErr: true,
		},
		{
			name:    "decimal without fractional part",
			jsonStr: `123.`,
			wantErr: true,
		},
		{
			name:    "exponent without digits",
			jsonStr: `123e`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.jsonStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParseBool(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		want    string
		wantErr bool
	}{
		{
			name:    "true",
			jsonStr: `true`,
			want:    "true",
			wantErr: false,
		},
		{
			name:    "false",
			jsonStr: `false`,
			want:    "false",
			wantErr: false,
		},
		{
			name:    "invalid boolean",
			jsonStr: `tru`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.jsonStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.String() != tt.want {
				t.Errorf("Parse() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestParseNull(t *testing.T) {
	tests := []struct {
		name    string
		jsonStr string
		want    string
		wantErr bool
	}{
		{
			name:    "null",
			jsonStr: `null`,
			want:    "null",
			wantErr: false,
		},
		{
			name:    "invalid null",
			jsonStr: `nul`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.jsonStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.String() != tt.want {
				t.Errorf("Parse() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}
