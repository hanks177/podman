//go:build remote
// +build remote

package infra

import (
	"errors"

	"github.com/hanks177/podman/v4/pkg/domain/entities"
)

// NewSystemEngine factory provides a libpod runtime for specialized system operations
func NewSystemEngine(setup entities.EngineSetup, facts *entities.PodmanConfig) (entities.SystemEngine, error) {
	return nil, errors.New("not implemented")
}
