package integration

import (
	"context"
	"testing"
	"time"

	"github.com/rohankapoor161/k8s-guardian/pkg/webhook"
)

func TestWebhookServer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server := webhook.NewServer(nil)
	
	// Start server in background
	go func() {
		if err := server.Start(ctx, ":8443"); err != nil {
			t.Logf("Server stopped: %v", err)
		}
	}()

	// Wait for server to be ready
	time.Sleep(100 * time.Millisecond)

	t.Log("Webhook server test complete")
}
