---
layout: commands
page_title: "Commands: Context create"
sidebar_title: "context create"
description: "Create a context."
---

# Waypoint Context create

Command: `waypoint context create`

Create a context.


## Usage

Usage: `waypoint context create [options] NAME`


  Creates a new context.

#### Global Options

- `-plain` - Plain output: no colors, no animation. The default is false.
- `-app=<string>` (`-a`) - App to target. Certain commands require a single app target for Waypoint configurations with multiple apps. If you have a single app, then this can be ignored.
- `-project=<string>` (`-p`) - Project to target.
- `-workspace=<string>` (`-w`) - Workspace to operate in.

#### Command Options

- `-set-default` - Set this context as the new default for the CLI. The default is true.
- `-server-addr=<string>` - Address for the server.
- `-server-auth-token=<string>` - Authentication token to use to connect to the server.
- `-server-platform=<string>` - The current platform that Waypoint server is running on. The default is n/a.
- `-server-tls` - If true, will connect to the server over TLS. The default is true.
- `-server-tls-skip-verify` - If true, will not validate TLS cert presented by the server. The default is false.
- `-server-require-auth` - If true, will send authentication details. The default is false.

