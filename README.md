
# Go String Package

The Go String package provides a wrapper around Go's built-in string type, offering a chainable API for string manipulation. This makes your string operations more fluent and intuitive.

## Installation

To use the String package, first, ensure you have Go installed on your system. Then, import the package into your project:

```go
import . "path/to/String"
```

Replace `"path/to/String"` with the actual path where the package is located.

## Types

### String

The `String` struct is the core type of the String package, wrapping a pointer to a Go `string`.

## Functions

### `S(s any) String`

Initializes a new `String` instance from any value by converting it to a string using `fmt.Sprint(s)`.

### Methods

#### `Contains(substr string) bool`

Checks if the `String` contains the specified substring.

#### `IncludesAny(substrs ...string) bool`

Determines if the `String` includes any of the specified substrings.

#### `Len() int`

Returns the length of the `String`.

#### `Replace(old string, nw string, count ...int) String`

Replaces the first `n` occurrences of `old` with `new`. If `count` is not specified, defaults to 1.

#### `ReplaceAll(oldnew ...string) String`

Replaces all occurrences of `old` with `new`. Parameters are passed as a variadic list where the first element is `old` and the second is `new`.

#### `ToLower() String`

Converts all characters in the `String` to lowercase.

#### `ToUpper() String`

Converts all characters in the `String` to uppercase.

#### `TrimSpace() String`

Trims whitespace from both ends of the `String`.

## Example Usage

```go
package main

import (
  "fmt"
  . "path/to/String" // Replace with the actual import path
)

func main() {
  s := String.S(" Hello, WORLD!!! ")
  fmt.Println(s.TrimSpace().ToLower().ReplaceAll("world", "gopher").Value) // Output: "hello, gopher!!!"
}
```

## Contributing

Contributions to the String package are welcome. Please feel free to submit issues or pull requests to the project repository.

## License

The String package is licensed under the MIT License. See the LICENSE file in the project repository for more details.
