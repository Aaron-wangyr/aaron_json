# Aaron JSON

> A Go JSON processing library providing unified object operations

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## Introduction

Aaron JSON is a JSON processing library designed for Go, aiming to provide a more elegant and convenient JSON operation experience. 

## Quick Start

## Installation

### Option 1: From GitHub (Recommended)
```bash
go get -u github.com/Aaron-wangyr/aaron_json
```

### Option 2: Local Development
```bash
# Clone the repository
git clone https://github.com/Aaron-wangyr/aaron_json.git
cd aaron_json

# Use as local module
go mod tidy
```

### Option 3: As a Go Module
Add to your `go.mod`:
```go
require github.com/Aaron-wangyr/aaron_json v1.0.0
```

```go
package main

import (
    "fmt"
    aaronjson "github.com/Aaron-wangyr/aaron_json"
)

func main() {
    jsonStr := `{
        "user": {
            "name": "Alice",
            "age": 30,
            "emails": ["alice@example.com", "alice@work.com"]
        }
    }`
    var err error
    data, err := aaronjson.Parse(jsonStr)
    if err != nil {
        panic(err)
    }
    
    // Chain-style access
    name, err := data.Get("user", "name")
    if err != nil {
        panic(err)
    }
    age, err := data.Get("user", "age")
    if err != nil {
        panic(err)
    }
    
    if name.IsString() {
        nameStr, _ := name.AsString()
        fmt.Printf("Name: %s\n", nameStr)
    }
    if age.IsInt() {
        ageInt, _ := age.AsInt()
        fmt.Printf("Age: %.0f\n", ageInt)
    }
    
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

## More Examples

### Creating JSON from Go values

```go
// Create JSON from struct
user := User{
    Name:   "Bob",
    Age:    25,
    Emails: []string{"bob@example.com"},
}

jsonValue, err := aaronjson.Marshal(user)
if err != nil {
    panic(err)
}

fmt.Println(jsonValue.String()) // Output JSON string
```

### Working with Arrays

```go
jsonStr := `["apple", "banana", "cherry"]`
data, _ := aaronjson.Parse(jsonStr)

if data.IsArray() {
    arr, _ := data.AsArray()
    
    // Access by index
    first, _ := arr.Index(0)
    fmt.Println(first.String()) // "apple"
    
    // Append new item
    arr.Append(aaronjson.NewJsonString("date"))
    
    // Get length
    fmt.Printf("Array length: %d\n", arr.Length())
}
```

### Working with Objects

```go
data, _ := aaronjson.Parse(`{"name": "Alice", "age": 30}`)

if data.IsObject() {
    obj, _ := data.AsObject()
    
    // Set new key-value
    obj.Set("city", aaronjson.NewJsonString("New York"))
    
    // Get all keys
    keys := obj.Keys()
    fmt.Printf("Keys: %v\n", keys)
    
    // Remove key
    obj.Remove("age")
}
```

## API Reference

### Main Functions

- `Parse(jsonStr string) (JsonValue, error)` - Parse JSON string
- `ParseByte(jsonData []byte) (JsonValue, error)` - Parse JSON bytes  
- `Marshal(v interface{}) (JsonValue, error)` - Convert Go value to JSON

### JsonValue Interface Methods

- Type checking: `IsString()`, `IsInt()`, `IsFloat()`, `IsBool()`, `IsNull()`, `IsArray()`, `IsObject()`
- Type conversion: `AsString()`, `AsInt()`, `AsFloat()`, `AsBool()`, `AsArray()`, `AsObject()`
- Access methods: `Get(keys ...string)`, `Index(i int)`, `Length()`, `Keys()`
- Serialization: `String()`, `PrettyString()`, `Unmarshal(v interface{})`

## Usage in Your Project

Once installed, you can use it in your Go project:

```go
package main

import (
    "fmt"
    "log"
    
    aaronjson "github.com/Aaron-wangyr/aaron_json"
)

func main() {
    // Example: Processing API response
    apiResponse := `{
        "status": "success",
        "data": {
            "users": [
                {"id": 1, "name": "Alice", "active": true},
                {"id": 2, "name": "Bob", "active": false}
            ]
        }
    }`
    
    data, err := aaronjson.Parse(apiResponse)
    if err != nil {
        log.Fatal(err)
    }
    
    // Check API status
    status, _ := data.Get("status")
    if status.String() == "\"success\"" {
        fmt.Println("API call successful!")
        
        // Process users
        users, _ := data.Get("data", "users")
        if users.IsArray() {
            userArray, _ := users.AsArray()
            for i := 0; i < userArray.Length(); i++ {
                user, _ := userArray.Index(i)
                name, _ := user.Get("name")
                active, _ := user.Get("active")
                
                nameStr, _ := name.AsString()
                activeBool, _ := active.AsBool()
                
                fmt.Printf("User: %s, Active: %v\n", nameStr, activeBool)
            }
        }
    }
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
