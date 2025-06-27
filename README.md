# Aaron JSON - Modern JSON Processing Library with Unmarshal Support

Aaron JSON is a modern Go JSON processing library that provides chain-style access, type-safe operations, and comprehensive Unmarshal functionality.

## Features

### ✅ **Unmarshal Methods Added**
- `Unmarshal(v interface{}) error` - Unmarshals JsonData into Go values
- `UnmarshalTo(v interface{}) error` - Alias for Unmarshal (API consistency)

### ✅ **Supported Target Types**
- **Basic Types**: `string`, `int`, `int32`, `int64`, `float32`, `float64`, `bool`
- **Complex Types**: `struct`, `map[string]interface{}`, `[]interface{}`
- **Interface{}**: Automatically converts to appropriate Go types
- **Nested Structures**: Full support for embedded structs and slices

### ✅ **Struct Tag Support**
- Standard json tags: `json:"field_name"`
- Omitempty support: `json:"field_name,omitempty"`
- Automatic field mapping from JSON keys to struct fields

### ✅ **Chain Compatibility**
- Seamless integration with chain operations
- Error propagation through the entire chain
- Type-safe access before unmarshaling

## Usage Examples

### Basic Struct Unmarshaling
```go
type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"`
}

data, err := Parse(`{"name": "John", "age": 30, "email": "john@example.com"}`)
if err != nil {
    // Handle parse error
    return
}
var person Person
err = data.Unmarshal(&person)
```

### Chain Access + Unmarshal
```go
// Get nested field and unmarshal
var name string
err := data.Get("user").Get("name").Unmarshal(&name)

// Get array element and unmarshal  
var firstUser User
err := data.Get("users").Index(0).Unmarshal(&firstUser)

// Get all users
var users []User
err := data.Get("users").Unmarshal(&users)
```

### Map and Interface Unmarshaling
```go
// Unmarshal to map
var result map[string]interface{}
err := data.Unmarshal(&result)

// Unmarshal to interface{} (auto-converts)
var value interface{}
err := data.Unmarshal(&value)
```

### Complex Nested Structures
```go
type Company struct {
    Name      string    `json:"name"`
    Address   Address   `json:"address"`
    Employees []Person  `json:"employees"`
    Founded   int       `json:"founded"`
}

data := Parse(complexJSON)
var company Company
err := data.Unmarshal(&company)
```

## API Overview

### Chain-Style Access
```go
data.Get("key")                    // Get object field
data.Index(0)                      // Get array element  
data.Get("users").Index(0)         // Chain operations
```

### Type Checking
```go
data.IsString()                    // Check if string
data.IsArray()                     // Check if array
data.IsValid()                     // Check if valid
data.Exists()                      // Check if exists
```

### Type Conversion
```go
str, ok := data.AsString()         // Safe conversion
age := data.IntOr(25)              // With default value
```

### Immutable Operations
```go
newData := data.Set("key", value)  // Returns new JsonData
newData := data.Append(item)       // Returns new JsonData
newData := data.Remove("key")      // Returns new JsonData
```

### Unmarshal Integration
```go
// Direct unmarshal
data.Unmarshal(&target)

// Chain + unmarshal
data.Get("user").Unmarshal(&user)
data.Get("items").Index(0).Unmarshal(&item)

// Error-safe chain
if data.Get("user").IsValid() {
    data.Get("user").Unmarshal(&user)
}
```

## Error Handling

The library provides comprehensive error handling:
- Type mismatch errors
- Invalid JSON format errors  
- Non-pointer target errors
- Chain operation errors with full context

```go
data, err := Parse(`{"invalid": json}`)
if err != nil {
    // Handle parse error - contains full context
    fmt.Printf("Parse failed: %v", err)
    return
}
var result map[string]interface{}
if err := data.Unmarshal(&result); err != nil {
    // Handle unmarshal error - contains full context
    fmt.Printf("Unmarshal failed: %v", err)
}
```

## Testing

Comprehensive test coverage includes:
- ✅ Basic type unmarshaling
- ✅ Complex struct unmarshaling  
- ✅ Nested structure support
- ✅ Array and slice handling
- ✅ Map unmarshaling
- ✅ Interface{} auto-conversion
- ✅ Error condition testing
- ✅ Chain operation integration

All tests pass with 100% success rate.

## Performance & Design

### Advantages over encoding/json:
- **Chain-style access** before unmarshaling
- **No intermediate conversions** required
- **Integrated error handling** in the chain
- **More flexible access patterns**
- **Type-safe default values**

### Comparison Example:

**Standard Library (verbose):**
```go
var data interface{}
json.Unmarshal([]byte(jsonStr), &data)
userData := data.(map[string]interface{})["users"]
firstUser := userData.([]interface{})[0]
userBytes, _ := json.Marshal(firstUser)
var user User
json.Unmarshal(userBytes, &user)
```

**Aaron JSON (concise):**
```go
data, err := Parse(jsonStr)
if err != nil {
    // Handle parse error
    return
}
var user User
err = data.Get("users").Index(0).Unmarshal(&user)
```

## Implementation Status

✅ **Complete Feature Set:**
- Core Unmarshal implementation
- Struct field mapping with json tags  
- Basic and complex type conversions
- Error handling and propagation
- Chain operation compatibility
- Comprehensive test coverage
- Documentation and examples
- Full integration with existing interface

The Unmarshal functionality is **production-ready** and provides a modern, ergonomic approach to JSON processing in Go.
