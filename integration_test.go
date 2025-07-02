package aaronjson

import (
	"testing"
)

// TestCompleteWorkflow tests the complete JSON parsing and marshaling workflow
func TestCompleteWorkflow(t *testing.T) {
	// Original JSON data
	jsonData := `{
		"name": "John Doe",
		"age": 30,
		"email": "john@example.com",
		"active": true,
		"address": {
			"street": "123 Main St",
			"city": "New York",
			"zipcode": "10001"
		},
		"hobbies": ["reading", "swimming", "coding"],
		"metadata": null
	}`

	// Parse JSON
	parsed, err := Parse(jsonData)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	// Verify it's an object
	if !parsed.IsObject() {
		t.Fatalf("Parsed result should be an object, got: %s", parsed.String())
	}

	// Get object and verify fields
	obj, err := parsed.AsObject()
	if err != nil {
		t.Fatalf("AsObject() error = %v", err)
	}

	// Test getting string field
	name, err := obj.Get("name")
	if err != nil {
		t.Errorf("Get('name') error = %v", err)
	}
	if name.String() != "John Doe" {
		t.Errorf("Name = %v, want John Doe", name.String())
	}

	// Test getting nested object
	address, err := obj.Get("address")
	if err != nil {
		t.Errorf("Get('address') error = %v", err)
	}
	if !address.IsObject() {
		t.Error("Address should be an object")
	}

	// Test getting array
	hobbies, err := obj.Get("hobbies")
	if err != nil {
		t.Errorf("Get('hobbies') error = %v", err)
	}
	if !hobbies.IsArray() {
		t.Error("Hobbies should be an array")
	}

	// Test getting null field
	metadata, err := obj.Get("metadata")
	if err != nil {
		t.Errorf("Get('metadata') error = %v", err)
	}
	if !metadata.IsNull() {
		t.Error("Metadata should be null")
	}

	// Test marshal back to JSON-like structure
	type Person struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Active bool   `json:"active"`
	}

	var person Person
	err = parsed.Unmarshal(&person)
	if err != nil {
		t.Errorf("Unmarshal() error = %v", err)
	}

	if person.Name != "John Doe" {
		t.Errorf("Person.Name = %v, want John Doe", person.Name)
	}
	if person.Age != 30 {
		t.Errorf("Person.Age = %v, want 30", person.Age)
	}
	if !person.Active {
		t.Errorf("Person.Active = %v, want true", person.Active)
	}
}

func TestRoundTripMarshalParse(t *testing.T) {
	// Test data structure
	type Address struct {
		Street string `json:"street"`
		City   string `json:"city"`
	}

	type Person struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Address Address  `json:"address"`
		Hobbies []string `json:"hobbies"`
	}

	original := Person{
		Name: "Alice",
		Age:  25,
		Address: Address{
			Street: "456 Oak Ave",
			City:   "Boston",
		},
		Hobbies: []string{"painting", "hiking"},
	}

	// Marshal to JsonValue
	marshaled, err := Marshal(original)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	// Convert to string and parse back
	jsonStr := marshaled.String()
	t.Logf("Marshaled JSON: %s", jsonStr)

	parsed, err := Parse(jsonStr)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	// Unmarshal back to struct
	var result Person
	err = parsed.Unmarshal(&result)
	if err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}

	// Verify round-trip consistency
	if result.Name != original.Name {
		t.Errorf("Name: got %v, want %v", result.Name, original.Name)
	}
	if result.Age != original.Age {
		t.Errorf("Age: got %v, want %v", result.Age, original.Age)
	}
	if result.Address.Street != original.Address.Street {
		t.Errorf("Address.Street: got %v, want %v", result.Address.Street, original.Address.Street)
	}
	if result.Address.City != original.Address.City {
		t.Errorf("Address.City: got %v, want %v", result.Address.City, original.Address.City)
	}
	if len(result.Hobbies) != len(original.Hobbies) {
		t.Errorf("Hobbies length: got %v, want %v", len(result.Hobbies), len(original.Hobbies))
	}
	for i, hobby := range result.Hobbies {
		if hobby != original.Hobbies[i] {
			t.Errorf("Hobbies[%d]: got %v, want %v", i, hobby, original.Hobbies[i])
		}
	}
}

