# crossplane-provider-timeweb

Crossplane provider for [Timeweb Cloud](https://timeweb.cloud).

Status: **MVP** — 21 managed resources generated and verified against a live
Timeweb account.

## Approach

Generated with [Upjet](https://github.com/crossplane/upjet) from the official
[`timeweb-cloud/terraform-provider-timeweb-cloud`](https://github.com/timeweb-cloud/terraform-provider-timeweb-cloud)
(currently v1.6.16). Native (non-Upjet) managed resources will be added later
for API surfaces the Terraform provider doesn't expose.

## References

- Timeweb Cloud REST API docs: <https://timeweb.cloud/api-docs>
- Timeweb Terraform provider: <https://github.com/timeweb-cloud/terraform-provider-timeweb-cloud>

## Why this exists

Built to drive a Kubernetes-based migration to Timeweb Cloud, replacing
Terraform with declarative managed resources reconciled by Crossplane (and
typically ArgoCD on top).
