package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rohankapoor161/k8s-guardian/pkg/gates"
	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// Server handles admission requests
type Server struct {
	gates  []gates.Gate
	decoder *admission.Decoder
}

// NewServer creates a webhook server
func NewServer(gates []gates.Gate) *Server {
	return &Server{gates: gates}
}

// InjectDecoder injects the decoder
func (s *Server) InjectDecoder(d *admission.Decoder) error {
	s.decoder = d
	return nil
}

// Handle processes admission requests
func (s *Server) Handle(ctx context.Context, req admission.Request) admission.Response {
	pod := &corev1.Pod{}
	err := s.decoder.Decode(req, pod)
	if err != nil {
		return admission.Errored(400, err)
	}

	for _, gate := range s.gates {
		if err := gate.Validate(ctx, pod); err != nil {
			return admission.Denied(err.Error())
		}
	}

	return admission.Allowed("All gates passed")
}
