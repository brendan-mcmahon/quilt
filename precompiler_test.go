package main

import (
	"testing"
)

func TestConvertCustomSyntaxToHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty Tag",
			input:    `<[quilt-component]></[quilt-component]>`,
			expected: `<div data-component-name="quilt-component"></div>`,
		},
		{
			name:     "Self-Closing Tag",
			input:    `<[quilt-component] />`,
			expected: `<div data-component-name="quilt-component"></div>`,
		},
		{
			name:     "Tag With One Attribute",
			input:    `<[quilt-component] {name}="something"></[quilt-component]>`,
			expected: `<div data-component-name="quilt-component" data-attr-name="something"></div>`,
		},
		{
			name:     "Self-Closing Tag With One Attribute",
			input:    `<[quilt-component] {name}="something" />`,
			expected: `<div data-component-name="quilt-component" data-attr-name="something"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToDataAttributes(tt.input); got != tt.expected {
				t.Errorf("convertCustomSyntaxToHTML() = %v, want %v", got, tt.expected)
			}
		})
	}
}
