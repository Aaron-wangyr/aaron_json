package aaronjson

import (
	"testing"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		want    string
		wantErr bool
	}{
		{
			name:  "nil value",
			input: nil,
			want:  "null",
		},
		{
			name:  "string value",
			input: "hello",
			want:  "hello",
		},
		{
			name:  "int value",
			input: 42,
			want:  "42.000000",
		},
		{
			name:  "float value",
			input: 3.14,
			want:  "3.140000",
		},
		{
			name:  "bool true",
			input: true,
			want:  "true",
		},
		{
			name:  "bool false",
			input: false,
			want:  "false",
		},
		{
			name:  "slice of strings",
			input: []string{"hello", "world"},
			want:  "[\"hello\", \"world\"]",
		},
		{
			name:  "slice of ints",
			input: []int{1, 2, 3},
			want:  "[1.000000, 2.000000, 3.000000]",
		},
		{
			name:  "map with string keys",
			input: map[string]interface{}{"name": "John", "age": 30},
			want:  `{"age": 30.000000, "name": "John"}`,
		},
		{
			name:  "empty slice",
			input: []string{},
			want:  "[]",
		},
		{
			name:  "empty map",
			input: map[string]interface{}{},
			want:  "{}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.String() != tt.want {
				t.Errorf("Marshal() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestMarshalStruct(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	type PersonWithOmit struct {
		Name  string `json:"name"`
		Age   int    `json:"age,omitempty"`
		Email string `json:"email,omitempty"`
	}

	type PersonWithIgnore struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Password string `json:"-"`
	}

	tests := []struct {
		name  string
		input interface{}
		want  string
	}{
		{
			name:  "simple struct",
			input: Person{Name: "John", Age: 30},
			want:  `{"age": 30.000000, "name": "John"}`,
		},
		{
			name:  "struct with omitempty - non-zero values",
			input: PersonWithOmit{Name: "John", Age: 30, Email: "john@example.com"},
			want:  `{"age": 30.000000, "email": "john@example.com", "name": "John"}`,
		},
		{
			name:  "struct with omitempty - zero values",
			input: PersonWithOmit{Name: "John", Age: 0, Email: ""},
			want:  `{"name": "John"}`,
		},
		{
			name:  "struct with ignored field",
			input: PersonWithIgnore{Name: "John", Age: 30, Password: "secret"},
			want:  `{"age": 30.000000, "name": "John"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.input)
			if err != nil {
				t.Errorf("Marshal() error = %v", err)
				return
			}
			if got.String() != tt.want {
				t.Errorf("Marshal() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestMarshalPointer(t *testing.T) {
	str := "hello"
	num := 42

	tests := []struct {
		name  string
		input interface{}
		want  string
	}{
		{
			name:  "pointer to string",
			input: &str,
			want:  "hello",
		},
		{
			name:  "pointer to int",
			input: &num,
			want:  "42.000000",
		},
		{
			name:  "nil pointer",
			input: (*string)(nil),
			want:  "null",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.input)
			if err != nil {
				t.Errorf("Marshal() error = %v", err)
				return
			}
			if got.String() != tt.want {
				t.Errorf("Marshal() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestMarshalArray(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  string
	}{
		{
			name:  "array of strings",
			input: [3]string{"a", "b", "c"},
			want:  "[\"a\", \"b\", \"c\"]",
		},
		{
			name:  "array of ints",
			input: [2]int{1, 2},
			want:  "[1.000000, 2.000000]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.input)
			if err != nil {
				t.Errorf("Marshal() error = %v", err)
				return
			}
			if got.String() != tt.want {
				t.Errorf("Marshal() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestMarshalNestedStructures(t *testing.T) {
	type Address struct {
		Street string `json:"street"`
		City   string `json:"city"`
	}

	type Person struct {
		Name    string  `json:"name"`
		Address Address `json:"address"`
		Hobbies []string `json:"hobbies"`
	}

	person := Person{
		Name: "John",
		Address: Address{
			Street: "123 Main St",
			City:   "New York",
		},
		Hobbies: []string{"reading", "swimming"},
	}

	got, err := Marshal(person)
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
		return
	}

	// Check that it contains expected components
	result := got.String()
	if !contains(result, `"name": "John"`) {
		t.Error("Result should contain name field")
	}
	if !contains(result, `"address"`) {
		t.Error("Result should contain address field")
	}
	if !contains(result, `"hobbies"`) {
		t.Error("Result should contain hobbies field")
	}
}

func TestMarshalErrors(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{
			name:    "map with non-string keys",
			input:   map[int]string{1: "one", 2: "two"},
			wantErr: true,
		},
		{
			name:    "unsupported type (channel)",
			input:   make(chan int),
			wantErr: true,
		},
		{
			name:    "unsupported type (function)",
			input:   func() {},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Marshal(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMarshalInterface(t *testing.T) {
	var iface interface{} = "hello"
	got, err := Marshal(iface)
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
		return
	}
	if got.String() != "hello" {
		t.Errorf("Marshal() = %v, want hello", got.String())
	}

	// Test nil interface
	var nilIface interface{} = nil
	got, err = Marshal(nilIface)
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
		return
	}
	if got.String() != "null" {
		t.Errorf("Marshal() = %v, want null", got.String())
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && indexOf(s, substr) >= 0
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
