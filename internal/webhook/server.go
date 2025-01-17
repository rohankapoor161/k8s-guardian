package webhook

import (
    "encoding/json"
    "fmt"
    "net/http"

    admissionv1 "k8s.io/api/admission/v1"
)

// Server handles admission webhook requests
type Server struct {
    port    int
    mux     *http.ServeMux
}

// NewServer creates a new webhook server
func NewServer(port int) *Server {
    s := &Server{
        port: port,
        mux:  http.NewServeMux(),
    }
    s.mux.HandleFunc("/validate", s.handleValidate)
    s.mux.HandleFunc("/health", s.handleHealth)
    return s
}

// Start begins listening for webhook requests
func (s *Server) Start() error {
    return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux)
}

func (s *Server) handleValidate(w http.ResponseWriter, r *http.Request) {
    // Parse admissionReview request
    var review admissionv1.AdmissionReview
    if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    // Validation logic here
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(review)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("healthy"))
}
