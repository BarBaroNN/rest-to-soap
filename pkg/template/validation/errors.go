package validation

import (
	"fmt"
)

// ValidationError represents a template validation error
type ValidationError struct {
	TemplateName string
	Line         int
	Message      string
	Path         string
	Expected     string
	Found        string
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("template %s: %s (line %d)", e.TemplateName, e.Message, e.Line)
}

// FormatValidationErrors formats a list of validation errors
func FormatValidationErrors(errors []ValidationError) string {
	if len(errors) == 0 {
		return "No validation errors"
	}

	var result string
	for _, err := range errors {
		result += fmt.Sprintf("- %s\n", err.Error())
		if err.Path != "" {
			result += fmt.Sprintf("  Path: %s\n", err.Path)
		}
		if err.Expected != "" {
			result += fmt.Sprintf("  Expected: %s\n", err.Expected)
		}
		if err.Found != "" {
			result += fmt.Sprintf("  Found: %s\n", err.Found)
		}
	}
	return result
}
