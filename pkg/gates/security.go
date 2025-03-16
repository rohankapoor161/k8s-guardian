package gates

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

// SecurityGate enforces security policies
type SecurityGate struct {
	name        string
	description string
	enabled     bool
}

// NewSecurityGate creates a new security gate
func NewSecurityGate() *SecurityGate {
	return &SecurityGate{
		name:        "security",
		description: "Enforces security policies",
		enabled:     true,
	}
}

// Validate checks if pod meets security requirements
func (g *SecurityGate) Validate(ctx context.Context, pod *corev1.Pod) error {
	if !g.enabled {
		return nil
	}

	// Check privileged containers
	for _, container := range pod.Spec.Containers {
		if container.SecurityContext != nil && 
		   container.SecurityContext.Privileged != nil && 
		   *container.SecurityContext.Privileged {
			return fmt.Errorf("privileged containers not allowed: %s", container.Name)
		}
	}

	return nil
}

// Name returns gate name
func (g *SecurityGate) Name() string {
	return g.name
}
