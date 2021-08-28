# gocc (Golang Case Converter)

[![CodeFactor](https://www.codefactor.io/repository/github/dtomasi/gocc/badge)](https://www.codefactor.io/repository/github/dtomasi/di)
[![pre-commit.ci status](https://results.pre-commit.ci/badge/github/dtomasi/gocc/main.svg)](https://results.pre-commit.ci/latest/github/dtomasi/gocc/main)
![Go Unit Tests](https://github.com/dtomasi/gocc/actions/workflows/build.yml/badge.svg)
![CodeQL](https://github.com/dtomasi/gocc/actions/workflows/codeql-analysis.yml/badge.svg)

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
    // Input case style is automatically detected
    gocc.C("myCamelCaseString").ToSnakeCase()

    // Convert with known input type
    // Using Convert function with known type is much faster than the example above,
    // because there is no need to detect the input type
    gocc.C("myCamelCaseString").Convert(gocc.StyleCamelCase, gocc.StyleSnakeCase)

    // Validate strings -> returns true
    gocc.IsSnakeCase("my_snake_case_string")

    // Case type checking -> returns StyleKebabCase (int)
    gocc.C("kebab-case-string").Style()
    // returns "kebab"
    gocc.C("kebab-case-string").Style().String()
}

```

For a full list of available functions see documentation
