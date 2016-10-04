package places

import (
	"path/filepath"
)

// Recursive prints all places in a directory, recursiveley.
func Recursive(path string) error {
	if err := createMapsClient(); err != nil {
		return err
	}
	err := filepath.Walk(path, VisitPrintLocation)
	if err != nil {
		return err
	}
	return nil
}
