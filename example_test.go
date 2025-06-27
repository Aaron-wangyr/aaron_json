package aaronjson

import (
	"fmt"
	"log"
)

// Example_basicParsing demonstrates basic JSON parsing functionality
func Example_basicParsing() {
	// Parse a simple JSON string
	jsonStr := `{"name":"Aaron","age":25,"active":true}`
	data, err := Parse(jsonStr)
	if err != nil {
		log.Fatal(err)
	}

	name, _ := data.Get("name")
	age, _ := data.Get("age")
	active, _ := data.Get("active")

	fmt.Printf("Name: %s\n", name.String())
	fmt.Printf("Age: %s\n", age.String())
	fmt.Printf("Active: %s\n", active.String())

	// Output:
	// Name: Aaron
	// Age: 25.000000
	// Active: true
}

// Example_parseArray demonstrates JSON array parsing
func Example_parseArray() {
	jsonStr := `[1, 2.5, "hello", true, null]`
	data, err := Parse(jsonStr)
	if err != nil {
		log.Fatal(err)
	}

	length, _ := data.Length()
	fmt.Printf("Array length: %d\n", length)

	for i := 0; i < length; i++ {
		item, _ := data.Index(i)
		fmt.Printf("Item[%d]: %s (Type: %d)\n", i, item.String(), item.Type())
	}

	// Output:
	// Array length: 5
	// Item[0]: 1.000000 (Type: 2)
	// Item[1]: 2.500000 (Type: 2)
	// Item[2]: hello (Type: 3)
	// Item[3]: true (Type: 1)
	// Item[4]: null (Type: 0)
}

// Example_nestedJson demonstrates parsing nested JSON structures
func Example_nestedJson() {
	jsonStr := `{
		"user": {
			"name": "John",
			"age": 30,
			"hobbies": ["reading", "swimming", "coding"]
		},
		"settings": {
			"theme": "dark",
			"notifications": true
		}
	}`

	data, err := Parse(jsonStr)
	if err != nil {
		log.Fatal(err)
	}

	user, _ := data.Get("user")
	name, _ := user.Get("name")
	hobbies, _ := user.Get("hobbies")

	fmt.Printf("User name: %s\n", name.String())

	hobbiesLength, _ := hobbies.Length()
	fmt.Printf("Hobbies (%d):\n", hobbiesLength)
	for i := 0; i < hobbiesLength; i++ {
		hobby, _ := hobbies.Index(i)
		fmt.Printf("  - %s\n", hobby.String())
	}

	// Output:
	// User name: John
	// Hobbies (3):
	//   - reading
	//   - swimming
	//   - coding
}

// Example_typeChecking demonstrates type checking methods
func Example_typeChecking() {
	jsonStr := `{
		"text": "hello",
		"number": 42,
		"flag": true,
		"empty": null,
		"list": [1, 2, 3],
		"object": {"key": "value"}
	}`

	data, err := Parse(jsonStr)
	if err != nil {
		log.Fatal(err)
	}

	fields := []string{"text", "number", "flag", "empty", "list", "object"}
	for _, field := range fields {
		value, _ := data.Get(field)
		fmt.Printf("%s: ", field)
		if value.IsString() {
			fmt.Print("String")
		}
		if value.IsNumber() {
			fmt.Print("Number")
		}
		if value.IsBool() {
			fmt.Print("Bool")
		}
		if value.IsNull() {
			fmt.Print("Null")
		}
		if value.IsArray() {
			fmt.Print("Array")
		}
		if value.IsObject() {
			fmt.Print("Object")
		}
		fmt.Println()
	}

	// Output:
	// text: String
	// number: Number
	// flag: Bool
	// empty: Null
	// list: Array
	// object: Object
}

// Example_modifyJson demonstrates JSON modification operations
func Example_modifyJson() {
	// Create a new JSON object
	obj := NewJsonObject()
	obj.Set("name", NewJsonString("Alice"))
	obj.Set("age", NewJsonNumber(28))
	obj.Set("active", NewJsonBool(true))

	// Create and add an array
	arr := NewJsonArray()
	arr.Append(NewJsonString("golang"))
	arr.Append(NewJsonString("javascript"))
	arr.Append(NewJsonString("python"))
	obj.Set("skills", arr)

	// Access and modify values
	name, _ := obj.Get("name")
	fmt.Printf("Original name: %s\n", name.String())

	obj.Set("name", NewJsonString("Bob"))
	newName, _ := obj.Get("name")
	fmt.Printf("Updated name: %s\n", newName.String())

	// Modify array
	skills, _ := obj.Get("skills")
	skills.Append(NewJsonString("rust"))
	length, _ := skills.Length()
	fmt.Printf("Skills count: %d\n", length)

	// Output:
	// Original name: Alice
	// Updated name: Bob
	// Skills count: 4
}

