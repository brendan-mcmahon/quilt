package main

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func compileHtml(data string, components map[string]string) string {
	doc, err := html.Parse(strings.NewReader(data))
	if err != nil {
		return ""
	}

	var output bytes.Buffer
	processNode(&output, doc, components)

	return output.String()
}

func parseComponents(data string, components map[string]string) string {
	doc, err := html.ParseFragment(strings.NewReader(data), &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	})
	if err != nil {
		return ""
	}

	var output bytes.Buffer
	for _, n := range doc {
		processNode(&output, n, components)
	}

	return output.String()
}
func processNode(output *bytes.Buffer, n *html.Node, components map[string]string) {
	switch n.Type {
	case html.ElementNode:
		processElementNode(n, output, components)
	case html.TextNode:
		output.WriteString(html.EscapeString(n.Data))
	case html.CommentNode:
		output.WriteString(fmt.Sprintf("<!--%s-->", n.Data))
	case html.DoctypeNode:
		output.WriteString("<!DOCTYPE html>")
	case html.DocumentNode:
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			processNode(output, c, components)
		}
	}
}

func processElementNode(n *html.Node, output *bytes.Buffer, components map[string]string) {
	if name, ok := getAttr(n, "data-component-name"); ok {
		processComponent(output, n, name, components)
	} else {
		startTag := "<" + n.Data
		for _, attr := range n.Attr {
			startTag += fmt.Sprintf(` %s="%s"`, attr.Key, html.EscapeString(attr.Val))
		}
		startTag += ">"
		output.WriteString(startTag)

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			processNode(output, c, components)
		}

		endTag := fmt.Sprintf("</%s>", n.Data)
		output.WriteString(endTag)
	}
}

func processComponent(output *bytes.Buffer, n *html.Node, name string, components map[string]string) {
	if componentHTML, ok := components[name]; ok {
		attrs := extractAttributes(n)
		for k, v := range attrs {
			componentHTML = strings.ReplaceAll(componentHTML, "{"+k+"}", v)
		}
		parsedHTML := parseComponents(componentHTML, components)
		output.WriteString(parsedHTML)
	}
}

func extractAttributes(n *html.Node) map[string]string {
	attrs := make(map[string]string)
	for _, a := range n.Attr {
		if strings.HasPrefix(a.Key, "data-attr-") {
			attrs[strings.TrimPrefix(a.Key, "data-attr-")] = a.Val
		}
	}
	return attrs
}

func getAttr(n *html.Node, key string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val, true
		}
	}
	return "", false
}
