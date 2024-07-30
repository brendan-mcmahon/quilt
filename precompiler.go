package main

import (
	"fmt"
	"regexp"
	"strings"
)

func convertToDataAttributes(input string) string {
	tagRegex := regexp.MustCompile(`<\[([^\]]+)\](.*?)(/?>)|(</\[([^\]]+)\]>)`)
	attrRegex := regexp.MustCompile(`\{([^}]+)\}="([^"]*)"`)

	return tagRegex.ReplaceAllStringFunc(input, func(match string) string {
		if strings.HasPrefix(match, "</[") {
			return "</div>"
		}

		parts := tagRegex.FindStringSubmatch(match)
		componentName, attributes, selfClosing := parts[1], strings.TrimSpace(parts[2]), parts[3] == "/>"

		attributes = attrRegex.ReplaceAllString(attributes, `data-attr-$1="$2"`)

		result := fmt.Sprintf(`<div data-component-name="%s"`, componentName)
		if attributes != "" {
			result += " " + attributes
		}
		if selfClosing {
			result += "></div>"
		} else {
			result += ">"
		}

		return result
	})
}
