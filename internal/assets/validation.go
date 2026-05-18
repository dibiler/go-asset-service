package assets

import (
	"errors"
	"fmt"
	"strings"
)

var allowedTypes = map[string]bool{
	"server":         true,
	"database":       true,
	"container":      true,
	"cloud-resource": true,
	"storage":        true,
}

var allowedEnvironments = map[string]bool{
	"development": true,
	"staging":     true,
	"production":  true,
}

var allowedStatuses = map[string]bool{
	"active":   true,
	"inactive": true,
}

func ValidateAsset(a Asset) error {
	if strings.TrimSpace(a.ID) == "" {
		return errors.New("asset ID cannot be empty")
	}

	if strings.TrimSpace(a.Name) == "" {
		return errors.New("asset name cannot be empty")
	}

	if !allowedTypes[a.Type] {
		return fmt.Errorf("invalid asset type: %s", a.Type)

	}

	if !allowedEnvironments[a.Environment] {
		return fmt.Errorf("invalid asset environment: %s", a.Type)
	}

	if !allowedStatuses[a.Status] {
		return fmt.Errorf("invalid asset status: %s", a.Type)
	}

	if strings.TrimSpace(a.Location) == "" {
		return errors.New("location cannot be empty")
	}

	if a.CreatedAt.IsZero() {
		return errors.New("created_at cannot be empty")
	}
	if a.UpdatedAt.IsZero() {
		return errors.New("updated_at cannot be empty")
	}

	return nil
}
