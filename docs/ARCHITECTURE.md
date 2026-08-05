# Architecture

## High-Level Architecture

                    +----------------------+
                    |        CLI           |
                    +----------+-----------+
                               |
                               v
                    +----------------------+
                    |     Core Engine      |
                    +----------+-----------+
                               |
        +-----------+----------+-----------+-----------+
        |           |                      |           |
        v           v                      v           v
+---------------+ +---------------+ +---------------+ +---------------+
| Job Engine    | | Provider API  | | Scanner API   | | Report Engine |
+---------------+ +---------------+ +---------------+ +---------------+
        |                 |                  |                |
        |                 |                  |                |
        v                 v                  v                v
     Workflow         Linode             ZAP             HTML
                      Hetzner            Nuclei          JSON
                      AWS                Nmap            SARIF
                      ...                ...             Markdown

                               |
                               v
                    +----------------------+
                    | Notification Engine  |
                    +----------------------+

## Core Philosophy

- EphemeralScan is an orchestration engine.

- It does not perform security assessments itself.

- Instead, it orchestrates cloud infrastructure, third-party scanners, reporting, and notifications through well-defined interfaces.                    

## Components

- CLI
- Core Engine
- Job Engine
- Provider API
- Scanner API
- Report Engine
- Notification Engine
- Configuration Manager
- Storage

## Typical Workflow

User

↓

CLI

↓

Load Configuration

↓

Create Job

↓

Provision Infrastructure

↓

Execute Scanners

↓

Collect Results

↓

Generate Reports

↓

Send Notifications

↓

Destroy Infrastructure

↓

Exit

## Design Principles

- Interface-first architecture
- Plugin-oriented design
- Immutable reports
- Idempotent operations
- Stateless execution
- Minimal dependencies
- Cloud agnostic
- Scanner agnostic

## Component Responsibilities

TBD

## Future Extensions

Future versions may support:

- Disaster Recovery Testing
- Backup Verification
- Docker Image Auditing
- Kubernetes Auditing
- Terraform Validation
- Compliance Checks
- AI-assisted Report Analysis