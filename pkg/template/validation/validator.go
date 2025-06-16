package validation

import (
	"fmt"

	"rest-to-soap/pkg/template/core"
)

// Validator validates templates against WSDL paths
type Validator struct {
	wsdlPaths map[string][]string
}

// NewValidator creates a new template validator
func NewValidator() *Validator {
	return &Validator{
		wsdlPaths: make(map[string][]string),
	}
}

// AddWSDLPaths adds WSDL paths for validation
func (v *Validator) AddWSDLPaths(operation string, paths []string) {
	v.wsdlPaths[operation] = paths
}

// ValidateTemplate validates a template against WSDL paths
func (v *Validator) ValidateTemplate(name string, tmpl *core.TemplateInfo) ([]ValidationError, error) {
	var errors []ValidationError

	// Get WSDL paths for this operation
	wsdlPaths, ok := v.wsdlPaths[name]
	if !ok {
		return nil, fmt.Errorf("no WSDL paths defined for operation %s", name)
	}

	// Create a map of WSDL paths for quick lookup
	wsdlPathMap := make(map[string]bool)
	for _, path := range wsdlPaths {
		wsdlPathMap[path] = true
	}

	// Check each template path against WSDL paths
	for _, path := range tmpl.Paths {
		if !wsdlPathMap[path] {
			errors = append(errors, ValidationError{
				TemplateName: name,
				Message:      "path not found in WSDL",
				Path:         path,
			})
		}
	}

	// Check for missing required paths
	for _, wsdlPath := range wsdlPaths {
		found := false
		for _, path := range tmpl.Paths {
			if path == wsdlPath {
				found = true
				break
			}
		}
		if !found {
			errors = append(errors, ValidationError{
				TemplateName: name,
				Message:      "required path missing from template",
				Path:         wsdlPath,
			})
		}
	}

	return errors, nil
}

// ValidatePath validates a single path against WSDL paths
func (v *Validator) ValidatePath(operation, path string) error {
	wsdlPaths, ok := v.wsdlPaths[operation]
	if !ok {
		return fmt.Errorf("no WSDL paths defined for operation %s", operation)
	}

	for _, wsdlPath := range wsdlPaths {
		if path == wsdlPath {
			return nil
		}
	}

	return fmt.Errorf("path %s not found in WSDL for operation %s", path, operation)
}

// GetWSDLPaths returns all WSDL paths for an operation
func (v *Validator) GetWSDLPaths(operation string) ([]string, error) {
	paths, ok := v.wsdlPaths[operation]
	if !ok {
		return nil, fmt.Errorf("no WSDL paths defined for operation %s", operation)
	}
	return paths, nil
}
