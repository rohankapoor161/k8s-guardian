package gates

import (
	"context"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"
)

func TestSecurityGate_Validate(t *testing.T) {
	ctx := context.Background()
	gate := NewSecurityGate()

	tests := []struct {
		name    string
		pod     *corev1.Pod
		wantErr bool
	}{
		{
			name: "privileged container denied",
			pod: &corev1.Pod{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "app",
							Image: "nginx",
							SecurityContext: &corev1.SecurityContext{
								Privileged: pointer.Bool(true),
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "non-privileged allowed",
			pod: &corev1.Pod{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "app",
							Image: "nginx",
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := gate.Validate(ctx, tt.pod)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
