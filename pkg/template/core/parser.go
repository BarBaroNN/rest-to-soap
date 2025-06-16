package core

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/template"
)

// DefaultParser implements the Parser interface
type DefaultParser struct {
	funcMap template.FuncMap
}

// NewParser creates a new template parser
func NewParser() *DefaultParser {
	return &DefaultParser{
		funcMap: template.FuncMap{
			"join": strings.Join,
		},
	}
}

// ParseTemplate parses a template string into a TemplateInfo
func (p *DefaultParser) ParseTemplate(name, content string) (*TemplateInfo, error) {
	// Create template
	tmpl, err := template.New(name).Funcs(p.funcMap).Parse(content)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	// Extract paths
	paths, err := p.ExtractPaths(content)
	if err != nil {
		return nil, fmt.Errorf("failed to extract paths: %w", err)
	}

	return &TemplateInfo{
		Name:     name,
		Content:  content,
		Template: tmpl,
		Paths:    paths,
	}, nil
}

// ExtractPaths extracts all paths from a template
func (p *DefaultParser) ExtractPaths(content string) ([]string, error) {
	pathsSet := make(map[string]struct{})

	// Match {{ .Field }} and nested like {{ .A.B }}
	re := regexp.MustCompile(`{{\s*\.([a-zA-Z0-9_\.]+)\s*}}`)
	matches := re.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) > 1 {
			path := strings.TrimSpace(match[1])
			pathsSet[path] = struct{}{}
			// Add parent path if nested
			if idx := strings.Index(path, "."); idx > 0 {
				parent := path[:idx]
				pathsSet[parent] = struct{}{}
			}
		}
	}

	// Match all range variables: {{ range $i, $var := .Collection }}
	rangeRe := regexp.MustCompile(`{{\s*range\s+\$[^,]+,\s*\$([a-zA-Z0-9_]+)\s*:=\s*\.([a-zA-Z0-9_]+)\s*}}`)
	rangeMatches := rangeRe.FindAllStringSubmatch(content, -1)
	for _, match := range rangeMatches {
		if len(match) > 2 {
			collection := strings.TrimSpace(match[2])
			pathsSet[collection] = struct{}{}
		}
	}

	// Match nested paths in range blocks: {{$var.Field}}
	nestedRe := regexp.MustCompile(`{{\s*\$([a-zA-Z0-9_]+)\.([a-zA-Z0-9_]+)\s*}}`)
	nestedMatches := nestedRe.FindAllStringSubmatch(content, -1)
	for _, match := range nestedMatches {
		if len(match) > 2 {
			// Find the collection for this variable
			// We assume variable names are unique per range in this template
			varName := match[1]
			field := match[2]
			// Find the collection for this varName
			for _, rangeMatch := range rangeMatches {
				if len(rangeMatch) > 1 && rangeMatch[1] == varName {
					collection := strings.TrimSpace(rangeMatch[2])
					fullPath := collection + "." + field
					pathsSet[fullPath] = struct{}{}
					// Add parent path
					pathsSet[collection] = struct{}{}
				}
			}
		}
	}

	// Convert set to slice
	paths := make([]string, 0, len(pathsSet))
	for k := range pathsSet {
		paths = append(paths, k)
	}
	return paths, nil
}

// ExecuteTemplate executes a template with the given data
func (p *DefaultParser) ExecuteTemplate(tmpl *TemplateInfo, data interface{}) (string, error) {
	var buf bytes.Buffer
	if err := tmpl.Template.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}
	return buf.String(), nil
}
