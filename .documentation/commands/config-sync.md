---
layout: commands
page_title: "Commands: Config sync"
sidebar_title: "config sync"
description: "Synchronize declared variables and pipeline configs in a waypoint.hcl"
---

# Waypoint Config sync

Command: `waypoint config sync`

Synchronize declared variables and pipeline configs in a waypoint.hcl


## Usage

Usage: `waypoint config sync [options]`


  Synchronize declared application configuration in the waypoint.hcl file
  for existing and new deployments.

  Conflicting configuration keys will be overwritten. Configuration keys
  that do not exist in the waypoint.hcl file but exist on the server will not
  be deleted.

#### Global Options

- `-plain` - Plain output: no colors, no animation. The default is false.
- `-app=<string>` (`-a`) - App to target. Certain commands require a single app target for Waypoint configurations with multiple apps. If you have a single app, then this can be ignored.
- `-project=<string>` (`-p`) - Project to target.
- `-workspace=<string>` (`-w`) - Workspace to operate in.

#### Operation Options

- `-label=<key=value>` - Labels to set for this operation. Can be specified multiple times.
- `-local` - True to use a local runner to execute the operation, false to use a remote runner. 
If unset, Waypoint will automatically determine where the operation will occur, 
defaulting to remote if possible.
- `-remote-source=<key=value>` - Override configurations for how remote runners source data. This is specified to the data source type being used in your configuration. This is used for example to set a specific Git ref to run against.
- `-var=<key=value>` - Variable value to set for this operation. Can be specified multiple times.
- `-var-file=<string>` - HCL or JSON file containing variable values to set for this operation. If any "*.auto.wpvars" or "*.auto.wpvars.json" files are present, they will be automatically loaded.

