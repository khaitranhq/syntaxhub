package filesystem

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
)

const (
	DefaultTimeoutSeconds = 2
)

// GetGoModuleRootPath returns the root path of the current Go module
func GetGoModuleRootPath() (string, error) {
	goModPathBytes, err := exec.CommandContext(context.Background(), "go", "env", "GOMOD").Output()
	if err != nil {
		return "", fmt.Errorf("failed to get Go module path: %w", err)
	}

	goModPath := string(goModPathBytes)

	return filepath.Dir(goModPath), nil
}