// Example_arrayOperations demonstrates array manipulation
func Example_arrayOperations() {
	arr := NewJsonArray()

	// Add elements
	arr.Append(NewJsonString("first"))
	arr.Append(NewJsonString("second"))
	arr.Append(NewJsonString("third"))

	length, _ := arr.Length()
	fmt.Printf("Initial length: %d\n", length)

	// Access elements
	first, _ := arr.Index(0)
	fmt.Printf("First element: %s\n", first.String())

	// Modify element
	arr.SetByIndex(1, NewJsonString("modified"))
	second, _ := arr.Index(1)
	fmt.Printf("Modified second element: %s\n", second.String())

	// Remove element
	removed, _ := arr.RemoveByIndex(0)
	fmt.Printf("Removed element: %s\n", removed.String())

	newLength, _ := arr.Length()
	fmt.Printf("Final length: %d\n", newLength)

	// Output:
	// Initial length: 3
	// First element: first
	// Modified second element: modified
	// Removed element: first
	// Final length: 2
}

// Example_objectOperations demonstrates object manipulation
func Example_objectOperations() {
	obj := NewJsonObject()

	// Set values
	obj.Set("key1", NewJsonString("value1"))
	obj.Set("key2", NewJsonNumber(123))
	obj.Set("key3", NewJsonBool(false))

	// Get all keys
	keys, _ := obj.Keys()
	fmt.Printf("Keys: %v\n", keys)

	// Check if key exists and get value
	value, err := obj.Get("key1")
	if err == nil {
		fmt.Printf("key1 exists: %s\n", value.String())
	}

	// Remove a key
	removed, _ := obj.Remove("key2")
	if removed != nil {
		fmt.Printf("Removed key2: %s\n", removed.String())
	}

	// Final keys
	finalKeys, _ := obj.Keys()
	fmt.Printf("Final keys: %v\n", finalKeys)

	// Output:
	// Keys: [key1 key2 key3]
	// key1 exists: value1
	// Removed key2: 123.000000
	// Final keys: [key1 key3]
}

// Example_unmarshalBasicTypes demonstrates unmarshaling to basic types
func Example_unmarshalBasicTypes() {
	// String
	strJson, _ := Parse(`"hello world"`)
	var str string
	strJson.Unmarshal(&str)
	fmt.Printf("String: %s\n", str)

	// Number to int
	numJson, _ := Parse(`42`)
	var num int
	numJson.Unmarshal(&num)
	fmt.Printf("Number: %d\n", num)

	// Number to float
	floatJson, _ := Parse(`3.14`)
	var f float64
	floatJson.Unmarshal(&f)
	fmt.Printf("Float: %.2f\n", f)

	// Boolean
	boolJson, _ := Parse(`true`)
	var b bool
	boolJson.Unmarshal(&b)
	fmt.Printf("Boolean: %t\n", b)

	// Output:
	// String: hello world
	// Number: 42
	// Float: 3.14
	// Boolean: true
}

