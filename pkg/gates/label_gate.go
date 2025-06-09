package gates

import (
    "fmt"
    corev1 "k8s.io/api/core/v1"
)

// LabelGate validates required labels
func (v *Validator) validateLabels(labels map[string]string) []Issue {
    var issues []Issue
    
    required := []string{
        "app.kubernetes.io/name",
        "app.kubernetes.io/version",
    }
    
    for _, key := range required {
        if _, ok := labels[key]; !ok {
            issues = append(issues, Issue{
                Gate:     "labels",
                Rule:     "missing-required",
                Message:  fmt.Sprintf("missing required label: %s", key),
                Severity: SeverityError,
            })
        }
    }
    
    return issues
}
