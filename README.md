# security-notes

Personal notes, Semgrep rules, and (later) n-day write-ups from my own
authorized open-source vulnerability research.

## semgrep-rules/

Rules I write for myself while auditing OSS projects — not yet submitted to
the official [semgrep-rules](https://github.com/semgrep/semgrep-rules)
registry, since that requires a much lower false-positive rate across many
real codebases than a rule tuned on one project does. Each rule ships with
a synthetic test file (`*.test.<ext>`), validated with `semgrep --test`.

| Rule | Class | What it looks for |
|---|---|---|
| [`go-dial-unvalidated-host`](semgrep-rules/go-dial-unvalidated-host.yaml) | SSRF-adjacent / credential disclosure (CWE-918, CWE-522) | A Go struct populated from a request body is used to pick the host/port of an outbound `.Dial(...)` call, with no guard/validate call on the struct's fields anywhere earlier in the same function. If a caller with a narrow, legitimately-granted permission (not full admin) can influence that struct, the server can be made to connect — and disclose a stored credential to, or use as a port-scanning oracle against — any address the caller chooses. |

Run it:
```bash
semgrep --config semgrep-rules/go-dial-unvalidated-host.yaml path/to/code
semgrep --test --config semgrep-rules/ semgrep-rules/go-dial-unvalidated-host.test.go
```

## Scope and disclosure

Findings from real projects go through private channels (GitHub Private
Vulnerability Reporting or the project's own security contact) first. A
write-up for a specific CVE lands here only after the advisory is published
and the embargo has lifted — never before. Rule files use invented,
generic example code, not excerpts from any report still under embargo.

## AI disclosure

Parts of this repository (rule drafting, test fixtures, write-ups) are
produced with AI assistance (Claude Code). Every finding behind a write-up
here is reproduced and verified locally by me before being described.
