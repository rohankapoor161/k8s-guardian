package gates

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// ResourceGate enforces resource limits
type ResourceGate struct {
	minCPU     resource.Quantity
	minMemory  resource.Quantity
	maxCPU     resource.Quantity
	maxMemory  resource.Quantity
}

// NewResourceGate creates a resource validation gate
func NewResourceGate() *ResourceGate {
	return &ResourceGate{
		minCPU:    resource.MustParse("10m"),
		minMemory: resource.MustParse("32Mi"),
		maxCPU:    resource.MustParse("4000m"),
		maxMemory: resource.MustParse("8Gi"),
	}
}

// Validate checks resource requirements
func (g *ResourceGate) Validate(ctx context.Context, pod *corev1.Pod) error {
	for _, container := range pod.Spec.Containers {
		if container.Resources.Limits == nil {
			return fmt.Errorf("container %s: resource limits required", container.Name)
		}

		cpu := container.Resources.Limits.Cpu()
		memory := container.Resources.Limits.Memory()

		if cpu.IsZero() || memory.IsZero() {
			return fmt.Errorf("container %s: CPU and memory limits required", container.Name)
		}
	}

	return nil
}
