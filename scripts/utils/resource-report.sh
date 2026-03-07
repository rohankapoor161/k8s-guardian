#!/bin/bash
# Resource Usage Report Generator
# Generates daily resource utilization reports for Guardian services
# Author: Rohan Kapoor <rohan.kapoor@example.com>
# Date: 2026-03-07

set -euo pipefail

# Configuration
NAMESPACE="${GUARDIAN_NAMESPACE:-default}"
OUTPUT_DIR="${REPORT_DIR:-./reports}"
RETENTION_DAYS="${REPORT_RETENTION:-30}"
VERBOSE="${VERBOSE:-false}"

# Colors for terminal output
readonly RED='\033[0;31m'
readonly GREEN='\033[0;32m'
readonly YELLOW='\033[1;33m'
readonly NC='\033[0m' # No Color

# Logging helpers
log() { echo "[$(date -Iseconds)] $1"; }
info() { log "${GREEN}INFO${NC}: $1"; }
warn() { log "${YELLOW}WARN${NC}: $1" && return 0; }
error() { log "${RED}ERROR${NC}: $1" && return 1; }

# Check dependencies
check_deps() {
    local missing=()
    for cmd in kubectl jq date; do
        if ! command -v "$cmd" >/dev/null 2>&1; then
            missing+=("$cmd")
        fi
    done
    
    if [[ ${#missing[@]} -gt 0 ]]; then
        error "Missing required commands: ${missing[*]}"
        exit 1
    fi
}

# Generate timestamp for report file
generate_report_name() {
    local date_str
    date_str=$(date +%Y-%m-%d)
    echo "resource-report-${date_str}.json"
}

# Fetch pod resource usage from metrics API
fetch_metrics() {
    local namespace="$1"
    
    info "Fetching metrics for namespace: $namespace"
    
    # Note: This requires metrics-server to be installed
    kubectl top pods -n "$namespace" --containers 2>/dev/null | \
        awk 'NR>1 {print "{\"pod\":\""$1"\",\"container\":\""$2"\",\"cpu\":\""$3"\",\"memory\":\""$4"\"}"}' || \
        warn "Metrics not available (metrics-server may not be installed)"
}

# Cleanup old reports
cleanup_old_reports() {
    local report_dir="$1"
    local retention="$2"
    
    info "Cleaning up reports older than $retention days"
    
    find "$report_dir" -name "resource-report-*.json" -mtime +"$retention" -delete 2>/dev/null || true
}

# Main function
main() {
    info "Starting resource usage report generation"
    
    # Validate environment
    check_deps
    
    # Create output directory if needed
    mkdir -p "$OUTPUT_DIR"
    
    local report_file
    report_file="$OUTPUT_DIR/$(generate_report_name)"
    
    # Generate report header
    cat > "$report_file" << EOF
{
  "report_type": "resource_utilization",
  "generated_at": "$(date -Iseconds)",
  "namespace": "$NAMESPACE",
  "version": "1.0.0",
  "data": [
EOF
    
    # Fetch and append metrics
    fetch_metrics "$NAMESPACE" >> "$report_file" || true
    
    # Close JSON structure
    echo "  ]" >> "$report_file"
    echo "}" >> "$report_file"
    
    info "Report generated: $report_file"
    
    # Cleanup old reports
    cleanup_old_reports "$OUTPUT_DIR" "$RETENTION_DAYS"
    
    info "Resource report generation complete"
}

# Run main function
main "$@"
