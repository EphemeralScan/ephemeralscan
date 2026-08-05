# Development Standards

## General Principles

- Simplicity over cleverness.
- Readability over brevity.
- Explicit over implicit.
- Security before convenience.
- Small, composable components.
- Fail fast.
- Zero unnecessary dependencies.

---

## Repository Rules

- English only.
- One feature per Pull Request.
- One logical change per commit.
- No direct commits to `main`.
- Every change requires a Pull Request.

---

## Documentation

Every public component must have documentation.

Every architectural decision must be documented.

Documentation is part of the source code.

---

## Coding Standards

- Go formatting is mandatory.
- golangci-lint must pass.
- Static analysis must pass.
- No commented-out code.
- No dead code.

---

## Error Handling

Errors are never ignored.

Every returned error is either:

- handled;
- wrapped;
- returned.

---

## Logging

Structured logging only.

No printf debugging.

---

## Testing

Every public package must contain tests.

Regression bugs require regression tests.

---

## Security

No secrets in repository.

No credentials in source code.

No hardcoded API keys.

Least privilege by default.

---

## Dependencies

Every dependency must have a reason.

Avoid large frameworks.

Prefer Go standard library whenever possible.

---


## AI Principles

- When uncertain, ask instead of assuming.
- Architecture is more important than implementation.
- Prefer simple solutions over complex ones.
- Follow existing project patterns.
- Explain non-obvious design decisions.

## AI Usage

- AI may generate code.
- Humans approve code.
- AI never bypasses code review.
- AI-generated code follows the same standards as human-written code.
- If multiple solutions exist, explain the trade-offs and recommend one.

## AI must never:

- invent APIs;
- silently change architecture;
- introduce new dependencies without justification;
- modify public interfaces without discussion;
- commit directly to main;
- hardcode credentials;
- suppress errors;
- remove tests to make builds pass;
- optimize code before measuring performance;
- rewrite unrelated code during feature implementation.