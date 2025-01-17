---
layout: commands
page_title: "Commands: Config delete"
sidebar_title: "config delete"
description: "Delete a config variable."
---

# Waypoint Config delete

Command: `waypoint config delete`

Delete a config variable.


## Usage

Usage: `waypoint config delete <name> <name2> ...`


  Delete a config variable from the system. This cannot be undone.

#### Global Options

- `-plain` - Plain output: no colors, no animation. The default is false.
- `-app=<string>` (`-a`) - App to target. Certain commands require a single app target for Waypoint configurations with multiple apps. If you have a single app, then this can be ignored.
- `-project=<string>` (`-p`) - Project to target.
- `-workspace=<string>` (`-w`) - Workspace to operate in.

#### Command Options

- `-scope=<string>` - The scope for this configuration to delete. The configuration will only delete within this scope. This can be one of 'global', 'project', or 'app'. The default is project.
- `-label-scope=<string>` - If set, configuration will only be deleted if the deployment or operation has a matching label set.

