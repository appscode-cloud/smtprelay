# AGENTS.md

This file provides guidance to coding agents (e.g. Claude Code, claude.ai/code) when working with code in this repository.

## Repository purpose

The AppsCode fork of [decke/smtprelay](https://github.com/decke/smtprelay) — a small Go-based SMTP relay/proxy that accepts mail via SMTP and forwards it to another SMTP server. Used by AppsCode platform services to centralize outbound mail.

Module path is `go.bytebuilders.dev/smtprelay` (AppsCode-renamed). Git remote is `appscode-cloud/smtprelay`. Tracks upstream `decke/smtprelay`.

## Architecture

- Top-level Go files: `config.go`, `auth.go` (and tests). The SMTP server proper.
- `cmd/` — entry point.
- `Dockerfile.in` (PROD, distroless), `Dockerfile.dbg` (debian), `Dockerfile.ubi` (Red Hat certified).
- `hack/`, `Makefile` — AppsCode build harness.

## Common commands

- `make ci` — full CI pipeline.
- `make build` / `make all-build` — host or all-platform build.
- `make fmt`, `make lint`, `make unit-tests` / `make test` — standard.
- `make verify` — codegen + module-tidy verification.
- `make container` / `make push` / `make release` — image build/publish flow.

## Conventions

- Module path is `go.bytebuilders.dev/smtprelay`; imports must use that. Upstream is `github.com/decke/smtprelay` — keep AppsCode patches isolated.
- License: see `LICENSE` (matches upstream).
- Sign off commits (`git commit -s`).
- Single-purpose SMTP relay. Resist scope creep — adjacent functionality (templates, auth-token gen) belongs in `email-templates` / `smtprelay-auth-generator`.
- Three Dockerfiles, one binary — keep `Dockerfile.in`, `Dockerfile.dbg`, and `Dockerfile.ubi` in sync.
