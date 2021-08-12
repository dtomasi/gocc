package gocc

// Check Regular expressions
const (
	regexCheckSnackCase        = `^[a-z][a-z0-9]+(_[a-z0-9]+)*$`
	regexCheckUpperSnakeCase   = `^[A-Z][A-Z0-9]+(_[A-Z0-9]+)*$`
	regexCheckPascalCase       = `^[A-Z][a-z0-9]+([A-Za-z0-9]+)*$`
	regexCheckCamelCase        = `^[a-z0-9]+([A-Za-z0-9]+)*$`
	regexCheckKebabCase        = `^[a-z][a-z0-9]+(-[a-z0-9]+)*$`
	regexCheckUpperKebabCase   = `^[A-Z][A-Z0-9]+(-[A-Z0-9]+)*$`
	regexCheckDotNotation      = `^[a-z][a-z0-9]+(\.[a-z0-9]+)*$`
	regexCheckUpperDotNotation = `^[A-Z][A-Z0-9]+(\.[A-Z0-9]+)*$`
)
