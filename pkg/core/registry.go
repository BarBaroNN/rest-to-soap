package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// DefaultRegistry implements the Registry interface
type DefaultRegistry struct {
	parser Parser
	cache  Cache
	mu     sync.RWMutex
}

// NewRegistry creates a new template registry
func NewRegistry(parser Parser, cache Cache) *DefaultRegistry {
	return &DefaultRegistry{
		parser: parser,
		cache:  cache,
	}
}

// LoadTemplate loads a template from a file
func (r *DefaultRegistry) LoadTemplate(name, path string) error {
	// Read template file
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read template file: %w", err)
	}

	// Parse template
	tmpl, err := r.parser.ParseTemplate(name, string(content))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Store in cache
	r.cache.Set(name, tmpl)

	return nil
}

// GetTemplate returns a template by name
func (r *DefaultRegistry) GetTemplate(name string) (*TemplateInfo, error) {
	// Try to get from cache
	if tmpl, ok := r.cache.Get(name); ok {
		return tmpl, nil
	}

	return nil, fmt.Errorf("template %s not found", name)
}

// ExecuteTemplate executes a template with the given data
func (r *DefaultRegistry) ExecuteTemplate(name string, data interface{}) (string, error) {
	// Get template
	tmpl, err := r.GetTemplate(name)
	if err != nil {
		return "", err
	}

	// Execute template
	return r.parser.ExecuteTemplate(tmpl, data)
}

// ListTemplates returns a list of all loaded templates
func (r *DefaultRegistry) ListTemplates() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Get all templates from cache
	templates := make([]string, 0)
	for _, name := range r.cache.List() {
		templates = append(templates, name)
	}

	return templates
}

// LoadTemplatesFromDir loads all templates from a directory
func (r *DefaultRegistry) LoadTemplatesFromDir(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-template files
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".tmpl") {
			return nil
		}

		// Get template name from filename
		name := strings.TrimSuffix(info.Name(), ".tmpl")

		// Load template
		if err := r.LoadTemplate(name, path); err != nil {
			return fmt.Errorf("failed to load template %s: %w", name, err)
		}

		return nil
	})
}
