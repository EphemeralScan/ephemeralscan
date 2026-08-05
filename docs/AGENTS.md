# AGENTS.md

# EphemeralScan AI Agent Instructions

## Purpose

This file provides operational instructions for AI coding agents working on the EphemeralScan project.

Project architecture, coding standards, and roadmap are documented elsewhere.

---

# Before You Start

Read the following documents before making changes:

1. STANDARDS.md
2. ARCHITECTURE.md
3. ROADMAP.md
4. CONFIGURATION.md
5. CLI.md

Do not start implementation before understanding them.

---

# Responsibilities

AI agents may:

- implement approved features;
- write tests;
- improve documentation;
- fix bugs;
- refactor code without changing behavior.

AI agents must not make architectural decisions.

---

# Workflow

For every task:

1. Understand the request.
2. Follow existing architecture.
3. Keep changes minimal.
4. Run formatting.
5. Ensure the project builds.
6. Open a small Pull Request.

---

# Build Requirements

Before completing any task, always run:

```bash
go fmt ./...
go build ./...
```

Run tests when available.

Never submit code that does not build.

---

# Development Rules

- Keep Pull Requests small.
- One logical feature per Pull Request.
- Follow existing project structure.
- Prefer simple solutions.
- Prefer Go standard library.
- Keep code readable.
- Keep functions focused.

---

# AI Principles

- Ask instead of assuming.
- Explain trade-offs when appropriate.
- Recommend one solution.
- Respect existing architecture.
- Follow project standards.

---

# AI Must Never

AI must never:

- invent APIs;
- silently change architecture;
- introduce unnecessary dependencies;
- modify public interfaces without approval;
- commit directly to main;
- hardcode credentials;
- suppress errors;
- remove tests to make builds pass;
- optimize before measuring;
- rewrite unrelated code.

---

# Security

Never expose:

- API tokens;
- secrets;
- passwords;
- private keys.

Never log sensitive information.

---

# Goal

Write production-quality code that follows the project's architecture, standards, and philosophy.

When uncertain, ask.