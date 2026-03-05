#!/bin/bash
# Health Check Utility
# Performs comprehensive health checks on Guardian services

set -euo pipefail

GUARDIAN_URL="${GUARDIAN_URL:-http://localhost:8080}"
TIMEOUT_SECONDS="${HEALTH_CHECK_TIMEOUT:-30}"
VERBOSE="${VERBOSE:-false}"

log() {
    if [[ "$VERBOSE" == "true" ]]; then
        echo "[$(date -Iseconds)] $1"
    fi
}

check_endpoint() {
    local endpoint="$1"
    local expected_status="${2:-200}"
    
    log "Checking endpoint: $endpoint"
    
    local http_code
    http_code=$(curl -s -o /dev/null -w "%{http_code}" --max-time "$TIMEOUT_SECONDS" \
        "$GUARDIAN_URL$endpoint" 2>/dev/null || echo "000")
    
    if [[ "$http_code" == "$expected_status" ]]; then
        echo "✓ $endpoint ($http_code)"
        return 0
    else
        echo "✗ $endpoint (expected $expected_status, got $http_code)"
        return 1
    fi
}

# Main health check flow
echo "================================"
echo "Guardian Health Check"
echo "Target: $GUARDIAN_URL"
echo "================================"

failed=0

# Core endpoints
check_endpoint "/health" 200 || ((failed++))
check_endpoint "/ready" 200 || ((failed++))
check_endpoint "/metrics" 200 || ((failed++))

# Summary
echo "================================"
if [[ $failed -eq 0 ]]; then
    echo "Status: HEALTHY ✓"
    exit 0
else
    echo "Status: UNHEALTHY ✗ ($failed checks failed)"
    exit 1
fi
