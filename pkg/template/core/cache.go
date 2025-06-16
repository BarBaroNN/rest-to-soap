package core

import (
	"sync"
)

// MemoryCache implements the Cache interface using an in-memory map
type MemoryCache struct {
	templates map[string]*TemplateInfo
	mu        sync.RWMutex
}

// NewMemoryCache creates a new in-memory cache
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		templates: make(map[string]*TemplateInfo),
	}
}

// Get retrieves a template from cache
func (c *MemoryCache) Get(name string) (*TemplateInfo, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	tmpl, ok := c.templates[name]
	return tmpl, ok
}

// Set stores a template in cache
func (c *MemoryCache) Set(name string, tmpl *TemplateInfo) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.templates[name] = tmpl
}

// Delete removes a template from cache
func (c *MemoryCache) Delete(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.templates, name)
}

// Clear removes all templates from cache
func (c *MemoryCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.templates = make(map[string]*TemplateInfo)
}

// List returns all template names in the cache
func (c *MemoryCache) List() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	names := make([]string, 0, len(c.templates))
	for name := range c.templates {
		names = append(names, name)
	}
	return names
}
