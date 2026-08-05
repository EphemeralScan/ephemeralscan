# Command Line Interface (CLI)

## Overview

EphemeralScan uses a command-based CLI inspired by tools such as Git, Docker, and kubectl.

```text
ephemeralscan <command> [options]
```

The CLI is designed to be:

- Simple
- Predictable
- Automation-first
- Script-friendly
- Backward compatible whenever possible

---

# Commands

## scan

Execute a security assessment.

### Examples

```bash
ephemeralscan scan --config ephemeralscan.yaml

ephemeralscan scan --target https://example.com

ephemeralscan scan --profile production
```

---

## validate

Validate a configuration file without executing a scan.

### Examples

```bash
ephemeralscan validate ephemeralscan.yaml
```

---

## config

Manage configuration files.

### Commands

```bash
ephemeralscan config init

ephemeralscan config validate

ephemeralscan config show
```

### Description

- **init** – Create a default configuration template.
- **validate** – Validate the configuration file.
- **show** – Display the effective configuration.

---

## providers

List available cloud providers.

### Examples

```bash
ephemeralscan providers list

ephemeralscan providers info linode
```

---

## scanners

List available scanners.

### Examples

```bash
ephemeralscan scanners list

ephemeralscan scanners info zap
```

---

## reports

Work with existing reports.

### Examples

```bash
ephemeralscan reports list

ephemeralscan reports compare report1 report2

ephemeralscan reports export report.html
```

---

## doctor

Verify the local environment.

### Examples

```bash
ephemeralscan doctor
```

Typical output:

```text
✓ Configuration

✓ Docker

✓ Internet connectivity

✓ Cloud Provider

✓ API Credentials

✓ Required Tools

✓ Permissions
```

---

## version

Display version information.

### Examples

```bash
ephemeralscan version
```

Typical output:

```text
Version:    v0.1.0

Commit:     8f2d9e1

Build Date: 2026-08-04

Go Version: go1.25
```

---

# Global Options

```text
--config <file>

--verbose

--debug

--quiet

--output <directory>

--format html|json|markdown|sarif

--no-color

--help
```

---

# Exit Codes

| Code | Meaning |
|------:|---------|
| 0 | Success |
| 1 | General error |
| 2 | Invalid configuration |
| 3 | Provider error |
| 4 | Scanner error |
| 5 | Report generation error |
| 6 | Internal error |

---

# Design Principles

- No interactive mode.
- No hidden behavior.
- No automatic configuration generation.
- Explicit commands only.
- Human-friendly.
- Script-friendly.
- Stable command names.
- Clear and actionable error messages.
- Machine-readable exit codes.
- Automation-first.