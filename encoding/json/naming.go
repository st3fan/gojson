package json

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

type NamingStyle int

const (
	GoNaming         NamingStyle = 0
	JavascriptNaming             = 1
	PythonNaming                 = 2
	CSharpNaming                 = 3
)

func splitName(name string) []string {
	re := regexp.MustCompile("([A-Z][a-z0-9]*)")
	parts := re.FindAllString(name, -1)

	// Because I'm stupid and I can't get the regexp to do all the
	// work in one go, I am doing this partly manually.

	result := []string{}

	for i := 0; i < len(parts); {
		t := ""
		if len(parts[i]) == 1 {
			for i < len(parts) && len(parts[i]) == 1 {
				t += parts[i]
				i += 1
			}
			result = append(result, t)
		} else {
			result = append(result, parts[i])
			i += 1
		}
	}

	return result
}

func isAllUpper(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func removeCapitalization(s string) string {
	if s == "" {
		return ""
	}
	if isAllUpper(s) {
		return strings.ToLower(s)
	} else {
		r, n := utf8.DecodeRuneInString(s)
		return string(unicode.ToLower(r)) + s[n:]
	}
}

func encodeJavascriptName(name string) string {
	parts := splitName(name)
	parts[0] = removeCapitalization(parts[0])
	return strings.Join(parts, "")
}

func encodeCSharpName(name string) string {
	parts := splitName(name)
	parts[0] = removeCapitalization(parts[0])

	// Turn NewFOOThing into NewFooThing
	if len(parts) > 1 {
		for i := 1; i < len(parts); i++ {
			if isAllUpper(parts[i]) {
				parts[i] = strings.Title(strings.ToLower(parts[i]))
			}
		}
	}

	return strings.Join(parts, "")
}

func encodePythonName(name string) string {
	parts := splitName(name)
	for i, _ := range parts {
		parts[i] = strings.ToLower(parts[i])
	}
	return strings.Join(parts, "_")
}

func changeFieldNaming(name string, style NamingStyle) string {
	switch style {
	case JavascriptNaming:
		return encodeJavascriptName(name)
	case CSharpNaming:
		return encodeCSharpName(name)
	case PythonNaming:
		return encodePythonName(name)
	}
	return name
}
