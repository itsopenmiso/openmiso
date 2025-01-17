---
layout: commands
page_title: "Commands: Pipeline run"
sidebar_title: "pipeline run"
description: "Manually execute a pipeline by name."
---

# Waypoint Pipeline run

Command: `waypoint pipeline run`

Manually execute a pipeline by name.


## Usage

Usage: `waypoint pipeline run [options] <pipeline-name>`


	Run a pipeline by name. If run outside of a project dir, a '-project' flag is
	required. Before running a requested pipeline, this command will sync
	pipeline configs so it runs the most up to date configuration version for a
	pipeline.

	If '-reattach' is supplied, the CLI will attempt to reattach to an existing
	pipeline run. Defaults to latest, but if '-run' is specified, it will attach
	to that specific run by sequence number.

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

- `-id=<string>` - Run a pipeline by ID.
- `-reattach` - If set, will replay or reattach to an existing pipeline run. If '-run' is not specified, will attempt to read the latest run. The default is false.
- `-run=<int>` - Replay or attach to a specific pipeline run by sequence number.

