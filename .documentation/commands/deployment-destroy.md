---
layout: commands
page_title: "Commands: Deployment destroy"
sidebar_title: "deployment destroy"
description: "Destroy one or more deployments."
---

# Waypoint Deployment destroy

Command: `waypoint deployment destroy`

Destroy one or more deployments.


## Usage

Usage: `waypoint deployment destroy [options] [id...]`


  Destroy one or more deployments. This will "undeploy" this specific
  instance of an application.

  When no arguments are given, this will default to destroying ALL
  unreleased deployments. This will require interactive confirmation
  by the user unless the force flag (-force) is specified.

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

- `-all` - Delete ALL deployments, including released. The default is false.
- `-force` - Yes to all confirmations. The default is false.
