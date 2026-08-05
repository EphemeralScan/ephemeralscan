# Roadmap

This roadmap describes the planned evolution of EphemeralScan.

The project evolves incrementally. Every milestone should be considered stable before work begins on the next one.

---

# v0.1 — Minimum Viable Product

**Status:** In Progress

## Goal

Execute a complete security assessment using a temporary cloud instance.

## Features

- [ ] CLI
- [ ] YAML configuration
- [ ] Job Engine
- [ ] Linode provider
- [ ] cloud-init bootstrap
- [ ] Docker runtime
- [ ] OWASP ZAP Baseline Scan
- [ ] HTML report
- [ ] JSON report
- [ ] Structured logging
- [ ] Automatic infrastructure cleanup

## Success Criteria

- One command executes a complete assessment.
- Disposable infrastructure is mandatory.
- Infrastructure is always destroyed.
- Reports are reproducible.
- No manual intervention is required.

---

# v0.2 — Multiple Scanners

**Status:** Planned

## Goal

Support multiple security scanners through a common interface.

## Features

- [ ] Scanner SDK
- [ ] Nuclei
- [ ] Nmap
- [ ] Parallel execution
- [ ] Result aggregation

---

# v0.3 — Reports & Notifications

**Status:** Planned

## Goal

Improve reporting and result delivery.

## Features

- [ ] Markdown reports
- [ ] SARIF export
- [ ] Telegram notifications
- [ ] Email notifications
- [ ] Report comparison

---

# v0.4 — Multi Target

**Status:** Planned

## Goal

Support multiple targets within a single execution.

## Features

- [ ] Multiple targets
- [ ] Profiles
- [ ] Batch execution
- [ ] Scheduling support

---

# v0.5 — Multi Provider

**Status:** Planned

## Goal

Support multiple cloud providers.

## Features

- [ ] Provider SDK
- [ ] Hetzner
- [ ] DigitalOcean
- [ ] Vultr
- [ ] AWS

---

# v0.6 — History & Storage

**Status:** Planned

## Goal

Maintain execution history and collected artifacts.

## Features

- [ ] SQLite
- [ ] Scan history
- [ ] Artifact storage
- [ ] Report indexing
- [ ] Trend analysis

---

# v0.7 — Plugin Ecosystem

**Status:** Planned

## Goal

Allow third-party extensions.

## Features

- [ ] Plugin SDK
- [ ] Custom providers
- [ ] Custom scanners
- [ ] Custom reporters
- [ ] Custom notification modules

---

# v0.8 — Automation & CI/CD

**Status:** Planned

## Goal

Integrate EphemeralScan into automated security workflows.

## Features

- [ ] GitHub Actions
- [ ] GitLab CI
- [ ] Webhooks
- [ ] Exit codes
- [ ] Machine-readable outputs
- [ ] Scheduled execution
- [ ] Post-deployment validation
- [ ] Infrastructure-as-Code integration

---

# v1.0 — Stable

**Status:** Planned


## Goal

First production-ready release.

## Features

- [ ] Stable APIs
- [ ] Complete documentation
- [ ] Examples
- [ ] Integration tests
- [ ] Release binaries
- [ ] Installation packages
- [ ] Cross-platform support

---

# v2.0 — AI Assistance

**Status:** Planned

## Goal

Improve report analysis using AI.

## Features

- [ ] AI Provider SDK
- [ ] Risk prioritization
- [ ] Finding correlation
- [ ] Natural-language summaries
- [ ] Remediation suggestions

## Principles

- AI never replaces security scanners.
- AI analyzes existing findings only.
- AI recommendations must remain explainable.
- AI integration is optional.