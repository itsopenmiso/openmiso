---
layout: commands
page_title: "Commands: Pipeline inspect"
sidebar_title: "pipeline inspect"
description: "Inspect the full details of a pipeline by id"
---

# Waypoint Pipeline inspect

Command: `waypoint pipeline inspect`

Inspect the full details of a pipeline by id


## Usage

Usage: `waypoint pipeline inspect [options] <pipeline-id>`


  Inspect the full details of a pipeline by id or name for a project.
  Including the '-run' option will show full details of a specific run of this pipeline. 

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

- `-json` - Output the Pipeline as json. The default is false.
- `-name=<string>` - Inspect a pipeline by name.
- `-run=<int>` - Inspect a specific run sequence.
