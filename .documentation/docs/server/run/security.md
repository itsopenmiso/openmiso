---
layout: docs
page_title: Server Security
description: |-
  The Waypoint server must be secured accordingly. This section will outline the risks (and explicit non-risks) of a Waypoint server and mitigations the server may have towards these.
---

# Server Security

The Waypoint server must be secured accordingly. This section will outline
the risks (and explicit non-risks) of a Waypoint server and mitigations the
server may have towards these. This is not an explicit threat model; we plan on
refining and publishing a threat model in the near future.

The Waypoint server supports TLS for all inbound connections, and we recommend
always using this. The Waypoint server also requires authentication for all
connections, and multiple authentication tokens can be created and revoked.
Authentication is not optional.

The Waypoint server doesn't yet support fine-grained authorization. Anyone
with an authentication token can access any part of the system. Note though
that the Waypoint server _does not_ execute operations (as noted later). One
exception to this is the [entrypoint](../docs/entrypoint), which uses a special
authentication token that can only access entrypoint-related APIs.

The state stored by Waypoint in-memory and on-disk is not encrypted. We
cannot always detect tampering of the data. The state stored is primarily
metadata that does not impact runtime behavior of Waypoint, with the exception
of [application configuration](../docs/app-config).

The Waypoint server _does not_ require access to credentials or platforms
for builds, deploys, etc. The Waypoint server _does not_ run these operations.
These operations are coordinated with the server but executed locally in
the CLI. The server does store metadata about operations such as time of
execution, logs, etc.

The Waypoint server maintains long-running inbound connections from
any [entrypoints](../docs/entrypoint). These can be used to trigger an
[exec session](../docs/exec) if exec is enabled.

The Waypoint server stores any static [application configuration](../docs/app-config).
This may contain secrets.