// Example_unmarshalStruct demonstrates unmarshaling to struct
func Example_unmarshalStruct() {
	type Person struct {
		Name   string  `json:"name"`
		Age    int     `json:"age"`
		Height float64 `json:"height"`
		Active bool    `json:"active"`
		Email  string  `json:"email"`
	}

	jsonStr := `{
		"name": "Alice",
		"age": 30,
		"height": 165.5,
		"active": true,
		"email": "alice@example.com"
	}`

	data, _ := Parse(jsonStr)
	var person Person
	err := data.Unmarshal(&person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Height: %.1f\n", person.Height)
	fmt.Printf("Active: %t\n", person.Active)
	fmt.Printf("Email: %s\n", person.Email)

	// Output:
	// Name: Alice
	// Age: 30
	// Height: 165.5
	// Active: true
	// Email: alice@example.com
}

// Example_unmarshalSlice demonstrates unmarshaling to slice
func Example_unmarshalSlice() {
	jsonStr := `[1, 2, 3, 4, 5]`
	data, _ := Parse(jsonStr)

	var numbers []int
	err := data.Unmarshal(&numbers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Numbers: %v\n", numbers)

	// String slice
	strJsonStr := `["apple", "banana", "cherry"]`
	strData, _ := Parse(strJsonStr)

	var fruits []string
	strData.Unmarshal(&fruits)
	fmt.Printf("Fruits: %v\n", fruits)

	// Output:
	// Numbers: [1 2 3 4 5]
	// Fruits: [apple banana cherry]
}

// Example_unmarshalMap demonstrates unmarshaling to map
func Example_unmarshalMap() {
	jsonStr := `{
		"apple": 1.20,
		"banana": 0.80,
		"cherry": 3.50
	}`

	data, _ := Parse(jsonStr)

	var prices map[string]float64
	err := data.Unmarshal(&prices)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Prices: %v\n", prices)

	// Interface map
	var genericMap map[string]interface{}
	data.Unmarshal(&genericMap)
	fmt.Printf("Generic map: %v\n", genericMap)

	// Output:
	// Prices: map[apple:1.2 banana:0.8 cherry:3.5]
	// Generic map: map[apple:1.2 banana:0.8 cherry:3.5]
}

// Example_marshalBasicTypes demonstrates marshaling from basic types
func Example_marshalBasicTypes() {
	// Marshal string
	strData, _ := Marshal("hello")
	fmt.Printf("String: %s\n", strData.String())

	// Marshal number
	numData, _ := Marshal(42)
	fmt.Printf("Number: %s\n", numData.String())

	// Marshal float
	floatData, _ := Marshal(3.14)
	fmt.Printf("Float: %s\n", floatData.String())

	// Marshal boolean
	boolData, _ := Marshal(true)
	fmt.Printf("Boolean: %s\n", boolData.String())

	// Marshal nil
	nullData, _ := Marshal(nil)
	fmt.Printf("Null: %s\n", nullData.String())

	// Output:
	// String: hello
	// Number: 42.000000
	// Float: 3.140000
	// Boolean: true
	// Null: null
}

// Example_marshalStruct demonstrates marshaling from struct
func Example_marshalStruct() {
	type User struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Email    string `json:"email,omitempty"`
		Active   bool   `json:"active"`
		Password string `json:"-"` // This field will be ignored
	}

	user := User{
		Name:     "John",
		Age:      25,
		Email:    "john@example.com",
		Active:   true,
		Password: "secret",
	}

	data, _ := Marshal(user)
	obj, _ := data.AsObject()

	keys, _ := obj.Keys()
	fmt.Printf("Keys: %v\n", keys)

	name, _ := obj.Get("name")
	age, _ := obj.Get("age")
	email, _ := obj.Get("email")
	active, _ := obj.Get("active")

	fmt.Printf("Name: %s\n", name.String())
	fmt.Printf("Age: %s\n", age.String())
	fmt.Printf("Email: %s\n", email.String())
	fmt.Printf("Active: %s\n", active.String())

	// Output:
	// Keys: [active age email name]
	// Name: John
	// Age: 25.000000
	// Email: john@example.com
	// Active: true
}

// Example_marshalSlice demonstrates marshaling from slice
func Example_marshalSlice() {
	numbers := []int{1, 2, 3, 4, 5}
	data, _ := Marshal(numbers)

	length, _ := data.Length()
	fmt.Printf("Array length: %d\n", length)

	for i := 0; i < length; i++ {
		item, _ := data.Index(i)
		fmt.Printf("Item[%d]: %s\n", i, item.String())
	}

	// Output:
	// Array length: 5
	// Item[0]: 1.000000
	// Item[1]: 2.000000
	// Item[2]: 3.000000
	// Item[3]: 4.000000
	// Item[4]: 5.000000
}

// Example_marshalMap demonstrates marshaling from map
func Example_marshalMap() {
	data := map[string]interface{}{
		"name":   "Alice",
		"age":    30,
		"active": true,
		"scores": []int{85, 90, 95},
	}

	jsonData, _ := Marshal(data)
	obj, _ := jsonData.AsObject()

	keys, _ := obj.Keys()
	fmt.Printf("Keys: %v\n", keys)

	name, _ := obj.Get("name")
	age, _ := obj.Get("age")
	scores, _ := obj.Get("scores")

	fmt.Printf("Name: %s\n", name.String())
	fmt.Printf("Age: %s\n", age.String())

	scoresLength, _ := scores.Length()
	fmt.Printf("Scores (%d): ", scoresLength)
	for i := 0; i < scoresLength; i++ {
		score, _ := scores.Index(i)
		fmt.Printf("%s ", score.String())
	}
	fmt.Println()

	// Output:
	// Keys: [active age name scores]
	// Name: Alice
	// Age: 30.000000
	// Scores (3): 85.000000 90.000000 95.000000
}

// Example_complexNestedStructure demonstrates complex nested data structures
func Example_complexNestedStructure() {
	type Address struct {
		Street  string `json:"street"`
		City    string `json:"city"`
		Country string `json:"country"`
	}

	type Person struct {
		Name     string            `json:"name"`
		Age      int               `json:"age"`
		Address  Address           `json:"address"`
		Hobbies  []string          `json:"hobbies"`
		Metadata map[string]string `json:"metadata"`
	}

	person := Person{
		Name: "Bob",
		Age:  35,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
		Hobbies: []string{"reading", "gaming", "cooking"},
		Metadata: map[string]string{
			"department": "engineering",
			"level":      "senior",
		},
	}

	// Marshal to JsonData
	data, _ := Marshal(person)

	// Access nested data
	address, _ := data.Get("address")
	city, _ := address.Get("city")
	fmt.Printf("City: %s\n", city.String())

	hobbies, _ := data.Get("hobbies")
	firstHobby, _ := hobbies.Index(0)
	fmt.Printf("First hobby: %s\n", firstHobby.String())

	metadata, _ := data.Get("metadata")
	dept, _ := metadata.Get("department")
	fmt.Printf("Department: %s\n", dept.String())

	// Output:
	// City: New York
	// First hobby: reading
	// Department: engineering
}

// Example_errorHandling demonstrates error handling
func Example_errorHandling() {
	// Invalid JSON
	_, err := Parse(`{"name": "John", "age":}`)
	if err != nil {
		fmt.Printf("Parse error handled: %v\n", err != nil)
	}

	// Valid JSON but invalid access
	data, _ := Parse(`{"name": "John"}`)

	// Try to access non-existent key
	_, err = data.Get("age")
	if err != nil {
		fmt.Printf("Key not found error handled: %v\n", err != nil)
	}

	// Try to use array methods on object
	_, err = data.Index(0)
	if err != nil {
		fmt.Printf("Wrong type error handled: %v\n", err != nil)
	}

	// Try to unmarshal incompatible types
	var num int
	stringJson, _ := Parse(`"not a number"`)
	err = stringJson.Unmarshal(&num)
	if err != nil {
		fmt.Printf("Unmarshal type error handled: %v\n", err != nil)
	}

	// Output:
	// Parse error handled: true
	// Key not found error handled: true
	// Wrong type error handled: true
	// Unmarshal type error handled: true
}

// Example_edgeCases demonstrates edge cases and special scenarios
func Example_edgeCases() {
	// Empty object and array
	emptyObj, _ := Parse(`{}`)
	emptyArr, _ := Parse(`[]`)

	objLen, _ := emptyObj.Length()
	arrLen, _ := emptyArr.Length()
	fmt.Printf("Empty object length: %d\n", objLen)
	fmt.Printf("Empty array length: %d\n", arrLen)

	// Null values
	nullData, _ := Parse(`null`)
	fmt.Printf("Is null: %t\n", nullData.IsNull())

	// Large numbers
	largeNum, _ := Parse(`123456789.123456789`)
	var f float64
	largeNum.Unmarshal(&f)
	fmt.Printf("Large number: %.6f\n", f)

	// Empty string
	emptyStr, _ := Parse(`""`)
	var s string
	emptyStr.Unmarshal(&s)
	fmt.Printf("Empty string length: %d\n", len(s))

	// Boolean edge cases
	trueVal, _ := Parse(`true`)
	falseVal, _ := Parse(`false`)
	fmt.Printf("True: %s, False: %s\n", trueVal.String(), falseVal.String())

	// Output:
	// Empty object length: 0
	// Empty array length: 0
	// Is null: true
	// Large number: 123456789.123457
	// Empty string length: 0
	// True: true, False: false
}

// Example_performanceScenario demonstrates usage with larger data sets
func Example_performanceScenario() {
	// Create a larger data structure
	obj := NewJsonObject()

	// Add multiple fields
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("field_%d", i)
		obj.Set(key, NewJsonString(fmt.Sprintf("value_%d", i)))
	}

	// Add an array with multiple elements
	arr := NewJsonArray()
	for i := 0; i < 5; i++ {
		arr.Append(NewJsonNumber(float64(i * 10)))
	}
	obj.Set("numbers", arr)

	// Access and count elements
	keys, _ := obj.Keys()
	fmt.Printf("Object has %d keys\n", len(keys))

	numbers, _ := obj.Get("numbers")
	numLength, _ := numbers.Length()
	fmt.Printf("Array has %d elements\n", numLength)

	// Calculate sum of array elements
	sum := 0.0
	for i := 0; i < numLength; i++ {
		item, _ := numbers.Index(i)
		var val float64
		item.Unmarshal(&val)
		sum += val
	}
	fmt.Printf("Sum of array elements: %.0f\n", sum)

	// Output:
	// Object has 6 keys
	// Array has 5 elements
	// Sum of array elements: 100
}

