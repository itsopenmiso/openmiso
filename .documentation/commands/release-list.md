---
layout: commands
page_title: "Commands: Release list"
sidebar_title: "release list"
description: "List releases."
---

# Waypoint Release list

Command: `waypoint release list`

List releases.


## Usage

Usage: `waypoint release list [options]`


  Lists the releases that were created if the platform includes a releaser.

#### Global Options

- `-plain` - Plain output: no colors, no animation. The default is false.
- `-app=<string>` (`-a`) - App to target. Certain commands require a single app target for Waypoint configurations with multiple apps. If you have a single app, then this can be ignored.
- `-project=<string>` (`-p`) - Project to target.
- `-workspace=<string>` (`-w`) - Workspace to operate in.

#### Command Options

- `-workspace-all` - List builds in all workspaces for this project and application. The default is false.
- `-verbose` (`-V`) - Display more details about each release. The default is false.
- `-url` (`-u`) - Display release URL. The default is false.
- `-json` - Output the release information as JSON. The default is false.
- `-long-ids` - Show long identifiers rather than sequence numbers. The default is false.

#### Filter Options

- `-state=<string>` - Filter values to have the given status. One possible value from: error, running, success, unknown.
- `-physical-state=<string>` - Show values in the given physical states. One possible value from: any, created, destroyed, pending. The default is created.
- `-order-by=<string>` - Order the values by which field. One possible value from: start-time, complete-time.
- `-desc` - Sort the values in descending order. The default is false.
- `-limit=<uint>` - How many values to show.