func TestEdgeCases(t *testing.T) {
	// Test deeply nested structures
	deepJSON := `{"a":{"b":{"c":{"d":{"e":"deep"}}}}}`
	parsed, err := Parse(deepJSON)
	if err != nil {
		t.Fatalf("Parse deep JSON error = %v", err)
	}

	// Navigate deep structure
	val, err := parsed.Get("a")
	if err != nil {
		t.Errorf("Get('a') error = %v", err)
	}
	val, err = val.Get("b")
	if err != nil {
		t.Errorf("Get('b') error = %v", err)
	}
	val, err = val.Get("c")
	if err != nil {
		t.Errorf("Get('c') error = %v", err)
	}
	val, err = val.Get("d")
	if err != nil {
		t.Errorf("Get('d') error = %v", err)
	}
	val, err = val.Get("e")
	if err != nil {
		t.Errorf("Get('e') error = %v", err)
	}
	if val.String() != "deep" {
		t.Errorf("Deep value = %v, want deep", val.String())
	}

	// Test large array
	largeArrayJSON := `[1,2,3,4,5,6,7,8,9,10]`
	parsed, err = Parse(largeArrayJSON)
	if err != nil {
		t.Fatalf("Parse large array error = %v", err)
	}

	arr, err := parsed.AsArray()
	if err != nil {
		t.Fatalf("AsArray() error = %v", err)
	}

	length, _ := arr.Length()
	if length != 10 {
		t.Errorf("Array length = %v, want 10", length)
	}

	// Test mixed type array
	mixedJSON := `[1, "hello", true, null, {"key": "value"}, [1,2,3]]`
	parsed, err = Parse(mixedJSON)
	if err != nil {
		t.Fatalf("Parse mixed array error = %v", err)
	}

	arr, err = parsed.AsArray()
	if err != nil {
		t.Fatalf("AsArray() error = %v", err)
	}

	// Check each element type
	elem0, _ := arr.Index(0)
	if !elem0.IsInt() {
		t.Error("Element 0 should be int")
	}

	elem1, _ := arr.Index(1)
	if !elem1.IsString() {
		t.Error("Element 1 should be string")
	}

	elem2, _ := arr.Index(2)
	if !elem2.IsBool() {
		t.Error("Element 2 should be bool")
	}

	elem3, _ := arr.Index(3)
	if !elem3.IsNull() {
		t.Error("Element 3 should be null")
	}

	elem4, _ := arr.Index(4)
	if !elem4.IsObject() {
		t.Error("Element 4 should be object")
	}

	elem5, _ := arr.Index(5)
	if !elem5.IsArray() {
		t.Error("Element 5 should be array")
	}
}

func TestPrettyPrint(t *testing.T) {
	// Test pretty printing for complex structures
	jsonData := `{"name":"John","age":30,"hobbies":["reading","coding"]}`
	parsed, err := Parse(jsonData)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	pretty := parsed.PrettyString()
	if pretty == "" {
		t.Error("PrettyString() should not return empty string")
	}

	// Pretty string should be different from regular string (should have formatting)
	regular := parsed.String()
	if pretty == regular {
		t.Error("PrettyString() should be different from String() for complex objects")
	}
}

func TestErrorHandling(t *testing.T) {
	// Test various error conditions
	errorCases := []string{
		`{`,              // Unterminated object
		`[`,              // Unterminated array
		`"unterminated`,  // Unterminated string
		`{"key":}`,       // Missing value
		`{"key": value}`, // Unquoted value
		`{key: "value"}`, // Unquoted key
		`[1,2,]`,         // Trailing comma
		`{"a":1,}`,       // Trailing comma in object
		`123.`,           // Invalid number
		`--123`,          // Invalid number
		`trues`,          // Invalid boolean
		`nulls`,          // Invalid null
		``,               // Empty string
		`    `,           // Only whitespace
	}

	for i, jsonStr := range errorCases {
		t.Run(jsonStr, func(t *testing.T) {
			_, err := Parse(jsonStr)
			if err == nil {
				t.Errorf("Case %d: Parse(%q) should return error but didn't", i, jsonStr)
			}
		})
	}
}
