#!/bin/bash
# Log Analyzer Utility
# Parses Guardian logs and outputs common failure patterns
# Usage: ./log-analyzer.sh [log-file-path] [--hours=24]

set -euo pipefail

# Configuration
DEFAULT_LOG_PATH="/var/log/guardian/guardian.log"
DEFAULT_HOURS=24
LOG_PATH="${1:-$DEFAULT_LOG_PATH}"
HOURS="${2:-$DEFAULT_HOURS}"

# Colors for output
RED='\033[0;31m'
YELLOW='\033[1;33m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

print_header() {
    echo ""
    echo "=========================================="
    echo "Guardian Log Analyzer"
    echo "Log: $LOG_PATH"
    echo "Window: Last $HOURS hours"
    echo "=========================================="
}

analyze_errors() {
    echo -e "\n${YELLOW}Error Summary:${NC}"
    
    # Count errors by type (simulated - would work with actual logs)
    if [[ -f "$LOG_PATH" ]]; then
        # Count admission errors
        local admission_errors
        admission_errors=$(grep -c "admission error" "$LOG_PATH" 2>/dev/null || echo "0")
        
        # Count webhook timeouts
        local webhook_timeouts
        webhook_timeouts=$(grep -c "webhook timeout" "$LOG_PATH" 2>/dev/null || echo "0")
        
        # Count policy violations
        local policy_violations
        policy_violations=$(grep -c "policy violation" "$LOG_PATH" 2>/dev/null || echo "0")
        
        echo "  Admission Errors: $admission_errors"
        echo "  Webhook Timeouts: $webhook_timeouts"
        echo "  Policy Violations: $policy_violations"
    else
        echo -e "  ${YELLOW}⚠ Log file not found: $LOG_PATH${NC}"
    fi
}

analyze_latency() {
    echo -e "\n${YELLOW}Latency Analysis:${NC}"
    
    if [[ -f "$LOG_PATH" ]]; then
        # Check for slow requests (>1s)
        local slow_requests
        slow_requests=$(grep -c "duration=.*[1-9]s" "$LOG_PATH" 2>/dev/null || echo "0")
        
        echo "  Slow Requests (>1s): $slow_requests"
        echo "  (Run with --verbose for detailed breakdown)"
    else
        echo -e "  ${YELLOW}⚠ No log data available${NC}"
    fi
}

print_recommendations() {
    echo -e "\n${GREEN}Recommendations:${NC}"
    
    cat << 'EOF'
  • Review any webhook timeouts - may indicate downstream service issues
  • Admission errors often correlate with invalid resource specs
  • Consider increasing timeout if latency spikes correlate with traffic
  • Escalate sustained error rates >5% to on-call engineer

EOF
}

print_footer() {
    echo "=========================================="
    echo "Analysis complete. See /docs/runbooks/ for"
    echo "detailed troubleshooting procedures."
    echo "=========================================="
}

# Main execution
print_header
analyze_errors
analyze_latency
print_recommendations
print_footer

exit 0
