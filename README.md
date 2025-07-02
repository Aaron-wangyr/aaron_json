# Aaron JSON

> A Go JSON processing library providing unified object operations

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## Introduction

Aaron JSON is a JSON processing library designed for Go, aiming to provide a more elegant and convenient JSON operation experience. 

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
    var err error
    data, err := aaron_json.Parse(jsonStr)
    if err != nil {
        panic(err)
    }
    
    // Chain-style access
    name,err := data.Get("user","name")
    age,err := data.Get("user","age")
    if name.IsString(){
        nameStr,_:=name.AsString()
        fmt.Printf("Name: %s",nameStr)
        fmt.Printf("Name: %s",name.String())
    }
    if age.IsNumber(){
        
    }
    fmt.Printf("Name: %s, Age: %d\n", name., age)
    
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
