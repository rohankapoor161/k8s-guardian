# Incident Response Runbook

## Overview

This runbook provides standardized procedures for handling production incidents involving the Kubernetes Guardian platform.

## Severity Levels

### SEV-1: Critical
- Complete service outage
- Data loss or corruption
- Security breach

**Response time:** 15 minutes  
**Escalation:** Immediate to VP Platform + On-call engineer

### SEV-2: High
- Major feature degradation
- Performance impact >50%
- Partial availability loss

**Response time:** 30 minutes  
**Escalation:** Team lead + On-call engineer

### SEV-3: Medium
- Minor feature degradation
- Performance impact <50%
- Non-critical alert noise

**Response time:** 2 hours  
**Escalation:** On-call engineer only

## Initial Response Checklist

1. **Acknowledge** the incident in PagerDuty
2. **Join** the incident Slack channel (#incidents-YYYY-MM-DD)
3. **Assess** impact scope and severity
4. **Document** timeline in incident doc
5. **Communicate** status in #status-updates

## Post-Incident Review

All SEV-1 and SEV-2 incidents require a retrospective within 48 hours.

Template: [Postmortem Template](https://docs.google.com/templates/postmortem)

---
*Last updated: 2026-03-04*
