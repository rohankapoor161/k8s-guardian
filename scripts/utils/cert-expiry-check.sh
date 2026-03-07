#!/bin/bash
# cert-expiry-check.sh - Utility to check TLS certificate expiration
# Author: Rohan Kapoor (VP Platform)
# Created: March 2026
# 
# This script checks the expiration dates of TLS certificates
# used by the Guardian webhook. Should be run as part of the
# monthly security audit.
#
# Usage: ./cert-expiry-check.sh [-d|--days WARNING_DAYS]

set -euo pipefail

NAMESPACE="guardian-system"
SECRET_NAME="guardian-cert"
WARNING_DAYS=${WARNING_DAYS:-30}

usage() {
    echo "Usage: $0 [-d|--days WARNING_DAYS]"
    echo "  -d, --days    Number of days before expiration to warn (default: 30)"
    exit 1
}

while [[ $# -gt 0 ]]; do
    case $1 in
        -d|--days)
            WARNING_DAYS="$2"
            shift 2
            ;;
        -h|--help)
            usage
            ;;
        *)
            echo "Unknown option: $1"
            usage
            ;;
    esac
done

# Check if kubectl is available
if ! command -v kubectl &> /dev/null; then
    echo "Error: kubectl not found in PATH"
    exit 1
fi

# Fetch certificate from secret
echo "Checking certificate expiry for $SECRET_NAME in namespace $NAMESPACE..."

CERT_DATA=$(kubectl get secret -n "$NAMESPACE" "$SECRET_NAME" -o jsonpath='{.data.tls\.crt}' 2>/dev/null | base64 -d 2>/dev/null || true)

if [[ -z "$CERT_DATA" ]]; then
    echo "Warning: Could not retrieve certificate from secret $SECRET_NAME"
    echo "This may indicate:"
    echo "  - The secret does not exist"
    echo "  - Insufficient permissions"
    echo "  - Certificate data is missing"
    exit 1
fi

# Parse expiration date
EXPIRY_DATE=$(echo "$CERT_DATA" | openssl x509 -noout -enddate 2>/dev/null | cut -d= -f2 || echo "Unknown")
DAYS_UNTIL_EXPIRY=$(echo "$CERT_DATA" | openssl x509 -noout -days 2>/dev/null || echo "0")

# Output results
echo ""
echo "Certificate Details:"
echo "  Secret:    $SECRET_NAME"
echo "  Namespace: $NAMESPACE"
echo "  Expires:   $EXPIRY_DATE"
echo "  Days Left: $DAYS_UNTIL_EXPIRY"
echo ""

if [[ "$DAYS_UNTIL_EXPIRY" -lt "$WARNING_DAYS" ]]; then
    echo "WARNING: Certificate expires in $DAYS_UNTIL_EXPIRY days (threshold: $WARNING_DAYS)"
    echo "Action Required: Schedule certificate renewal"
    exit 2
else
    echo "OK: Certificate valid for $DAYS_UNTIL_EXPIRY days"
    exit 0
fi
