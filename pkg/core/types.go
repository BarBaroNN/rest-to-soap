package core

import (
	"text/template"
)

// TemplateInfo represents a parsed template with its metadata
type TemplateInfo struct {
	Name     string
	Content  string
	Template *template.Template
	Paths    []string // Extracted paths from template
}

// Registry manages template loading and caching
type Registry interface {
	// LoadTemplate loads a template from a file
	LoadTemplate(name, path string) error

	// GetTemplate returns a template by name
	GetTemplate(name string) (*TemplateInfo, error)

	// ExecuteTemplate executes a template with the given data
	ExecuteTemplate(name string, data interface{}) (string, error)

	// ListTemplates returns a list of all loaded templates
	ListTemplates() []string
}

// Parser handles template parsing and execution
type Parser interface {
	// ParseTemplate parses a template string into a TemplateInfo
	ParseTemplate(name, content string) (*TemplateInfo, error)

	// ExtractPaths extracts all paths from a template
	ExtractPaths(content string) ([]string, error)

	// ExecuteTemplate executes a template with the given data
	ExecuteTemplate(tmpl *TemplateInfo, data interface{}) (string, error)
}

// Cache provides template caching functionality
type Cache interface {
	// Get retrieves a template from cache
	Get(name string) (*TemplateInfo, bool)

	// Set stores a template in cache
	Set(name string, tmpl *TemplateInfo)

	// Delete removes a template from cache
	Delete(name string)

	// Clear removes all templates from cache
	Clear()

	// List returns all template names in the cache
	List() []string
}
