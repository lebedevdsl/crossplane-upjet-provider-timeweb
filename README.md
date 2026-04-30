# crossplane-provider-timeweb

Crossplane provider for [Timeweb Cloud](https://timeweb.cloud).

Status: **planning** — see [`docs/PLAN.md`](docs/PLAN.md). No code yet.

## Approach

Generated with [Upjet](https://github.com/crossplane/upjet) from the official
[`timeweb-cloud/terraform-provider-timeweb-cloud`](https://github.com/timeweb-cloud/terraform-provider-timeweb-cloud)
(currently v1.6.15). Native (non-Upjet) managed resources will be added later
for API surfaces the Terraform provider doesn't expose, generated from the
vendored OpenAPI spec at [`docs/openapi-timeweb.json`](docs/openapi-timeweb.json).

## Why this exists

Drives the InYan platform migration to Timeweb Cloud — see
[`../timeweb/docs/initial-proposal.md`](../timeweb/docs/initial-proposal.md).
Crossplane replaces the Terraform layer described there with Kubernetes-native
managed resources reconciled by ArgoCD.

## Repo layout

See `docs/PLAN.md` §2.
