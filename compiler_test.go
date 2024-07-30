package main

import (
	"testing"
)

func TestParseComponents(t *testing.T) {
	tests := []struct {
		name       string
		components map[string]string
		input      string
		expected   string
	}{
		{
			name: "Simple Tag",
			components: map[string]string{
				"quilt-component": "<div>This is a quilt component</div>",
			},
			input:    `<div data-component-name="quilt-component"></div>`,
			expected: `<div>This is a quilt component</div>`,
		},
		{
			name: "Tag With One Attribute",
			components: map[string]string{
				"quilt-component": "<div>This is a quilt component with a count of {count}</div>",
			},
			input:    `<div data-component-name="quilt-component" data-attr-count="4"></div>`,
			expected: `<div>This is a quilt component with a count of 4</div>`,
		},
		{
			name: "Tag With Multiple Attributes",
			components: map[string]string{
				"quilt-component": "<div>This is a quilt component with a count of {count} and a name of {name}</div>",
			},
			input:    `<div data-component-name="quilt-component" data-attr-count="4" data-attr-name="quilt"></div>`,
			expected: `<div>This is a quilt component with a count of 4 and a name of quilt</div>`,
		},
		{
			name: "Tag With Multiple Components",
			components: map[string]string{
				"quilt-component": "<div>This is a quilt component with a count of {count} and a name of {name}</div>",
				"bear-component":  "<div>This is a bear component with a count of {count} and a name of {name}</div>",
			},
			input:    `<div data-component-name="quilt-component" data-attr-count="4" data-attr-name="quilt"></div><div data-component-name="bear-component" data-attr-count="2" data-attr-name="Bear"></div>`,
			expected: `<div>This is a quilt component with a count of 4 and a name of quilt</div><div>This is a bear component with a count of 2 and a name of Bear</div>`,
		},
		{
			name: "Tag With Multiple Components and Text",
			components: map[string]string{
				"quilt-component": "<div>This is a quilt component with a count of {count} and a name of {name}</div>",
				"bear-component":  "<div>This is a bear component with a count of {count} and a name of {name}</div>",
			},
			input:    `<div data-component-name="quilt-component" data-attr-count="4" data-attr-name="quilt"></div>Some text<div data-component-name="bear-component" data-attr-count="2" data-attr-name="Bear"></div>`,
			expected: `<div>This is a quilt component with a count of 4 and a name of quilt</div>Some text<div>This is a bear component with a count of 2 and a name of Bear</div>`,
		},
		{
			name: "Tag With Nested Components",
			components: map[string]string{
				"quilt-component": `<div>This is a quilt component with a count of {count} and a name of {name}<div data-component-name="bear-component" data-attr-count="2" data-attr-name="Bear"></div></div>`,
				"bear-component":  "<div>This is a bear component with a count of {count} and a name of {name}</div>",
			},
			input:    `<div data-component-name="quilt-component" data-attr-count="4" data-attr-name="quilt"></div>`,
			expected: `<div>This is a quilt component with a count of 4 and a name of quilt<div>This is a bear component with a count of 2 and a name of Bear</div></div>`,
		},
		{
			name: "Tag With Nested Components and Text",
			components: map[string]string{
				"quilt-component": `<div>This is a quilt component with a count of {count} and a name of {name}<div data-component-name="bear-component" data-attr-count="2" data-attr-name="Bear"></div></div>`,
				"bear-component":  "<div>This is a bear component with a count of {count} and a name of {name}</div>",
			},
			input:    `<div data-component-name="quilt-component" data-attr-count="4" data-attr-name="quilt"></div>Some text<div data-component-name="bear-component" data-attr-count="2" data-attr-name="Bear"></div>`,
			expected: `<div>This is a quilt component with a count of 4 and a name of quilt<div>This is a bear component with a count of 2 and a name of Bear</div></div>Some text<div>This is a bear component with a count of 2 and a name of Bear</div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseComponents(tt.input, tt.components)
			if got != tt.expected {
				t.Errorf("parseComponents() = %v, want %v", got, tt.expected)
			}
		})
	}
}
