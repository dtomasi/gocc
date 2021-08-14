package gocc

func toCamelCaseFromCaseStyle(from CaseStyle, s string) string {

	switch from {
	case StylePascalCase:
		return lcFirst(s)
	case StyleSnakeCase:
		return customDelimiterToCamelCase(s, DelimiterSnakeCase)
	case StyleKebabCase:
		return customDelimiterToCamelCase(s, DelimiterKebabCase)
	case StyleDotNotation:
		return customDelimiterToCamelCase(s, DelimiterDotNotation)
	}

	return s
}

func customDelimiterToCamelCase(s string, delimiter string) string {
	return lcFirst(customDelimiterToPascalCase(s, delimiter))
}
