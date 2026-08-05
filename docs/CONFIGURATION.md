# Configuration

## Overview

EphemeralScan uses a single YAML configuration file.

Default filename:

```text
ephemeralscan.yaml
```

The configuration format is designed to be:

- Explicit
- Predictable
- Human-readable
- Automation-friendly
- Versioned

---

# Configuration Structure

```yaml
version: 1

job:
  name: example-scan

target:
  url: https://example.com

provider:
  name: linode
  region: ca-central
  instance_type: g6-nanode-1

scanners:
  - name: zap
    profile: baseline

reports:
  directory: ./reports
  formats:
    - html
    - json

cleanup:
  always_destroy: true
```

---

# Sections

## version

Configuration schema version.

```yaml
version: 1
```

Required.

---

## job

Execution metadata.

```yaml
job:
  name: example-scan
```

Fields:

- `name`

---

## target

Defines the assessment target.

```yaml
target:
  url: https://example.com
```

Fields:

- `url`

Required.

---

## provider

Defines the temporary cloud infrastructure.

```yaml
provider:
  name: linode
  region: ca-central
  instance_type: g6-nanode-1
```

Fields:

- `name`
- `region`
- `instance_type`

Credentials are never stored here.

---

## scanners

Defines the scanners to execute.

```yaml
scanners:
  - name: zap
    profile: baseline
```

Multiple scanners are supported.

---

## reports

Defines report generation.

```yaml
reports:
  directory: ./reports

  formats:
    - html
    - json
```

Fields:

- `directory`
- `formats`

---

## cleanup

Defines infrastructure cleanup.

```yaml
cleanup:
  always_destroy: true
```

### v0.1 Restriction

Infrastructure cleanup is mandatory.

The option cannot be disabled.

If cleanup is disabled, EphemeralScan must terminate before creating any cloud resources.

---

# Credentials

Credentials are supplied using environment variables.

Example:

```bash
export EPHEMERALSCAN_LINODE_TOKEN=xxxxxxxx
```

Credentials must never:

- appear in configuration files;
- appear in Git;
- appear in logs;
- appear in reports.

---

# Environment Variables

Every configuration value may be overridden.

Format:

```text
EPHEMERALSCAN_<SECTION>_<FIELD>
```

Example:

```bash
export EPHEMERALSCAN_PROVIDER_REGION=us-east
```

Priority:

1. Command-line arguments
2. Environment variables
3. Configuration file
4. Default values

---

# Configuration Commands

Create configuration:

```bash
ephemeralscan config init
```

Validate configuration:

```bash
ephemeralscan config validate
```

Show effective configuration:

```bash
ephemeralscan config show
```

---

# Validation

Validation must detect:

- missing required fields;
- invalid values;
- unsupported providers;
- unsupported scanners;
- invalid URLs;
- invalid report formats;
- conflicting options;
- unsafe configuration.

Validation must never create cloud resources.

---

# Design Principles

- No interactive configuration.
- No hidden behavior.
- No secrets in configuration files.
- No automatic configuration changes.
- Configuration is versioned.
- Invalid configuration fails before infrastructure provisioning.
- Backward compatibility whenever possible.