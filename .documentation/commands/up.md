---
layout: commands
page_title: "Commands: Up"
sidebar_title: "up"
description: "Perform the build, deploy, and release steps"
---

# Waypoint Up

Command: `waypoint up`

Perform the build, deploy, and release steps


## Usage

Usage: `waypoint up [options]`


  Perform the build, deploy, and release steps.

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

#### Command Options

- `-prune` - Prune old unreleased deployments. The default is true.
- `-prune-retain=<int>` - The number of unreleased deployments to keep. If this isn't set or is set to any negative number, then this will default to 1 on the server. If you want to prune all unreleased deployments, set this to 0. The default is -1.
