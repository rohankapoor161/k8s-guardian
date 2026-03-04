#!/bin/bash
# Log Rotation Utility
# Rotates and compresses guardian logs to prevent disk pressure

set -euo pipefail

LOG_DIR="${GUARDIAN_LOG_DIR:-/var/log/guardian}"
RETENTION_DAYS="${LOG_RETENTION_DAYS:-30}"
COMPRESS_AFTER_DAYS=7

# Create directory if missing
mkdir -p "$LOG_DIR"

# Rotate current logs
cd "$LOG_DIR" || exit 1

# Find and compress logs older than COMPRESS_AFTER_DAYSind . -name "*.log" -type f -mtime +$COMPRESS_AFTER_DAYS -not -name "*.gz" | while read -r logfile; do
    echo "Compressing: $logfile"
    gzip -k "$logfile" && rm "$logfile"
done

# Remove compressed logs older than RETENTION_DAYSind . -name "*.gz" -type f -mtime +$RETENTION_DAYS -delete

# Log rotation summary
echo "[$(date -Iseconds)] Log rotation complete"
echo "  - Directory: $LOG_DIR"
echo "  - Retention: $RETENTION_DAYS days"
echo "  - Compressed: $(find . -name '*.gz' | wc -l) files"

# Optional: Send metrics to monitoring
if command -v curl >/dev/null 2>&1; then
    curl -s -X POST "${METRICS_ENDPOINT:-http://localhost:9090/metrics}" \
        -d "log_rotation_completed $(date +%s)"
fi
