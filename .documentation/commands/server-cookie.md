---
layout: commands
page_title: "Commands: Server cookie"
sidebar_title: "server cookie"
description: "Output server cookie value"
---

# Waypoint Server cookie

Command: `waypoint server cookie`

Output server cookie value


## Usage

Usage: `waypoint server cookie [options]`


  Output the server cookie value.

  The server cookie is used in API calls to superficially ensure that
  you're communicating with the proper cluster. This isn't mean to be a
  security mechanism. This is an optional way to prevent errant API calls
  to the incorrect Waypoint cluster (if you're running multiple).

  Some unauthenticated API endpoints require the cookie value be set to
  protect against random noise, such as the runner registration endpoint.

  While the cookie isn't a security mechanism, it should be kept secret
  to prevent unnecessary API noise to a cluster.

#### Global Options

- `-plain` - Plain output: no colors, no animation. The default is false.
- `-app=<string>` (`-a`) - App to target. Certain commands require a single app target for Waypoint configurations with multiple apps. If you have a single app, then this can be ignored.
- `-project=<string>` (`-p`) - Project to target.
- `-workspace=<string>` (`-w`) - Workspace to operate in.
