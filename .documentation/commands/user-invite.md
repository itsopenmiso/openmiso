---
layout: commands
page_title: "Commands: User invite"
sidebar_title: "user invite"
description: "Invite a user to join the Waypoint server"
---

# Waypoint User invite

Command: `waypoint user invite`

Invite a user to join the Waypoint server


## Usage

Usage: `waypoint user invite [options]`


  Generate an invite token that can be used to log in.

  You must be logged in already to generate an invite token. If you need to
  log in, use the "waypoint login" command.

  This generates a new invite token. An invite token can be exchanged for
  a login token. If your Waypoint server has OIDC (non-token) auth enabled,
  it is recommended to instead invite users using your UI URL or directly
  via "waypoint login".

#### Global Options

- `-plain` - Plain output: no colors, no animation. The default is false.
- `-app=<string>` (`-a`) - App to target. Certain commands require a single app target for Waypoint configurations with multiple apps. If you have a single app, then this can be ignored.
- `-project=<string>` (`-p`) - Project to target.
- `-workspace=<string>` (`-w`) - Workspace to operate in.

#### Command Options

- `-expires-in=<duration>` - The duration until the token expires. i.e. '5m'. The default is 24h0m0s.
- `-username=<string>` - Invite a new user and provide a username hint. The user may still change their username after signing up. If this isn't specified, an invite token for your current user will be generated and no new user signup is performed.
