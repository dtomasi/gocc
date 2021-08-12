package gocc

// Check Regular expressions
const (
	RegexCheckSnackCase      = `^[a-z][a-z0-9]+(_[a-z0-9]+)*$`
	RegexCheckUpperSnakeCase = `^[A-Z][A-Z0-9]+(_[A-Z0-9]+)*$`
	RegexCheckPascalCase     = `^[A-Z][a-z0-9]+([A-Za-z0-9]+)*$`
	RegexCheckCamelCase      = `^[a-z0-9]+([A-Za-z0-9]+)*$`
	RegexCheckKebabCase      = `^[a-z][a-z0-9]+(-[a-z0-9]+)*$`
	RegexCheckUpperKebabCase = `^[A-Z][A-Z0-9]+(-[A-Z0-9]+)*$`
)
