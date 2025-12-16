# go-strings

A collection of useful string manipulation functions for Go with Unicode support.

## Description

The `go-strings` package provides additional string manipulation functions that are not included in Go's standard library. All functions correctly handle Unicode characters and are optimized for performance.

## Installation

```bash
go get github.com/ayden1st/go-strings
```

## Functions

### Reverse(s string) string

Returns a new string with the characters of the input string in reverse order. Uses a rune slice to ensure correct handling of Unicode characters.

**Example:**
```go
package main

import (
    "fmt"
    "github.com/ayden1st/go-strings"
)

func main() {
    fmt.Println(go_strings.Reverse("hello"))        // Output: olleh
    fmt.Println(go_strings.Reverse("привет"))       // Output: тевирп
    fmt.Println(go_strings.Reverse("👋🌍"))          // Output: 🌍👋
}
```

### StripNonAlphanumeric(s string) string

Removes all characters from the input string except letters (Unicode), digits (Unicode), and spaces. Uses strings.Builder for efficient string construction.

**Example:**
```go
package main

import (
    "fmt"
    "github.com/ayden1st/go-strings"
)

func main() {
    fmt.Println(go_strings.StripNonAlphanumeric("Hello, world! 123"))  // Output: Hello world 123
    fmt.Println(go_strings.StripNonAlphanumeric("Привет, мир! 456"))   // Output: Привет мир 456
    fmt.Println(go_strings.StripNonAlphanumeric("@#$%^&*()"))           // Output: 
}
```

### StripBlanks(s string) string

Removes all spaces (' ') from the input string.

**Example:**
```go
package main

import (
    "fmt"
    "github.com/ayden1st/go-strings"
)

func main() {
    fmt.Println(go_strings.StripBlanks("hello world test"))    // Output: helloworldtest
    fmt.Println(go_strings.StripBlanks("  hello  world  "))    // Output: helloworld
    fmt.Println(go_strings.StripBlanks("привет мир"))          // Output: приветмир
}
```

## Testing

To run the tests:

```bash
cd go-strings
go test -v
```

To run benchmarks:

```bash
go test -bench=.
```

## License

This project is licensed under the MIT License.

## Version

Current version: 0.1.0