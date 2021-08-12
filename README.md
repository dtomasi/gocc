# gocc (Golang Case Converter)

## Documentation

https://pkg.go.dev/github.com/dtomasi/gocc

## Installation

    go get github.com/dtomasi/gocc

## Usage

```go
import (
    "github.com/dtomasi/gocc"
)

func func main() {
    // Convert string to snake case -> returns "my_camel_case_string"
    gocc.C("myCamelCaseString").ToSnakeCase()
    
    // Validate strings -> returns true
    gocc.IsSnakeCase("my_snake_case_string")
    
    // Case type checking -> returns StyleKebabCase (int)
    gocc.C("kebab-case-string").Style()
    // returns "kebab"
    gocc.C("kebab-case-string").Style().String()
}

```

For a full list of available functions see documentation
