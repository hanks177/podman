package abi

import (
	"context"

	"github.com/hanks177/podman/v4/libpod/define"
	"github.com/hanks177/podman/v4/pkg/domain/entities"
)

func (ic *ContainerEngine) HealthCheckRun(ctx context.Context, nameOrID string, options entities.HealthCheckOptions) (*define.HealthCheckResults, error) {
	status, err := ic.Libpod.HealthCheck(nameOrID)
	if err != nil {
		return nil, err
	}
	hcStatus := define.HealthCheckUnhealthy
	if status == define.HealthCheckSuccess {
		hcStatus = define.HealthCheckHealthy
	}
	report := define.HealthCheckResults{
		Status: hcStatus,
	}
	return &report, nil
}
