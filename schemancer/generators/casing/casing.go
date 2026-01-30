package casing

import (
	"strings"
	"unicode"
)

// CommonAcronyms is a lookup table of commonly used acronyms that should be uppercased
var CommonAcronyms = map[string]string{
	"id":   "ID",
	"api":  "API",
	"http": "HTTP",
	"https": "HTTPS",
	"url":  "URL",
	"uri":  "URI",
	"uuid": "UUID",
	"sql":  "SQL",
	"json": "JSON",
	"xml":  "XML",
	"html": "HTML",
	"css":  "CSS",
	"rpc":  "RPC",
	"rest": "REST",
	"tcp":  "TCP",
	"udp":  "UDP",
	"ip":   "IP",
	"db":   "DB",
	"ui":   "UI",
	"ux":   "UX",
	"io":   "IO",
	"os":   "OS",
	"cpu":  "CPU",
	"gpu":  "GPU",
	"ram":  "RAM",
	"ascii": "ASCII",
	"utf":  "UTF",
	"oauth": "OAuth",
	"jwt":  "JWT",
	"ssl":  "SSL",
	"tls":  "TLS",
	"ssh":  "SSH",
	"ftp":  "FTP",
	"dns":  "DNS",
	"smtp": "SMTP",
	"imap": "IMAP",
	"pop":  "POP",
	"grpc": "GRPC",
	"cors": "CORS",
	"csrf": "CSRF",
	"xss":  "XSS",
}

// ToPascalCase converts a string to PascalCase with proper acronym handling
func ToPascalCase(s string) string {
	words := SplitWords(s)
	for i, w := range words {
		if len(w) > 0 {
			lower := strings.ToLower(w)
			if acronym, ok := CommonAcronyms[lower]; ok {
				words[i] = acronym
			} else {
				words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
			}
		}
	}
	return strings.Join(words, "")
}

// ToCamelCase converts a string to camelCase with proper acronym handling
func ToCamelCase(s string) string {
	pascal := ToPascalCase(s)
	if len(pascal) == 0 {
		return ""
	}

	// Handle case where first word is an acronym
	// For example: "ID" -> "id", "API" -> "api"
	// But preserve mixed case: "IDField" -> "idField"
	firstRune := rune(pascal[0])
	if len(pascal) > 1 {
		secondRune := rune(pascal[1])
		// If first two chars are both uppercase (likely acronym at start)
		// lowercase the entire first word/acronym
		if unicode.IsUpper(secondRune) {
			// Find where the acronym ends
			i := 1
			for i < len(pascal) && unicode.IsUpper(rune(pascal[i])) {
				i++
			}
			// If we're not at the end and next char is lowercase,
			// keep the last uppercase char with the next word
			// e.g., "HTTPServer" -> "httpServer" not "httPserver"
			if i < len(pascal) && i > 1 {
				i--
			}
			return strings.ToLower(pascal[:i]) + pascal[i:]
		}
	}

	return strings.ToLower(string(firstRune)) + pascal[1:]
}

// ToSnakeCase converts a string to snake_case
func ToSnakeCase(s string) string {
	words := SplitWords(s)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return strings.Join(words, "_")
}

// ToKebabCase converts a string to kebab-case
func ToKebabCase(s string) string {
	words := SplitWords(s)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return strings.Join(words, "-")
}

// SplitWords splits a string into words based on delimiters and case transitions
func SplitWords(s string) []string {
	var words []string
	var current strings.Builder

	for i, r := range s {
		if r == '_' || r == '-' || r == ' ' {
			if current.Len() > 0 {
				words = append(words, current.String())
				current.Reset()
			}
			continue
		}

		if unicode.IsUpper(r) && i > 0 {
			prev := rune(s[i-1])
			if unicode.IsLower(prev) {
				if current.Len() > 0 {
					words = append(words, current.String())
					current.Reset()
				}
			}
		}

		current.WriteRune(r)
	}

	if current.Len() > 0 {
		words = append(words, current.String())
	}

	return words
}
