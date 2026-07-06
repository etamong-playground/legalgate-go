# Security Policy

## Reporting a Vulnerability

Please report security issues privately through this repository's **Security** tab
("Report a vulnerability"), not via public issues or pull requests.

This library is a fail-closed policy gate. If you find a case where a malformed,
missing, or ambiguous effective date causes `PolicyInForce` to return `true`
(i.e. it fails *open*), treat it as a security issue and report it privately.
