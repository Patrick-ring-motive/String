
# Go String Package

The Go String package provides a wrapper around Go's built-in string type, offering a chainable API for string manipulation. This makes your string operations more fluent and intuitive.

## Installation

To use the String package, first, ensure you have Go installed on your system. Then, import the package into your project:

```go
import . "github.com/Patrick-ring-motive/go-String"
```


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

#### `strings` and `strconv`

All of the methods in the  `strings` and `strconv` packages from the standard library that take a string as input have been implemented following a similar pattern where the first string input parameter is pulled out given a method that takes in the remaining input parameters eg. 

```go
func doAStandardStringOperation(s string,i int)string {}
//becomes
func (s String) doAStandardStringOperation (i int) String {}
```

For more information on the standard string operations, see the standard documentation
* [string](https://go.patrickring.net/strings)

## Example Usage

```go
package main

import (
  "fmt"
  . "github.com/Patrick-ring-motive/go-String" 
)

func main() {
  s := S(" Hello, WORLD!!! ")
  fmt.Println(s.TrimSpace().ToLower().ReplaceAll("world", "gopher").Value) // Output: "hello, gopher!!!"
}
```

## Contributing

Contributions to the String package are welcome. Please feel free to submit issues or pull requests to the project repository.

## License

The String package is licensed under the MIT License. See the LICENSE file in the project repository for more details.
