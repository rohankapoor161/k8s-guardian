package config

import (
    "testing"
)

func TestLoadConfig(t *testing.T) {
    // Test loading from non-existent file uses defaults
    cfg, err := Load("")
    if err != nil {
        t.Fatalf("expected no error for empty config path, got %v", err)
    }
    if cfg == nil {
        t.Fatal("expected config, got nil")
    }
}

func TestGateConfig_Validate(t *testing.T) {
    cfg := &GateConfig{}
    err := cfg.Validate()
    if err != nil {
        t.Errorf("validation failed: %v", err)
    }
}