// Example_roundTrip demonstrates marshal/unmarshal round trip
func Example_roundTrip() {
	type Data struct {
		Name   string `json:"name"`
		Values []int  `json:"values"`
		Active bool   `json:"active"`
		Meta   struct {
			Version string `json:"version"`
		} `json:"meta"`
	}

	original := Data{
		Name:   "Test",
		Values: []int{1, 2, 3},
		Active: true,
	}
	original.Meta.Version = "1.0.0"

	// Marshal to JsonData
	jsonData, _ := Marshal(original)

	// Unmarshal back to struct
	var result Data
	jsonData.Unmarshal(&result)

	fmt.Printf("Original name: %s\n", original.Name)
	fmt.Printf("Result name: %s\n", result.Name)
	fmt.Printf("Values match: %v\n", fmt.Sprintf("%v", original.Values) == fmt.Sprintf("%v", result.Values))
	fmt.Printf("Active match: %v\n", original.Active == result.Active)
	fmt.Printf("Version match: %v\n", original.Meta.Version == result.Meta.Version)

	// Output:
	// Original name: Test
	// Result name: Test
	// Values match: true
	// Active match: true
	// Version match: true
}

// Example_prettyString demonstrates pretty printing functionality
func Example_prettyString() {
	// Test basic types
	fmt.Println("=== Basic Types ===")

	str := NewJsonString("hello\nworld")
	fmt.Printf("String: %s\n", str.PrettyString())

	num := NewJsonNumber(42.5)
	fmt.Printf("Number: %s\n", num.PrettyString())

	integer := NewJsonNumber(42)
	fmt.Printf("Integer: %s\n", integer.PrettyString())

	boolTrue := NewJsonBool(true)
	fmt.Printf("Boolean: %s\n", boolTrue.PrettyString())

	null := NewJsonNull()
	fmt.Printf("Null: %s\n", null.PrettyString())

	// Test complex structure
	fmt.Println("\n=== Complex Structure ===")

	obj := NewJsonObject()
	obj.Set("name", NewJsonString("John"))
	obj.Set("age", NewJsonNumber(30))
	obj.Set("active", NewJsonBool(true))

	arr := NewJsonArray()
	arr.Append(NewJsonString("item1"))
	arr.Append(NewJsonNumber(123))
	arr.Append(NewJsonBool(false))

	obj.Set("items", arr)

	// Nested object
	nested := NewJsonObject()
	nested.Set("city", NewJsonString("New York"))
	nested.Set("country", NewJsonString("USA"))
	obj.Set("address", nested)

	fmt.Println(obj.PrettyString())

	// Output:
	// === Basic Types ===
	// String: "hello\nworld"
	// Number: 42.5
	// Integer: 42
	// Boolean: true
	// Null: null
	//
	// === Complex Structure ===
	// {
	//   "active": true,
	//   "address": {
	//     "city": "New York",
	//     "country": "USA"
	//   },
	//   "age": 30,
	//   "items": [
	//     "item1",
	//     123,
	//     false
	//   ],
	//   "name": "John"
	// }
}

// Example_prettyStringEscaping demonstrates string escaping in pretty printing
func Example_prettyStringEscaping() {
	// Test string with special characters
	specialStr := "Hello\n\t\"World\"\r\nWith\\Backslash"
	str := NewJsonString(specialStr)
	fmt.Printf("Escaped string: %s\n", str.PrettyString())

	// Test in object context
	obj := NewJsonObject()
	obj.Set("message", str)
	obj.Set("control_chars", NewJsonString("\b\f\n\r\t"))

	fmt.Println("Object with escaped strings:")
	fmt.Println(obj.PrettyString())

	// Output:
	// Escaped string: "Hello\n\t\"World\"\r\nWith\\Backslash"
	// Object with escaped strings:
	// {
	//   "control_chars": "\b\f\n\r\t",
	//   "message": "Hello\n\t\"World\"\r\nWith\\Backslash"
	// }
}
