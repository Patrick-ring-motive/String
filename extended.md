
# Extended String Package Documentation

This package extends Go's standard `strings` and `strconv` packages with additional functionality for working with strings, while also introducing new utilities for handling various string-related types in a flexible and type-safe manner.

## Overview

The package introduces a `String` type, along with a suite of functions that operate on this type and the built-in `string` type. It aims to provide a more versatile and powerful toolset for string manipulation by integrating standard library capabilities with custom enhancements.

## Types

### `String`

A wrapper around the built-in `string` type, providing additional methods for manipulation.

### `Strings`

A collection of `String` objects, offering bulk operations and transformations.

### `StringTypes`

An interface that encompasses both the built-in `string` type and the custom `String` type (including their pointer variants), facilitating operations that can accept any of these types.

## Key Functions and Methods

- `UnwrapStr[STR StringTypes](s STR) string`: Unwraps a `StringTypes` variable, extracting the underlying string value.
- `NewString[STR StringTypes](s STR) String`: Constructs a new `String` instance from a `StringTypes` variable.
- `NewStrings(ss []string) Strings`: Constructs a `Strings` collection from a slice of `string`.
- `S(s any) String`: Creates a `String` instance from any type by converting it to a string.

## Extended Standard Library Functions

This package includes all existing functions from the Go standard library's `strings` and `strconv` packages that interact with strings. For documentation on these functions, please refer to the standard library documentation.

## Custom Additions

Several new functions have been added to further extend the functionality provided by the standard library. These include:

- Custom parsing and formatting functions for handling complex types as strings.
- Enhanced comparison and search functions tailored to the `String` and `Strings` types.
- Additional constructors and utilities for working with `String` and `Strings` types more effectively.

## Usage

```go
// Example of creating a new String
str := String.NewString("example")

// Using an extended standard library function
modifiedStr := str.ToUpper()

// Using a custom addition
result := modifiedStr.CustomFunction()
```

## Installation

To use this package, import it into your Go project:

```go
import "github.com/Patrick-ring-motive/go-String"
```

Replace `"github.com/your-username/your-repo-name/String"` with the actual path to the package.

## Contributing

Contributions to the package are welcome. Please submit issues or pull requests to the project repository on GitHub.

## License

This package is licensed under the MIT License. See the LICENSE file in the project repository for more details.
