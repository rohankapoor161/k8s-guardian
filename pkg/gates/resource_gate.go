package gates

import (
    "fmt"
    corev1 "k8s.io/api/core/v1"
)

// ResourceGate validates container resource limits and requests
func (v *Validator) validateResources(container corev1.Container) []Issue {
    var issues []Issue

    // Check if limits are set
    if container.Resources.Limits == nil || len(container.Resources.Limits) == 0 {
        issues = append(issues, Issue{
            Gate:     "resources",
            Rule:     "require-limits",
            Message:  fmt.Sprintf("container '%s' missing resource limits", container.Name),
            Severity: SeverityError,
        })
        return issues
    }

    // Validate memory limit exists
    if _, ok := container.Resources.Limits[corev1.ResourceMemory]; !ok {
        issues = append(issues, Issue{
            Gate:     "resources",
            Rule:     "require-memory-limit",
            Message:  fmt.Sprintf("container '%s' missing memory limit", container.Name),
            Severity: SeverityError,
        })
    }

    // Validate CPU limit exists
    if _, ok := container.Resources.Limits[corev1.ResourceCPU]; !ok {
        issues = append(issues, Issue{
            Gate:     "resources",
            Rule:     "require-cpu-limit",
            Message:  fmt.Sprintf("container '%s' missing CPU limit", container.Name),
            Severity: SeverityError,
        })
    }

    // Check if requests are set
    if container.Resources.Requests == nil || len(container.Resources.Requests) == 0 {
        issues = append(issues, Issue{
            Gate:     "resources",
            Rule:     "require-requests",
            Message:  fmt.Sprintf("container '%s' missing resource requests", container.Name),
            Severity: SeverityWarning,
        })
    }

    return issues
}
