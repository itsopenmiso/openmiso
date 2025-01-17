---
layout: commands
page_title: "Commands: Config source-delete"
sidebar_title: "config source-delete"
description: "Delete the configuration for a dynamic source plugin"
---

# Waypoint Config source-delete

Command: `waypoint config source-delete`

Delete the configuration for a dynamic source plugin


## Usage

Usage: `waypoint config source-delete [options]`


  Delete the configuration for a dynamic configuration source plugin.

  To use this command, you should specify a "-type" flag. Please see the
  documentation for the config source type you're configuring for details on
  what configuration fields are available.

  Configuration for this command is global. The "-app", "-project", and
  "-workspace" flags are ignored on this command.

#### Global Options

- `-plain` - Plain output: no colors, no animation. The default is false.
- `-app=<string>` (`-a`) - App to target. Certain commands require a single app target for Waypoint configurations with multiple apps. If you have a single app, then this can be ignored.
- `-project=<string>` (`-p`) - Project to target.
- `-workspace=<string>` (`-w`) - Workspace to operate in.

#### Command Options

- `-type=<string>` - Dynamic source type to delete, such as 'vault'.
- `-scope=<string>` - The scope for this configuration source. The configuration source will only delete within this scope. This can be one of 'global', 'project', or 'app'. The default is global.

