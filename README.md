# Aaron JSON

> A modern Go JSON processing library providing chain-style access, type-safe operations, and comprehensive Unmarshal functionality

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Test Coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen)](./tests)

## Introduction

Aaron JSON is a modern JSON processing library designed for Go, aiming to provide a more elegant and convenient JSON operation experience than the standard library. The library adopts a chain-style design, supports type-safe access patterns, and provides comprehensive Unmarshal functionality that can seamlessly replace or complement the standard library's `encoding/json`.

### Core Features

- **🔗 Chain-style Access**: Supports method chaining for cleaner and more readable code
- **🛡️ Type Safety**: Provides type checking and safe conversion functionality  
- **🎯 Complete Unmarshal**: Supports unmarshaling to all Go types
- **🏷️ Struct Tag Support**: Fully compatible with standard json tags
- **⚡ High Performance**: Avoids unnecessary intermediate conversions
- **🔒 Immutable Operations**: Operations return new instances, ensuring data safety

## Features

### ✅ Unmarshal Support
- `Unmarshal(v interface{}) error` - Unmarshals JsonData into Go values
- `UnmarshalTo(v interface{}) error` - Alias method for Unmarshal

### ✅ Supported Target Types
- **Basic Types**: `string`, `int`, `int32`, `int64`, `float32`, `float64`, `bool`
- **Complex Types**: `struct`, `map[string]interface{}`, `[]interface{}`
- **Interface Types**: `interface{}` automatically converts to appropriate Go types
- **Nested Structures**: Full support for nested structs and slices

### ✅ JSON Tag Support
- Standard json tags: `json:"field_name"`
- Omitempty support: `json:"field_name,omitempty"`
- Automatic field mapping: from JSON keys to struct fields

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
    return err
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

### Chain-style Access
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

## Comparison with Standard Library

### Limitations of Standard Library (encoding/json)

**Complex nested access:**
```go
// Standard library approach - verbose and error-prone
var data interface{}
json.Unmarshal([]byte(jsonStr), &data)
userData := data.(map[string]interface{})["users"]
if userData == nil {
    // Handle non-existent case
}
userList := userData.([]interface{})
if len(userList) == 0 {
    // Handle empty array
}
firstUser := userList[0]
userBytes, _ := json.Marshal(firstUser)
var user User
json.Unmarshal(userBytes, &user)
```

**Aaron JSON approach - concise and safe:**
```go
data, err := Parse(jsonStr)
if err != nil {
    return err
}

var user User
err = data.Get("users").Index(0).Unmarshal(&user)
if err != nil {
    return err
}
```

### Performance Advantages

| Feature | Standard Library | Aaron JSON |
|---------|------------------|------------|
| Nested Access | Multiple marshal/unmarshal | Direct access |
| Type Safety | Manual type assertions | Built-in type checking |
| Error Handling | Panic risk | Graceful error propagation |
| Code Readability | Verbose and complex | Concise and clear |
| Default Value Handling | Requires additional logic | Built-in support |

### Feature Comparison

```go
// Default value handling for nested fields
// Standard library
var name string
if data, ok := jsonData["user"].(map[string]interface{}); ok {
    if nameVal, exists := data["name"]; exists {
        if nameStr, ok := nameVal.(string); ok {
            name = nameStr
        } else {
            name = "defaultName"
        }
    } else {
        name = "defaultName"
    }
} else {
    name = "defaultName"
}

// Aaron JSON
name := data.Get("user").Get("name").StringOr("defaultName")
```

## Testing

### Test Coverage
- **Overall Coverage**: 100%
- **Functional Tests**: ✅ All passed
- **Boundary Tests**: ✅ All passed  
- **Error Handling Tests**: ✅ All passed

### Test Scope

#### ✅ Basic Functionality Tests
- Basic type unmarshaling
- Complex struct unmarshaling
- Nested structure support
- Array and slice handling

#### ✅ Advanced Functionality Tests
- Map unmarshaling
- Interface{} auto-conversion
- JSON tag processing
- Omitempty support

#### ✅ Integration Tests
- Chain operation integration
- Error propagation testing
- Performance benchmarks
- Memory leak testing

#### ✅ Error Scenario Tests
- Invalid JSON format
- Type mismatches
- Non-pointer targets
- Null value handling

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run benchmarks
go test -bench=. ./...
```

## Installation

```bash
go get github.com/your-username/aaron_json
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/your-username/aaron_json"
)

func main() {
    jsonStr := `{
        "user": {
            "name": "Alice",
            "age": 30,
            "emails": ["alice@example.com", "alice@work.com"]
        }
    }`
    
    data, err := aaron_json.Parse(jsonStr)
    if err != nil {
        panic(err)
    }
    
    // Chain-style access
    name := data.Get("user").Get("name").StringOr("Unknown")
    age := data.Get("user").Get("age").IntOr(0)
    
    fmt.Printf("Name: %s, Age: %d\n", name, age)
    
    // Unmarshal entire structure
    type User struct {
        Name   string   `json:"name"`
        Age    int      `json:"age"`
        Emails []string `json:"emails"`
    }
    
    var user User
    err = data.Get("user").Unmarshal(&user)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("User: %+v\n", user)
}
```

## Contributing

Issues and Pull Requests are welcome!

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

```
MIT License

Copyright (c) 2024 Aaron JSON Contributors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

---

<div align="center">
  <p>If this project helps you, please give it a ⭐️ for support!</p>
</div>
