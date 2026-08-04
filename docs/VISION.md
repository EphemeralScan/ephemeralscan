# Vision

## Why EphemeralScan Exists

EphemeralScan started as a personal project.

I wanted a simple way to periodically verify the security posture of my homelab and Internet-facing services from an external network without maintaining dedicated scanning infrastructure.

Instead of keeping another VPS running 24/7, I wanted an ephemeral approach:

- create a temporary cloud instance;
- execute one or more security assessment jobs;
- collect reports;
- destroy the infrastructure completely.

The idea proved useful enough to evolve into a general-purpose open-source framework.

---

## Mission

EphemeralScan is an open-source framework for performing disposable, cloud-based security assessments.

The project creates short-lived infrastructure on demand, executes security assessment jobs, securely collects the results, and destroys the environment afterwards.

---

## Goals

- Security by isolation.
- Provider-independent architecture.
- Scanner-independent architecture.
- Simple CLI.
- Reproducible scans.
- Disposable infrastructure.
- Plugin-based architecture.
- High-quality reports.
- Automation-friendly.
- Open-source from day one.

---

## Non-Goals

EphemeralScan is NOT:

- a vulnerability scanner;
- a penetration testing framework;
- an exploit framework;
- a vulnerability database;
- a SIEM;
- an EDR.

EphemeralScan orchestrates existing security tools.

---

## Core Principles

1. Infrastructure is temporary.
2. Everything is reproducible.
3. Every action is logged.
4. Every component is replaceable.
5. AI is optional.
6. Cloud providers are interchangeable.
7. Scanners are interchangeable.
8. Reports are immutable.
9. Security comes before convenience.
10. Simple things should remain simple.

---

## Long-Term Vision

EphemeralScan aims to become a modular framework for securely executing ephemeral cloud workloads.

Security assessment is the first implemented workload, but the architecture should support many others without redesigning the core.