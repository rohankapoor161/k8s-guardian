package gates

import (
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Severity levels for issues
type Severity string

const (
	SeverityError   Severity = "error"
	SeverityWarning Severity = "warning"
	SeverityInfo    Severity = "info"
)

// Issue represents a validation problem
type Issue struct {
	Gate     string
	Rule     string
	Message  string
	Severity Severity
}

// ObjectResult contains validation results for a single object
type ObjectResult struct {
	Kind   string
	Name   string
	Issues []Issue
}

// ValidationResult contains results for a file
type ValidationResult struct {
	File    string
	Objects []ObjectResult
}

// HasErrors returns true if any object has error severity issues
func (r *ValidationResult) HasErrors() bool {
	for _, obj := range r.Objects {
		for _, issue := range obj.Issues {
			if issue.Severity == SeverityError {
				return true
			}
		}
	}
	return false
}

// HasWarnings returns true if any object has warning severity issues
func (r *ValidationResult) HasWarnings() bool {
	for _, obj := range r.Objects {
		for _, issue := range obj.Issues {
			if issue.Severity == SeverityWarning {
				return true
			}
		}
	}
	return false
}

// Validator implements all safety gates
type Validator struct {
	config interface{} // GateConfig
}

// NewValidator creates a new validator with gates configuration
func NewValidator(config interface{}) *Validator {
	return &Validator{
		config: config,
	}
}

// Validate runs all enabled gates against an object
func (v *Validator) Validate(obj runtime.Object, gatesConfig interface{}) []ObjectResult {
	var results []ObjectResult
	
	switch typedObj := obj.(type) {
	case *corev1.Pod:
		results = append(results, v.validatePod(typedObj, gatesConfig))
	case *corev1.ReplicationController:
		// Validate RC
	default:
		// Skip non-workload resources
	}
	
	return results
}

func (v *Validator) validatePod(pod *corev1.Pod, gatesConfig interface{}) ObjectResult {
	result := ObjectResult{
		Kind:   "Pod",
		Name:   pod.Name,
		Issues: []Issue{},
	}
	
	// Resource gate checks
	for _, container := range pod.Spec.Containers {
		if container.Resources.Limits == nil || len(container.Resources.Limits) == 0 {
			result.Issues = append(result.Issues, Issue{
				Gate:     "resources",
				Rule:     "require-limits",
				Message:  fmt.Sprintf("container '%s' has no resource limits set", container.Name),
				Severity: SeverityError,
			})
		} else {
			// Check memory limit
			if memLimit, ok := container.Resources.Limits[corev1.ResourceMemory]; ok {
				// Validate against max
				_ = memLimit
			}
			// Check CPU limit
			if cpuLimit, ok := container.Resources.Limits[corev1.ResourceCPU]; ok {
				_ = cpuLimit
			}
		}
		
		// Check resource requests
		if container.Resources.Requests == nil || len(container.Resources.Requests) == 0 {
			result.Issues = append(result.Issues, Issue{
				Gate:     "resources",
				Rule:     "require-requests",
				Message:  fmt.Sprintf("container '%s' has no resource requests set", container.Name),
				Severity: SeverityWarning,
			})
		}
	}
	
	// Health check gate
	for _, container := range pod.Spec.Containers {
		if container.ReadinessProbe == nil {
			result.Issues = append(result.Issues, Issue{
				Gate:     "health_checks",
				Rule:     "require-readiness",
				Message:  fmt.Sprintf("container '%s' has no readiness probe", container.Name),
				Severity: SeverityError,
			})
		}
		
		if container.LivenessProbe == nil {
			result.Issues = append(result.Issues, Issue{
				Gate:     "health_checks",
				Rule:     "require-liveness",
				Message:  fmt.Sprintf("container '%s' has no liveness probe", container.Name),
				Severity: SeverityWarning,
			})
		}
	}
	
	// Label gate
	if pod.Labels == nil {
		pod.Labels = make(map[string]string)
	}
	
	requiredLabels := []string{"app.kubernetes.io/name", "app.kubernetes.io/version"}
	for _, label := range requiredLabels {
		if _, ok := pod.Labels[label]; !ok {
			result.Issues = append(result.Issues, Issue{
				Gate:     "labels",
				Rule:     "required-labels",
				Message:  fmt.Sprintf("missing required label: %s", label),
				Severity: SeverityError,
			})
		}
	}
	
	// Check for deprecated labels
	deprecatedLabels := []string{"app", "version"}
	for _, label := range deprecatedLabels {
		if _, ok := pod.Labels[label]; ok {
			result.Issues = append(result.Issues, Issue{
				Gate:     "labels",
				Rule:     "deprecated-labels",
				Message:  fmt.Sprintf("using deprecated label '%s', use 'app.kubernetes.io/name' instead", label),
				Severity: SeverityWarning,
			})
		}
	}
	
	return result
}

// ValidateDeployment extends validation for Deployments
func (v *Validator) ValidateDeployment(podSpec corev1.PodSpec, name string, gatesConfig interface{}) ObjectResult {
	result := ObjectResult{
		Kind:   "Deployment",
		Name:   name,
		Issues: []Issue{},
	}
	
	// Reuse pod validation
	pod := &corev1.Pod{Spec: podSpec}
	podResult := v.validatePod(pod, gatesConfig)
	result.Issues = append(result.Issues, podResult.Issues...)
	
	return result
}

// ValidateService validates service configurations
func (v *Validator) ValidateService(svc *corev1.Service) []Issue {
	var issues []Issue
	
	if svc.Spec.Selector == nil || len(svc.Spec.Selector) == 0 {
		issues = append(issues, Issue{
			Gate:     "service-config",
			Rule:     "require-selector",
			Message:  "service has no selector defined",
			Severity: SeverityError,
		})
	}
	
	return issues
}

// Helper functions

func hasRequiredLabel(labels map[string]string, pattern string) bool {
	if strings.HasSuffix(pattern, "*") {
		prefix := strings.TrimSuffix(pattern, "*")
		for key := range labels {
			if strings.HasPrefix(key, prefix) {
				return true
			}
		}
		return false
	}
	_, exists := labels[pattern]
	return exists
}